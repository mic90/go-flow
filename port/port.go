package port

import (
	"time"
)

type PortWriter interface {
	read() interface{}
	GetTimestamp() time.Time
}

type PortReader interface {
	write(interface{}) error
	IsRequiredNew() bool
	ValueChanged() bool
}

type OutputPort struct {
	Timestamp time.Time `json:"timestamp"`
}

type InputPort struct {
	RequiredNew bool
}
