package port

import (
	"fmt"
	"github.com/mauricelam/genny/generic"
	"github.com/rs/xid"
	"reflect"
	"sync"
)

type ValueType generic.Type

type OutputPortValueType struct {
	BaseOutputPort
	Mutex sync.RWMutex
	Value ValueType
}

type InputPortValueType struct {
	BaseInputPort
	Mutex     sync.RWMutex
	Value     ValueType
	PrevValue ValueType
}

func NewOutputPortValueType() *OutputPortValueType {
	return &OutputPortValueType{Mutex: sync.RWMutex{}}
}

func NewInputPortValueType(blockingType BlockingType) *InputPortValueType {
	return &InputPortValueType{Mutex: sync.RWMutex{}, BaseInputPort: BaseInputPort{blockingType, false}}
}

func (port *OutputPortValueType) GetID() xid.ID {
	return port.id
}

func (port *OutputPortValueType) Write(value ValueType) error {
	port.Mutex.Lock()
	defer port.Mutex.Unlock()

	port.Value = value
	port.id = xid.New()
	return nil
}

func (port *OutputPortValueType) read() interface{} {
	port.Mutex.RLock()
	defer port.Mutex.RUnlock()

	return port.Value
}

func (port *InputPortValueType) IsBlockingNew() bool {
	return port.blockingType == PortBlockingNew
}

func (port *InputPortValueType) IsBlockingDiff() bool {
	return port.blockingType == PortBlockingDiff
}

func (port *InputPortValueType) ValueChanged() bool {
	return port.Value != port.PrevValue
}

func (port *InputPortValueType) ValueNew() bool {
	return port.valueNew
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
	port.PrevValue = port.Value
	port.valueNew = true
	port.Value = valueOfValue.Convert(typeOfPortValue).Interface().(ValueType)
	return nil
}

func (port *InputPortValueType) Read() ValueType {
	port.Mutex.RLock()
	defer port.Mutex.RUnlock()

	// reset value freshness on read
	port.valueNew = false

	return port.Value
}
