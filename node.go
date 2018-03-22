package goflow

import (
	"encoding/json"
	"errors"
)

type NodeProperty struct {
	Name     string
	Value    interface{}
	ReadOnly bool
}

type Node interface {
	GetName() string
	GetDescription() string
	GetID() string
	SetID(id string)

	ToJSONString() []byte

	GetInputs() map[string]*InputPort
	GetOutputs() map[string]*OutputPort

	Process() error
}

type BaseNode struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	ID          string                 `json:"id"`
	Inputs      map[string]*InputPort  `json:"inputs"`
	Outputs     map[string]*OutputPort `json:"outputs"`
	Properties  []NodeProperty
}

func (node *BaseNode) GetName() string {
	return node.Name
}

func (node *BaseNode) GetDescription() string {
	return node.Description
}

func (node *BaseNode) GetID() string {
	return node.ID
}

func (node *BaseNode) SetID(id string) {
	node.ID = id
}

func (node *BaseNode) GetInputs() map[string]*InputPort {
	return node.Inputs
}

func (node *BaseNode) GetOutputs() map[string]*OutputPort {
	return node.Outputs
}

func (node *BaseNode) Process() error {
	return errors.New("override me")
}

func (node *BaseNode) ToJSONString() []byte {
	jsonString, err := json.Marshal(node)
	if err != nil {
		panic(err)
	}
	return jsonString
}
