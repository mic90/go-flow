// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/mauricelam/genny

package port

import (
	"fmt"

	"reflect"

	"sync"

	"time"
)

type OutputPortByte struct {
	OutputPort
	Mutex sync.RWMutex
	Value byte
}

type InputPortByte struct {
	InputPort
	Mutex sync.RWMutex
	Value byte
}

func NewOutputPortByte() *OutputPortByte {
	return &OutputPortByte{Mutex: sync.RWMutex{}}
}

func NewInputPortByte(requiredNew bool) *InputPortByte {
	return &InputPortByte{Mutex: sync.RWMutex{}, InputPort: InputPort{RequiredNew: requiredNew}}
}

func (port *OutputPortByte) GetTimestamp() time.Time {
	return port.Timestamp
}

func (port *OutputPortByte) Write(value byte) error {
	port.Mutex.Lock()
	defer port.Mutex.Unlock()

	port.Value = value
	port.Timestamp = time.Now()
	return nil
}

func (port *OutputPortByte) read() interface{} {
	port.Mutex.RLock()
	defer port.Mutex.RUnlock()

	return port.Value
}

func (port *InputPortByte) IsRequiredNew() bool {
	return port.RequiredNew
}

func (port *InputPortByte) write(value interface{}) error {
	port.Mutex.Lock()
	defer port.Mutex.Unlock()

	typeOfValue := reflect.TypeOf(value)
	typeOfPortValue := reflect.TypeOf(port.Value)
	if !typeOfValue.ConvertibleTo(typeOfPortValue) {
		return fmt.Errorf("incompatible value types, given type: %v, could not be converted to: %v", typeOfValue, typeOfPortValue)
	}
	valueOfValue := reflect.ValueOf(value)
	port.Value = valueOfValue.Convert(typeOfPortValue).Interface().(byte)
	return nil
}

func (port *InputPortByte) Read() byte {
	port.Mutex.RLock()
	defer port.Mutex.RUnlock()

	return port.Value
}

type OutputPortInt struct {
	OutputPort
	Mutex sync.RWMutex
	Value int
}

type InputPortInt struct {
	InputPort
	Mutex sync.RWMutex
	Value int
}

func NewOutputPortInt() *OutputPortInt {
	return &OutputPortInt{Mutex: sync.RWMutex{}}
}

func NewInputPortInt(requiredNew bool) *InputPortInt {
	return &InputPortInt{Mutex: sync.RWMutex{}, InputPort: InputPort{RequiredNew: requiredNew}}
}

func (port *OutputPortInt) GetTimestamp() time.Time {
	return port.Timestamp
}

func (port *OutputPortInt) Write(value int) error {
	port.Mutex.Lock()
	defer port.Mutex.Unlock()

	port.Value = value
	port.Timestamp = time.Now()
	return nil
}

func (port *OutputPortInt) read() interface{} {
	port.Mutex.RLock()
	defer port.Mutex.RUnlock()

	return port.Value
}

func (port *InputPortInt) IsRequiredNew() bool {
	return port.RequiredNew
}

func (port *InputPortInt) write(value interface{}) error {
	port.Mutex.Lock()
	defer port.Mutex.Unlock()

	typeOfValue := reflect.TypeOf(value)
	typeOfPortValue := reflect.TypeOf(port.Value)
	if !typeOfValue.ConvertibleTo(typeOfPortValue) {
		return fmt.Errorf("incompatible value types, given type: %v, could not be converted to: %v", typeOfValue, typeOfPortValue)
	}
	valueOfValue := reflect.ValueOf(value)
	port.Value = valueOfValue.Convert(typeOfPortValue).Interface().(int)
	return nil
}

func (port *InputPortInt) Read() int {
	port.Mutex.RLock()
	defer port.Mutex.RUnlock()

	return port.Value
}

type OutputPortInt8 struct {
	OutputPort
	Mutex sync.RWMutex
	Value int8
}

type InputPortInt8 struct {
	InputPort
	Mutex sync.RWMutex
	Value int8
}

func NewOutputPortInt8() *OutputPortInt8 {
	return &OutputPortInt8{Mutex: sync.RWMutex{}}
}

func NewInputPortInt8(requiredNew bool) *InputPortInt8 {
	return &InputPortInt8{Mutex: sync.RWMutex{}, InputPort: InputPort{RequiredNew: requiredNew}}
}

