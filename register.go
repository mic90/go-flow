package goflow

import (
	"fmt"
)

var nodesRegister map[NodeRegisterMember]func(...interface{}) Node

type NodeRegisterMember struct {
	Name    string
	Version string
}

func (member *NodeRegisterMember) GetName() string {
	return member.Name
}

func (member *NodeRegisterMember) GetVersion() string {
	return member.Version
}

func (member *NodeRegisterMember) GetCombined() string {
	return fmt.Sprintf("%s:%s", member.Name, member.Version)
}

func init() {
	nodesRegister = make(map[NodeRegisterMember]func(...interface{}) Node)
}

// RegisterNode registers new node factory with given name and version
// If node with given name and version already exists it will be overriden
// This is global function that is not thread-safe
func RegisterNode(nodeName, version string, factoryMethod func(...interface{}) Node) {
	nodesRegister[NodeRegisterMember{nodeName, version}] = factoryMethod
}

// GetNode returns node factory function which could be used to create new node
// If no node was found nil will be returned
// This is global function that is not thread-safe
func GetNode(nodeName, version string) (func(...interface{}) Node, error) {
	val, ok := nodesRegister[NodeRegisterMember{nodeName, version}]
	if !ok {
		return nil, fmt.Errorf("couldn't find given node %s:%s", nodeName, version)
	}
	return val, nil
}

// GetRegisteredNodes returns list of all registered nodes names with its versions
func GetRegisteredNodes() []NodeRegisterMember {
	members := make([]NodeRegisterMember, 0, len(nodesRegister))
	for k := range nodesRegister {
		members = append(members, k)
	}
	return members
}
