package goflow

import "fmt"

const (
	defaultConnectorsCap = 12
)

type InputPort struct {
	Name       string           `json:"name"`
	Value      interface{}      `json:"value"`
	Connector  chan interface{} `json:"-"`
	IsOptional bool             `json:"isOptional"`
}

func NewInputPort(name string, isOptional bool) *InputPort {
	return &InputPort{name, nil, nil, isOptional}
}

func (port *InputPort) Read() interface{} {
	return port.Value
}

type OutputPort struct {
	Name        string             `json:"name"`
	Connectors  []chan interface{} `json:"-"`
	InputsCount int                `json:"-"`
}

func NewOutputPort(name string) *OutputPort {
	connectors := make([]chan interface{}, 0, defaultConnectorsCap)
	return &OutputPort{name, connectors, 0}
}

// ConnectWith connects this port with given input port
// On connection a new channel is created and stored inside this port
// Channel is then passed to the input port where it can be used to retrieve values
func (port *OutputPort) ConnectWith(inPort *InputPort) {
	if inPort.Connector != nil {
		panic(fmt.Sprintf("Can't connect output %s to input %s. Input port alread connected", port.Name, inPort.Name))
	}
	newConnector := make(chan interface{})
	port.Connectors = append(port.Connectors, newConnector)
	inPort.Connector = newConnector
}

// Write writes given value into all connected ports
func (port *OutputPort) Write(value interface{}) {
	for i := range port.Connectors {
		select {
		case port.Connectors[i] <- value:
		default:
		}
	}
}
