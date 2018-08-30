package port

import (
	"github.com/mic90/go-flow/property"
	"github.com/twinj/uuid"
	"time"
)

type Connector interface {
	IsInputBlocking() bool
	IsOutputNew() bool
	GetID() string
	Trigger() error
}

type BaseConnector struct {
	ID        string
	Timestamp time.Time
}

func (c BaseConnector) GetID() string {
	return c.ID
}

type PortConnector struct {
	BaseConnector
	FromPort PortWriter
	ToPort   PortReader
}

func NewPortConnector(fromPort PortWriter, toPort PortReader) *PortConnector {
	id := uuid.NewV4().String()
	timestamp := time.Now()
	return &PortConnector{BaseConnector{id, timestamp}, fromPort, toPort}
}

func (p *PortConnector) IsInputBlocking() bool {
	return p.ToPort.IsRequiredNew()
}

func (p *PortConnector) IsOutputNew() bool {
	outputNew := p.FromPort.GetTimestamp().After(p.Timestamp)
	if outputNew {
		p.Timestamp = p.FromPort.GetTimestamp()
	}
	return outputNew
}

func (p *PortConnector) Trigger() error {
	return p.ToPort.write(p.FromPort.read())
}

type PropertyConnector struct {
	BaseConnector
	FromProperty property.PropertyReader
	ToPort       PortReader
}

func NewPropertyConnector(property property.PropertyReader, toPort PortReader) *PropertyConnector {
	id := uuid.NewV4().String()
	timestamp := time.Now()
	return &PropertyConnector{BaseConnector{id, timestamp}, property, toPort}
}

func (p *PropertyConnector) IsInputBlocking() bool {
	return p.ToPort.IsRequiredNew()
}

func (p *PropertyConnector) IsOutputNew() bool {
	propertyNew := p.FromProperty.(property.PropertyTimestampedReader).GetTimestamp().After(p.Timestamp)
	if propertyNew {
		p.Timestamp = p.FromProperty.(property.PropertyTimestampedReader).GetTimestamp()
	}
	return propertyNew
}

func (p *PropertyConnector) Trigger() error {
	return p.ToPort.write(p.FromProperty.Read())
}
