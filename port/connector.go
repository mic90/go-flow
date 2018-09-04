package port

import (
	"github.com/mic90/go-flow/property"
	"github.com/rs/xid"
	"github.com/twinj/uuid"
)

type Connector interface {
	IsInputBlocking() bool
	IsInputBlockingDiff() bool
	IsOutputNew() bool
	IsInputDiff() bool
	GetID() string
	Trigger() error
}

type BaseConnector struct {
	ID       string
	OutputID xid.ID
}

func (c BaseConnector) GetID() string {
	return c.ID
}

type PortConnector struct {
	BaseConnector
	FromPort OutputPort
	ToPort   InputPort
}

func NewPortConnector(fromPort OutputPort, toPort InputPort) *PortConnector {
	id := uuid.NewV4().String()
	timestamp := xid.ID{}
	return &PortConnector{BaseConnector{id, timestamp}, fromPort, toPort}
}

func (p *PortConnector) IsInputBlocking() bool {
	return p.ToPort.IsBlockingNew()
}

func (p *PortConnector) IsInputBlockingDiff() bool {
	return p.ToPort.IsBlockingDiff()
}

func (p *PortConnector) IsOutputNew() bool {
	outputNew := p.FromPort.GetID().Compare(p.OutputID) != 0
	if outputNew {
		p.OutputID = p.FromPort.GetID()
	}
	return outputNew
}

func (p *PortConnector) IsInputDiff() bool {
	return p.ToPort.ValueChanged()
}

func (p *PortConnector) Trigger() error {
	return p.ToPort.write(p.FromPort.read())
}

type PropertyConnector struct {
	BaseConnector
	FromProperty property.PropertyReader
	ToPort       InputPort
}

func NewPropertyConnector(property property.PropertyReader, toPort InputPort) *PropertyConnector {
	id := uuid.NewV4().String()
	propertyID := xid.ID{}
	return &PropertyConnector{BaseConnector{id, propertyID}, property, toPort}
}

func (p *PropertyConnector) IsInputBlocking() bool {
	return p.ToPort.IsBlockingNew()
}

func (p *PropertyConnector) IsInputBlockingDiff() bool {
	return p.ToPort.IsBlockingDiff()
}

func (p *PropertyConnector) IsOutputNew() bool {
	propertyNew := p.FromProperty.(property.PropertyIDReader).GetID().Compare(p.OutputID) != 0
	if propertyNew {
		p.OutputID = p.FromProperty.(property.PropertyIDReader).GetID()
	}
	return propertyNew
}

func (p *PropertyConnector) IsInputDiff() bool {
	return p.ToPort.ValueChanged()
}

func (p *PropertyConnector) Trigger() error {
	return p.ToPort.write(p.FromProperty.Read())
}