func (port *OutputPortInt8) GetTimestamp() time.Time {
	return port.Timestamp
}

func (port *OutputPortInt8) Write(value int8) error {
	port.Mutex.Lock()
	defer port.Mutex.Unlock()

	port.Value = value
	port.Timestamp = time.Now()
	return nil
}

func (port *OutputPortInt8) read() interface{} {
	port.Mutex.RLock()
	defer port.Mutex.RUnlock()

	return port.Value
}

func (port *InputPortInt8) IsRequiredNew() bool {
	return port.RequiredNew
}

func (port *InputPortInt8) write(value interface{}) error {
	port.Mutex.Lock()
	defer port.Mutex.Unlock()

	typeOfValue := reflect.TypeOf(value)
	typeOfPortValue := reflect.TypeOf(port.Value)
	if !typeOfValue.ConvertibleTo(typeOfPortValue) {
		return fmt.Errorf("incompatible value types, given type: %v, could not be converted to: %v", typeOfValue, typeOfPortValue)
	}
	valueOfValue := reflect.ValueOf(value)
	port.Value = valueOfValue.Convert(typeOfPortValue).Interface().(int8)
	return nil
}

func (port *InputPortInt8) Read() int8 {
	port.Mutex.RLock()
	defer port.Mutex.RUnlock()

	return port.Value
}

type OutputPortInt16 struct {
	OutputPort
	Mutex sync.RWMutex
	Value int16
}

type InputPortInt16 struct {
	InputPort
	Mutex sync.RWMutex
	Value int16
}

func NewOutputPortInt16() *OutputPortInt16 {
	return &OutputPortInt16{Mutex: sync.RWMutex{}}
}

func NewInputPortInt16(requiredNew bool) *InputPortInt16 {
	return &InputPortInt16{Mutex: sync.RWMutex{}, InputPort: InputPort{RequiredNew: requiredNew}}
}

func (port *OutputPortInt16) GetTimestamp() time.Time {
	return port.Timestamp
}

func (port *OutputPortInt16) Write(value int16) error {
	port.Mutex.Lock()
	defer port.Mutex.Unlock()

	port.Value = value
	port.Timestamp = time.Now()
	return nil
}

func (port *OutputPortInt16) read() interface{} {
	port.Mutex.RLock()
	defer port.Mutex.RUnlock()

	return port.Value
}

func (port *InputPortInt16) IsRequiredNew() bool {
	return port.RequiredNew
}

func (port *InputPortInt16) write(value interface{}) error {
	port.Mutex.Lock()
	defer port.Mutex.Unlock()

	typeOfValue := reflect.TypeOf(value)
	typeOfPortValue := reflect.TypeOf(port.Value)
	if !typeOfValue.ConvertibleTo(typeOfPortValue) {
		return fmt.Errorf("incompatible value types, given type: %v, could not be converted to: %v", typeOfValue, typeOfPortValue)
	}
	valueOfValue := reflect.ValueOf(value)
	port.Value = valueOfValue.Convert(typeOfPortValue).Interface().(int16)
	return nil
}

func (port *InputPortInt16) Read() int16 {
	port.Mutex.RLock()
	defer port.Mutex.RUnlock()

	return port.Value
}

type OutputPortInt32 struct {
	OutputPort
	Mutex sync.RWMutex
	Value int32
}

type InputPortInt32 struct {
	InputPort
	Mutex sync.RWMutex
	Value int32
}

func NewOutputPortInt32() *OutputPortInt32 {
	return &OutputPortInt32{Mutex: sync.RWMutex{}}
}

func NewInputPortInt32(requiredNew bool) *InputPortInt32 {
	return &InputPortInt32{Mutex: sync.RWMutex{}, InputPort: InputPort{RequiredNew: requiredNew}}
}

func (port *OutputPortInt32) GetTimestamp() time.Time {
	return port.Timestamp
}

func (port *OutputPortInt32) Write(value int32) error {
	port.Mutex.Lock()
	defer port.Mutex.Unlock()

	port.Value = value
	port.Timestamp = time.Now()
	return nil
}

func (port *OutputPortInt32) read() interface{} {
	port.Mutex.RLock()
	defer port.Mutex.RUnlock()

	return port.Value
}

func (port *InputPortInt32) IsRequiredNew() bool {
	return port.RequiredNew
}

