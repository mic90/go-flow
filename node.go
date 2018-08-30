package goflow

import (
	"encoding/json"
	"errors"
	"github.com/mic90/go-flow/port"
)

//Node is the base struct used in the flow.
//Each node consists of inputs, outputs and its properties.
//The basic flow is that, when the inputs are received from connected nodes
//the node logic is executed.
//One must set all output values in the Process function to propagate its values to other nodes inputs
type Node interface {
	GetName() string
	GetDescription() string
	GetVersion() string

	GetID() string
	SetID(id string)

	ToJSONString() []byte

	AddConnector(id string)
	GetConnectors() []string

	Setup() error
	Process() error
}

//BaseNode is the base struct which must be added by composition to the user-defined nodes
type BaseNode struct {
	Name          string                     `json:"name"`
	Description   string                     `json:"description"`
	Version       string                     `json:"version"`
	ID            string                     `json:"id"`
	Inputs        map[string]port.PortReader `json:"inputs"`
	Outputs       map[string]port.PortWriter `json:"outputs"`
	ConnectorsIDs []string                   `json:"connectors"`
}

func (node *BaseNode) NewNodeFromJSONString(data string) (error, *BaseNode) {
	nodeData := BaseNode{}
	if err := json.Unmarshal([]byte(data), &nodeData); err != nil {
		return err, nil
	}
	return nil, &nodeData
}

//GetName returns the name of the node.
//The name should not contain spaces, be short and meaningfull
func (node *BaseNode) GetName() string {
	return node.Name
}

//GetDescription returns the description of the node
//The descrption could be as long as one would like.
//Provide all important informations here
func (node *BaseNode) GetDescription() string {
	return node.Description
}

//GetVersion returns version of given node
func (node *BaseNode) GetVersion() string {
	return node.Version
}

//GetID returns node id. The id will be generated automatically by graph when the node is added
func (node *BaseNode) GetID() string {
	return node.ID
}

//SetID sets the node id. This will be used by the graph
func (node *BaseNode) SetID(id string) {
	node.ID = id
}

func (node *BaseNode) AddConnector(id string) {
	node.ConnectorsIDs = append(node.ConnectorsIDs, id)
}

func (node *BaseNode) GetConnectors() []string {
	return node.ConnectorsIDs
}

//Setup this function will be triggered when the node is started.
//It will be triggred only once. Put all the initialization code here
func (node *BaseNode) Setup() error {
	return errors.New("override me")
}

//Process this is the main node function.
//It will be run in the loop, when this function is executed all non-optional inputs are already provided
//Put all the node logic here.
func (node *BaseNode) Process() error {
	return errors.New("override me")
}

//ToJSONString retruns json formatted string of the node
func (node *BaseNode) ToJSONString() []byte {
	jsonString, err := json.Marshal(node)
	if err != nil {
		panic(err)
	}
	return jsonString
}
