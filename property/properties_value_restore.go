package property

import (
	"encoding/json"
	"reflect"
)

type PropertiesValueRestorer interface {
	Store(interface{}) []byte
	Restore(jsonString []byte, properties interface{}) error
}

type JSONPropertiesValueRestorer struct{}

// Store returns json representation of all properties combined into one tuple
func (j JSONPropertiesValueRestorer) Store(object interface{}) []byte {
	jsonString, _ := json.Marshal(object)
	return jsonString
}

// Restore loads last used values from json string into current properties.
// If given property was saved into loaded json its value is set to the one from json.
// Values which are not present in loaded json are left intact with their default values set
func (j JSONPropertiesValueRestorer) Restore(jsonString []byte, properties interface{}) error {
	var objmap map[string]*json.RawMessage
	err := json.Unmarshal(jsonString, &objmap)
	if err != nil {
		return err
	}

	propertiesValue := reflect.ValueOf(properties).Elem()
	for propertyName := range objmap {
		propertiesField := propertiesValue.FieldByName(propertyName)
		if propertiesField.IsValid() {
			var propertyObjectMap map[string]*json.RawMessage
			err = json.Unmarshal(*objmap[propertyName], &propertyObjectMap)
			if err != nil {
				return err
			}
			var propertyLoadedValue interface{}
			err = json.Unmarshal(*propertyObjectMap["value"], &propertyLoadedValue)
			if err != nil {
				return err
			}
			propertiesField.Interface().(PropertyWriter).Write(propertyLoadedValue)
		}
	}
	return nil
}