func (port *InputPortInt32) write(value interface{}) error {
	port.Mutex.Lock()
	defer port.Mutex.Unlock()

	typeOfValue := reflect.TypeOf(value)
	typeOfPortValue := reflect.TypeOf(port.Value)
	if !typeOfValue.ConvertibleTo(typeOfPortValue) {
		return fmt.Errorf("incompatible value types, given type: %v, could not be converted to: %v", typeOfValue, typeOfPortValue)
	}
	valueOfValue := reflect.ValueOf(value)
	port.Value = valueOfValue.Convert(typeOfPortValue).Interface().(int32)
	return nil
}

func (port *InputPortInt32) Read() int32 {
	port.Mutex.RLock()
	defer port.Mutex.RUnlock()

	return port.Value
}

type OutputPortInt64 struct {
	OutputPort
	Mutex sync.RWMutex
	Value int64
}

type InputPortInt64 struct {
	InputPort
	Mutex sync.RWMutex
	Value int64
}

func NewOutputPortInt64() *OutputPortInt64 {
	return &OutputPortInt64{Mutex: sync.RWMutex{}}
}

func NewInputPortInt64(requiredNew bool) *InputPortInt64 {
	return &InputPortInt64{Mutex: sync.RWMutex{}, InputPort: InputPort{RequiredNew: requiredNew}}
}

func (port *OutputPortInt64) GetTimestamp() time.Time {
	return port.Timestamp
}

func (port *OutputPortInt64) Write(value int64) error {
	port.Mutex.Lock()
	defer port.Mutex.Unlock()

	port.Value = value
	port.Timestamp = time.Now()
	return nil
}

func (port *OutputPortInt64) read() interface{} {
	port.Mutex.RLock()
	defer port.Mutex.RUnlock()

	return port.Value
}

func (port *InputPortInt64) IsRequiredNew() bool {
	return port.RequiredNew
}

func (port *InputPortInt64) write(value interface{}) error {
	port.Mutex.Lock()
	defer port.Mutex.Unlock()

	typeOfValue := reflect.TypeOf(value)
	typeOfPortValue := reflect.TypeOf(port.Value)
	if !typeOfValue.ConvertibleTo(typeOfPortValue) {
		return fmt.Errorf("incompatible value types, given type: %v, could not be converted to: %v", typeOfValue, typeOfPortValue)
	}
	valueOfValue := reflect.ValueOf(value)
	port.Value = valueOfValue.Convert(typeOfPortValue).Interface().(int64)
	return nil
}

func (port *InputPortInt64) Read() int64 {
	port.Mutex.RLock()
	defer port.Mutex.RUnlock()

	return port.Value
}

type OutputPortUint struct {
	OutputPort
	Mutex sync.RWMutex
	Value uint
}

type InputPortUint struct {
	InputPort
	Mutex sync.RWMutex
	Value uint
}

func NewOutputPortUint() *OutputPortUint {
	return &OutputPortUint{Mutex: sync.RWMutex{}}
}

func NewInputPortUint(requiredNew bool) *InputPortUint {
	return &InputPortUint{Mutex: sync.RWMutex{}, InputPort: InputPort{RequiredNew: requiredNew}}
}

func (port *OutputPortUint) GetTimestamp() time.Time {
	return port.Timestamp
}

func (port *OutputPortUint) Write(value uint) error {
	port.Mutex.Lock()
	defer port.Mutex.Unlock()

	port.Value = value
	port.Timestamp = time.Now()
	return nil
}

func (port *OutputPortUint) read() interface{} {
	port.Mutex.RLock()
	defer port.Mutex.RUnlock()

	return port.Value
}

func (port *InputPortUint) IsRequiredNew() bool {
	return port.RequiredNew
}

func (port *InputPortUint) write(value interface{}) error {
	port.Mutex.Lock()
	defer port.Mutex.Unlock()

	typeOfValue := reflect.TypeOf(value)
	typeOfPortValue := reflect.TypeOf(port.Value)
	if !typeOfValue.ConvertibleTo(typeOfPortValue) {
		return fmt.Errorf("incompatible value types, given type: %v, could not be converted to: %v", typeOfValue, typeOfPortValue)
	}
	valueOfValue := reflect.ValueOf(value)
	port.Value = valueOfValue.Convert(typeOfPortValue).Interface().(uint)
	return nil
}

func (port *InputPortUint) Read() uint {
	port.Mutex.RLock()
	defer port.Mutex.RUnlock()

	return port.Value
}

