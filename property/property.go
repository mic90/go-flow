package property

import "time"

type PropertyTimestampedReader interface {
	PropertyReader
	GetTimestamp() time.Time
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

type PropertyTimestampReaderWriter interface {
	PropertyTimestampedReader
	PropertyWriter
}
