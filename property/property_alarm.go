package property

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

type PropertyAlarm struct {
	Description string        `json:"description"`
	Value       bool          `json:"value"`
	timestamp   time.Time     `json:"-"`
	mutex       *sync.RWMutex `json:"-"`
}

func NewPropertyAlarm(description string) *PropertyAlarm {
	return &PropertyAlarm{Description: description, Value: false, mutex: &sync.RWMutex{}}
}

func (alarm *PropertyAlarm) Read() interface{} {
	alarm.mutex.RLock()
	defer alarm.mutex.RUnlock()

	return alarm.Value
}

func (alarm *PropertyAlarm) Write(value interface{}) error {
	alarm.mutex.Lock()
	defer alarm.mutex.Unlock()

	typeOfValue := reflect.TypeOf(value)
	typeOfPortValue := reflect.TypeOf(alarm.Value)

	if typeOfValue != typeOfPortValue {
		return fmt.Errorf("incompatible value types, given type: %v, could not be converted to: %v", typeOfValue, typeOfPortValue)
	}

	alarm.Value = value.(bool)
	alarm.timestamp = time.Now()
	return nil
}

func (prop *PropertyAlarm) GetTimestamp() time.Time {
	prop.mutex.RLock()
	defer prop.mutex.RUnlock()

	return prop.timestamp
}

func (alarm *PropertyAlarm) ReadBool() bool {
	alarm.mutex.RLock()
	defer alarm.mutex.RUnlock()

	return alarm.Value
}

func (alarm *PropertyAlarm) WriteBool(value bool) {
	alarm.mutex.Lock()
	defer alarm.mutex.Unlock()

	alarm.Value = value
}