type OutputPortUint8 struct {
	OutputPort
	Mutex sync.RWMutex
	Value uint8
}

type InputPortUint8 struct {
	InputPort
	Mutex sync.RWMutex
	Value uint8
}

func NewOutputPortUint8() *OutputPortUint8 {
	return &OutputPortUint8{Mutex: sync.RWMutex{}}
}

func NewInputPortUint8(requiredNew bool) *InputPortUint8 {
	return &InputPortUint8{Mutex: sync.RWMutex{}, InputPort: InputPort{RequiredNew: requiredNew}}
}

func (port *OutputPortUint8) GetTimestamp() time.Time {
	return port.Timestamp
}

func (port *OutputPortUint8) Write(value uint8) error {
	port.Mutex.Lock()
	defer port.Mutex.Unlock()

	port.Value = value
	port.Timestamp = time.Now()
	return nil
}

func (port *OutputPortUint8) read() interface{} {
	port.Mutex.RLock()
	defer port.Mutex.RUnlock()

	return port.Value
}

func (port *InputPortUint8) IsRequiredNew() bool {
	return port.RequiredNew
}

func (port *InputPortUint8) write(value interface{}) error {
	port.Mutex.Lock()
	defer port.Mutex.Unlock()

	typeOfValue := reflect.TypeOf(value)
	typeOfPortValue := reflect.TypeOf(port.Value)
	if !typeOfValue.ConvertibleTo(typeOfPortValue) {
		return fmt.Errorf("incompatible value types, given type: %v, could not be converted to: %v", typeOfValue, typeOfPortValue)
	}
	valueOfValue := reflect.ValueOf(value)
	port.Value = valueOfValue.Convert(typeOfPortValue).Interface().(uint8)
	return nil
}

func (port *InputPortUint8) Read() uint8 {
	port.Mutex.RLock()
	defer port.Mutex.RUnlock()

	return port.Value
}

type OutputPortUint16 struct {
	OutputPort
	Mutex sync.RWMutex
	Value uint16
}

type InputPortUint16 struct {
	InputPort
	Mutex sync.RWMutex
	Value uint16
}

func NewOutputPortUint16() *OutputPortUint16 {
	return &OutputPortUint16{Mutex: sync.RWMutex{}}
}

func NewInputPortUint16(requiredNew bool) *InputPortUint16 {
	return &InputPortUint16{Mutex: sync.RWMutex{}, InputPort: InputPort{RequiredNew: requiredNew}}
}

func (port *OutputPortUint16) GetTimestamp() time.Time {
	return port.Timestamp
}

func (port *OutputPortUint16) Write(value uint16) error {
	port.Mutex.Lock()
	defer port.Mutex.Unlock()

	port.Value = value
	port.Timestamp = time.Now()
	return nil
}

func (port *OutputPortUint16) read() interface{} {
	port.Mutex.RLock()
	defer port.Mutex.RUnlock()

	return port.Value
}

func (port *InputPortUint16) IsRequiredNew() bool {
	return port.RequiredNew
}

func (port *InputPortUint16) write(value interface{}) error {
	port.Mutex.Lock()
	defer port.Mutex.Unlock()

	typeOfValue := reflect.TypeOf(value)
	typeOfPortValue := reflect.TypeOf(port.Value)
	if !typeOfValue.ConvertibleTo(typeOfPortValue) {
		return fmt.Errorf("incompatible value types, given type: %v, could not be converted to: %v", typeOfValue, typeOfPortValue)
	}
	valueOfValue := reflect.ValueOf(value)
	port.Value = valueOfValue.Convert(typeOfPortValue).Interface().(uint16)
	return nil
}

func (port *InputPortUint16) Read() uint16 {
	port.Mutex.RLock()
	defer port.Mutex.RUnlock()

	return port.Value
}

type OutputPortUint32 struct {
	OutputPort
	Mutex sync.RWMutex
	Value uint32
}

type InputPortUint32 struct {
	InputPort
	Mutex sync.RWMutex
	Value uint32
}

func NewOutputPortUint32() *OutputPortUint32 {
	return &OutputPortUint32{Mutex: sync.RWMutex{}}
}

func NewInputPortUint32(requiredNew bool) *InputPortUint32 {
	return &InputPortUint32{Mutex: sync.RWMutex{}, InputPort: InputPort{RequiredNew: requiredNew}}
}

func (port *OutputPortUint32) GetTimestamp() time.Time {
	return port.Timestamp
}

