package property

import (
	"fmt"
	"github.com/mauricelam/genny/generic"
	"github.com/rs/xid"
	"reflect"
	"sync"
)

type PropValueType generic.Type

type PropertyPropValueType struct {
	Value       PropValueType `json:"value"`
	ReadOnly    bool          `json:"readOnly"`
	UserVisible bool          `json:"userVisible"`
	Description string        `json:"description"`
	Unit        string        `json:"unit"`
	Min         int           `json:"min"`
	Max         int           `json:"max"`
	mutex       *sync.RWMutex `json:"-"`
	id          xid.ID        `json:"-"`
}

func NewPropertyPropValueType(description string, defaultValue PropValueType, min, max int, readOnly, userVisible bool, unit string) *PropertyPropValueType {
	mutex := &sync.RWMutex{}
	id := xid.ID{}
	return &PropertyPropValueType{defaultValue, readOnly, userVisible, description, unit, max, min, mutex, id}
}

func NewPropertyPropValueTypeRW(description string, defaultValue PropValueType, min, max int, userVisible bool, unit string) *PropertyPropValueType {
	return NewPropertyPropValueType(description, defaultValue, min, max, false, userVisible, unit)
}

func NewPropertyPropValueTypeRO(description string, defaultValue PropValueType, min, max int, userVisible bool, unit string) *PropertyPropValueType {
	return NewPropertyPropValueType(description, defaultValue, min, max, true, userVisible, unit)
}

func NewPropertyPropValueTypeUserView(description string, defaultValue PropValueType, min, max int, unit string) *PropertyPropValueType {
	return NewPropertyPropValueType(description, defaultValue, min, max, true, true, unit)
}

func NewPropertyPropValueTypeUserViewRaw(description string, defaultValue PropValueType, min, max int) *PropertyPropValueType {
	return NewPropertyPropValueType(description, defaultValue, min, max, true, true, UnitNone)
}

func (prop *PropertyPropValueType) Write(value interface{}) error {
	prop.mutex.Lock()
	defer prop.mutex.Unlock()

	typeOfValue := reflect.TypeOf(value)
	typeOfPortValue := reflect.TypeOf(prop.Value)

	if !typeOfValue.ConvertibleTo(typeOfPortValue) {
		return fmt.Errorf("incompatible value types, given type: %v, could not be converted to: %v", typeOfValue, typeOfPortValue)
	}
	valueOfValue := reflect.ValueOf(value)
	prop.Value = valueOfValue.Convert(typeOfPortValue).Interface().(PropValueType)
	prop.id = xid.New()
	return nil
}

func (prop *PropertyPropValueType) WritePropValueType(value PropValueType) {
	prop.mutex.Lock()
	defer prop.mutex.Unlock()

	prop.Value = value
	prop.id = xid.New()
}

func (prop *PropertyPropValueType) Read() interface{} {
	prop.mutex.RLock()
	defer prop.mutex.RUnlock()

	return prop.Value
}

func (prop *PropertyPropValueType) ReadPropValueType() PropValueType {
	prop.mutex.RLock()
	defer prop.mutex.RUnlock()

	return prop.Value
}

func (prop *PropertyPropValueType) GetID() xid.ID {
	prop.mutex.RLock()
	defer prop.mutex.RUnlock()

	return prop.id
}
