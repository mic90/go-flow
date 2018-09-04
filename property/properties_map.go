package property

import (
	"fmt"
	"github.com/rs/xid"
	"reflect"
	"strings"
	"sync"
)

type PropertyMapReader interface {
	Read(propertyName string) (interface{}, error)
	GetID(propertyName string) (*xid.ID, error)
}

type PropertyMapWriter interface {
	Write(propertyName string, value interface{}) error
}

type PropertyMapReaderWriter interface {
	PropertyMapReader
	PropertyMapWriter
}

type PropertiesMap struct {
	propertiesMapMutex sync.RWMutex                      `json:"-"`
	PropertiesMap      map[string]PropertyIDReaderWriter `json:"properties"`
}

func NewPropertiesMap(properties ...interface{}) *PropertiesMap {
	propertyType := reflect.TypeOf((*PropertyIDReaderWriter)(nil)).Elem()
	finalPropertiesMap := make(map[string]PropertyIDReaderWriter)
	for _, propertiesObj := range properties {
		objectType := reflect.TypeOf(propertiesObj).Elem()
		if objectType.Kind() != reflect.Struct {
			//TODO: do something maybe return some error ?
			continue
		}
		propertiesMap := structFieldsToMap(propertiesObj, propertyType).(map[string]PropertyIDReaderWriter)
		// write all struct field into map by field name
		for k, v := range propertiesMap {
			lowerCaseName := strings.ToLower(k)
			finalPropertiesMap[lowerCaseName] = v
		}
	}

	return &PropertiesMap{sync.RWMutex{}, finalPropertiesMap}
}

func (propc *PropertiesMap) Write(name string, value interface{}) error {
	propc.propertiesMapMutex.Lock()
	defer propc.propertiesMapMutex.Unlock()

	property, ok := propc.PropertiesMap[name]
	if !ok {
		return fmt.Errorf("no property with name %s was found", name)
	}
	property.Write(value)
	return nil
}

func (propc *PropertiesMap) Read(name string) (interface{}, error) {
	propc.propertiesMapMutex.Lock()
	defer propc.propertiesMapMutex.Unlock()
	property, ok := propc.PropertiesMap[name]
	if !ok {
		return nil, fmt.Errorf("no property with name %s was found", name)
	}
	return property.Read(), nil
}

func (propc *PropertiesMap) GetID(name string) (*xid.ID, error) {
	propc.propertiesMapMutex.Lock()
	defer propc.propertiesMapMutex.Unlock()
	property, ok := propc.PropertiesMap[name]
	if !ok {
		return nil, fmt.Errorf("no property with name %s was found", name)
	}
	propertyTime := property.GetID()
	return &propertyTime, nil
}

func structFieldsToMap(object interface{}, fieldType reflect.Type) interface{} {
	stringType := reflect.TypeOf("")
	mapType := reflect.MapOf(stringType, fieldType)
	fieldsMap := reflect.MakeMap(mapType)

	nodeValue := reflect.ValueOf(object).Elem()
	for i := 0; i < nodeValue.NumField(); i++ {
		fieldName := nodeValue.Type().Field(i).Name
		implements := nodeValue.Field(i).Type().Implements(fieldType)
		if implements {
			fieldsMap.SetMapIndex(reflect.ValueOf(fieldName), reflect.ValueOf(nodeValue.Field(i).Interface()))
		}
	}
	return fieldsMap.Interface()
}