func (port *OutputPortUint32) Write(value uint32) error {
	port.Mutex.Lock()
	defer port.Mutex.Unlock()

	port.Value = value
	port.Timestamp = time.Now()
	return nil
}

func (port *OutputPortUint32) read() interface{} {
	port.Mutex.RLock()
	defer port.Mutex.RUnlock()

	return port.Value
}

func (port *InputPortUint32) IsRequiredNew() bool {
	return port.RequiredNew
}

func (port *InputPortUint32) write(value interface{}) error {
	port.Mutex.Lock()
	defer port.Mutex.Unlock()

	typeOfValue := reflect.TypeOf(value)
	typeOfPortValue := reflect.TypeOf(port.Value)
	if !typeOfValue.ConvertibleTo(typeOfPortValue) {
		return fmt.Errorf("incompatible value types, given type: %v, could not be converted to: %v", typeOfValue, typeOfPortValue)
	}
	valueOfValue := reflect.ValueOf(value)
	port.Value = valueOfValue.Convert(typeOfPortValue).Interface().(uint32)
	return nil
}

func (port *InputPortUint32) Read() uint32 {
	port.Mutex.RLock()
	defer port.Mutex.RUnlock()

	return port.Value
}

type OutputPortUint64 struct {
	OutputPort
	Mutex sync.RWMutex
	Value uint64
}

type InputPortUint64 struct {
	InputPort
	Mutex sync.RWMutex
	Value uint64
}

func NewOutputPortUint64() *OutputPortUint64 {
	return &OutputPortUint64{Mutex: sync.RWMutex{}}
}

func NewInputPortUint64(requiredNew bool) *InputPortUint64 {
	return &InputPortUint64{Mutex: sync.RWMutex{}, InputPort: InputPort{RequiredNew: requiredNew}}
}

func (port *OutputPortUint64) GetTimestamp() time.Time {
	return port.Timestamp
}

func (port *OutputPortUint64) Write(value uint64) error {
	port.Mutex.Lock()
	defer port.Mutex.Unlock()

	port.Value = value
	port.Timestamp = time.Now()
	return nil
}

func (port *OutputPortUint64) read() interface{} {
	port.Mutex.RLock()
	defer port.Mutex.RUnlock()

	return port.Value
}

func (port *InputPortUint64) IsRequiredNew() bool {
	return port.RequiredNew
}

func (port *InputPortUint64) write(value interface{}) error {
	port.Mutex.Lock()
	defer port.Mutex.Unlock()

	typeOfValue := reflect.TypeOf(value)
	typeOfPortValue := reflect.TypeOf(port.Value)
	if !typeOfValue.ConvertibleTo(typeOfPortValue) {
		return fmt.Errorf("incompatible value types, given type: %v, could not be converted to: %v", typeOfValue, typeOfPortValue)
	}
	valueOfValue := reflect.ValueOf(value)
	port.Value = valueOfValue.Convert(typeOfPortValue).Interface().(uint64)
	return nil
}

func (port *InputPortUint64) Read() uint64 {
	port.Mutex.RLock()
	defer port.Mutex.RUnlock()

	return port.Value
}

type OutputPortBool struct {
	OutputPort
	Mutex sync.RWMutex
	Value bool
}

type InputPortBool struct {
	InputPort
	Mutex sync.RWMutex
	Value bool
}

func NewOutputPortBool() *OutputPortBool {
	return &OutputPortBool{Mutex: sync.RWMutex{}}
}

func NewInputPortBool(requiredNew bool) *InputPortBool {
	return &InputPortBool{Mutex: sync.RWMutex{}, InputPort: InputPort{RequiredNew: requiredNew}}
}

func (port *OutputPortBool) GetTimestamp() time.Time {
	return port.Timestamp
}

func (port *OutputPortBool) Write(value bool) error {
	port.Mutex.Lock()
	defer port.Mutex.Unlock()

	port.Value = value
	port.Timestamp = time.Now()
	return nil
}

func (port *OutputPortBool) read() interface{} {
	port.Mutex.RLock()
	defer port.Mutex.RUnlock()

	return port.Value
}

func (port *InputPortBool) IsRequiredNew() bool {
	return port.RequiredNew
}

