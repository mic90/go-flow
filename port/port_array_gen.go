package port

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"time"

	"github.com/mauricelam/genny/generic"
)

type ArrayValueType generic.Type

type ArrayOutputPortArrayValueType struct {
	OutputPort
	Mutex sync.RWMutex
	Value []ArrayValueType
}

type ArrayInputPortArrayValueType struct {
	InputPort
	Mutex sync.RWMutex
	Value []ArrayValueType
	PrevValue []ArrayValueType
}

func NewArrayOutputPortArrayValueType() *ArrayOutputPortArrayValueType {
	array := make([]ArrayValueType, 10)
	return &ArrayOutputPortArrayValueType{OutputPort: OutputPort{}, Mutex: sync.RWMutex{}, Value: array}
}

func NewArrayOutputPortArrayValueTypeLen(len int) *ArrayOutputPortArrayValueType {
	array := make([]ArrayValueType, len)
	return &ArrayOutputPortArrayValueType{OutputPort: OutputPort{}, Mutex: sync.RWMutex{}, Value: array}
}

func NewArrayInputPortArrayValueType(requiredNew bool) *ArrayInputPortArrayValueType {
	array := make([]ArrayValueType, 10)
	prevArray := make([]ArrayValueType, 10)
	return &ArrayInputPortArrayValueType{InputPort: InputPort{RequiredNew: requiredNew}, Mutex: sync.RWMutex{}, Value: array, PrevValue: prevArray}
}

func NewArrayInputPortArrayValueTypeLen(requiredNew bool, len int) *ArrayInputPortArrayValueType {
	array := make([]ArrayValueType, len)
	prevArray := make([]ArrayValueType, len)
	return &ArrayInputPortArrayValueType{InputPort: InputPort{RequiredNew: requiredNew}, Mutex: sync.RWMutex{}, Value: array, PrevValue: prevArray}
}

func (port *ArrayOutputPortArrayValueType) GetTimestamp() time.Time {
	return port.Timestamp
}

// Write will write input slice value into the current port
// Input slice will be copied into internal port slice
// If internal port slice len is < input value it will be extended
func (port *ArrayOutputPortArrayValueType) Write(value []ArrayValueType) error {
	port.Mutex.Lock()
	defer port.Mutex.Unlock()

	if len(port.Value) < len(value) {
		port.Value = append(port.Value, make([]ArrayValueType, len(value)-len(port.Value))...)
	}

	copy(port.Value, value)
	port.Timestamp = time.Now()
	return nil
}

func (port *ArrayInputPortArrayValueType) IsRequiredNew() bool {
	return port.RequiredNew
}

func (port *ArrayInputPortArrayValueType) ValueChanged() bool {
	for i, value := range port.Value {
		if value != port.PrevValue[i] {
			return true
		}
	}
	return false
}

// read will return value currently stored in port
// Needed only for internal usage by graph
func (port *ArrayOutputPortArrayValueType) read() interface{} {
	port.Mutex.RLock()
	defer port.Mutex.RUnlock()

	return port.Value
}

func (port *ArrayInputPortArrayValueType) write(value interface{}) error {
	port.Mutex.Lock()
	defer port.Mutex.Unlock()

	typeOfValue := reflect.TypeOf(value)
	typeOfPortValue := reflect.TypeOf(port.Value)
	if typeOfValue.Kind() != reflect.Slice {
		return errors.New("unsupported value type, expected slice of values")
	}
	if !typeOfValue.ConvertibleTo(typeOfPortValue) {
		return fmt.Errorf("incompatible value types, given type: %v, could not be converted to: %v", typeOfValue, typeOfPortValue)
	}

	valueOfValue := reflect.ValueOf(value)
	copy(port.PrevValue, port.Value)
	port.Value = valueOfValue.Convert(typeOfPortValue).Interface().([]ArrayValueType)
	return nil
}

func (port *ArrayInputPortArrayValueType) Read() []ArrayValueType {
	port.Mutex.RLock()
	defer port.Mutex.RUnlock()

	return port.Value
}
