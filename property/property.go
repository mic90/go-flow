package property

import (
	"github.com/rs/xid"
)

type PropertyIDReader interface {
	PropertyReader
	GetID() xid.ID
}

type PropertyReader interface {
	Read() interface{}
}

type PropertyWriter interface {
	Write(interface{}) error
}

type PropertyReaderWriter interface {
	PropertyReader
	PropertyWriter
}

type PropertyIDReaderWriter interface {
	PropertyIDReader
	PropertyWriter
}