func (port *InputPortBool) write(value interface{}) error {
	port.Mutex.Lock()
	defer port.Mutex.Unlock()

	typeOfValue := reflect.TypeOf(value)
	typeOfPortValue := reflect.TypeOf(port.Value)
	if !typeOfValue.ConvertibleTo(typeOfPortValue) {
		return fmt.Errorf("incompatible value types, given type: %v, could not be converted to: %v", typeOfValue, typeOfPortValue)
	}
	valueOfValue := reflect.ValueOf(value)
	port.Value = valueOfValue.Convert(typeOfPortValue).Interface().(bool)
	return nil
}

func (port *InputPortBool) Read() bool {
	port.Mutex.RLock()
	defer port.Mutex.RUnlock()

	return port.Value
}

type OutputPortFloat32 struct {
	OutputPort
	Mutex sync.RWMutex
	Value float32
}

type InputPortFloat32 struct {
	InputPort
	Mutex sync.RWMutex
	Value float32
}

func NewOutputPortFloat32() *OutputPortFloat32 {
	return &OutputPortFloat32{Mutex: sync.RWMutex{}}
}

func NewInputPortFloat32(requiredNew bool) *InputPortFloat32 {
	return &InputPortFloat32{Mutex: sync.RWMutex{}, InputPort: InputPort{RequiredNew: requiredNew}}
}

func (port *OutputPortFloat32) GetTimestamp() time.Time {
	return port.Timestamp
}

func (port *OutputPortFloat32) Write(value float32) error {
	port.Mutex.Lock()
	defer port.Mutex.Unlock()

	port.Value = value
	port.Timestamp = time.Now()
	return nil
}

func (port *OutputPortFloat32) read() interface{} {
	port.Mutex.RLock()
	defer port.Mutex.RUnlock()

	return port.Value
}

func (port *InputPortFloat32) IsRequiredNew() bool {
	return port.RequiredNew
}

func (port *InputPortFloat32) write(value interface{}) error {
	port.Mutex.Lock()
	defer port.Mutex.Unlock()

	typeOfValue := reflect.TypeOf(value)
	typeOfPortValue := reflect.TypeOf(port.Value)
	if !typeOfValue.ConvertibleTo(typeOfPortValue) {
		return fmt.Errorf("incompatible value types, given type: %v, could not be converted to: %v", typeOfValue, typeOfPortValue)
	}
	valueOfValue := reflect.ValueOf(value)
	port.Value = valueOfValue.Convert(typeOfPortValue).Interface().(float32)
	return nil
}

func (port *InputPortFloat32) Read() float32 {
	port.Mutex.RLock()
	defer port.Mutex.RUnlock()

	return port.Value
}

type OutputPortFloat64 struct {
	OutputPort
	Mutex sync.RWMutex
	Value float64
}

type InputPortFloat64 struct {
	InputPort
	Mutex sync.RWMutex
	Value float64
}

func NewOutputPortFloat64() *OutputPortFloat64 {
	return &OutputPortFloat64{Mutex: sync.RWMutex{}}
}

func NewInputPortFloat64(requiredNew bool) *InputPortFloat64 {
	return &InputPortFloat64{Mutex: sync.RWMutex{}, InputPort: InputPort{RequiredNew: requiredNew}}
}

func (port *OutputPortFloat64) GetTimestamp() time.Time {
	return port.Timestamp
}

func (port *OutputPortFloat64) Write(value float64) error {
	port.Mutex.Lock()
	defer port.Mutex.Unlock()

	port.Value = value
	port.Timestamp = time.Now()
	return nil
}

func (port *OutputPortFloat64) read() interface{} {
	port.Mutex.RLock()
	defer port.Mutex.RUnlock()

	return port.Value
}

func (port *InputPortFloat64) IsRequiredNew() bool {
	return port.RequiredNew
}

func (port *InputPortFloat64) write(value interface{}) error {
	port.Mutex.Lock()
	defer port.Mutex.Unlock()

	typeOfValue := reflect.TypeOf(value)
	typeOfPortValue := reflect.TypeOf(port.Value)
	if !typeOfValue.ConvertibleTo(typeOfPortValue) {
		return fmt.Errorf("incompatible value types, given type: %v, could not be converted to: %v", typeOfValue, typeOfPortValue)
	}
	valueOfValue := reflect.ValueOf(value)
	port.Value = valueOfValue.Convert(typeOfPortValue).Interface().(float64)
	return nil
}

func (port *InputPortFloat64) Read() float64 {
	port.Mutex.RLock()
	defer port.Mutex.RUnlock()

	return port.Value
}
