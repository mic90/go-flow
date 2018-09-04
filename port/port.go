package port

import (
	"github.com/rs/xid"
)

type BlockingType int

const (
	PortBlockingNone BlockingType = iota
	PortBlockingNew
	PortBlockingDiff
)

type OutputPort interface {
	read() interface{}

	GetID() xid.ID
}

type InputPort interface {
	write(interface{}) error

	IsBlockingNew() bool
	IsBlockingDiff() bool

	ValueChanged() bool
	ValueNew() bool
}

type BaseOutputPort struct {
	id xid.ID
}

type BaseInputPort struct {
	blockingType BlockingType
	valueNew bool
}
