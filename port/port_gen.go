package port

import (
	"fmt"
	"github.com/mauricelam/genny/generic"
	"reflect"
	"sync"
	"time"
)

type ValueType generic.Type

type OutputPortValueType struct {
	OutputPort
	Mutex sync.RWMutex
	Value ValueType
}

type InputPortValueType struct {
	InputPort
	Mutex sync.RWMutex
	Value ValueType
}

func NewOutputPortValueType() *OutputPortValueType {
	return &OutputPortValueType{Mutex: sync.RWMutex{}}
}

func NewInputPortValueType(requiredNew bool) *InputPortValueType {
	return &InputPortValueType{Mutex: sync.RWMutex{}, InputPort: InputPort{RequiredNew: requiredNew}}
}

func (port *OutputPortValueType) GetTimestamp() time.Time {
	return port.Timestamp
}

func (port *OutputPortValueType) Write(value ValueType) error {
	port.Mutex.Lock()
	defer port.Mutex.Unlock()

	port.Value = value
	port.Timestamp = time.Now()
	return nil
}

func (port *OutputPortValueType) read() interface{} {
	port.Mutex.RLock()
	defer port.Mutex.RUnlock()

	return port.Value
}

func (port *InputPortValueType) IsRequiredNew() bool {
	return port.RequiredNew
}

func (port *InputPortValueType) write(value interface{}) error {
	port.Mutex.Lock()
	defer port.Mutex.Unlock()

	typeOfValue := reflect.TypeOf(value)
	typeOfPortValue := reflect.TypeOf(port.Value)
	if !typeOfValue.ConvertibleTo(typeOfPortValue) {
		return fmt.Errorf("incompatible value types, given type: %v, could not be converted to: %v", typeOfValue, typeOfPortValue)
	}
	valueOfValue := reflect.ValueOf(value)
	port.Value = valueOfValue.Convert(typeOfPortValue).Interface().(ValueType)
	return nil
}

func (port *InputPortValueType) Read() ValueType {
	port.Mutex.RLock()
	defer port.Mutex.RUnlock()

	return port.Value
}
