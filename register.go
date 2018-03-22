package goflow

import (
	"errors"
	"fmt"
)

var nodesRegister map[nodeRegisterMember]func() Node

type nodeRegisterMember struct {
	Name    string
	Version string
}

func init() {
	nodesRegister = make(map[nodeRegisterMember]func() Node)
}

// RegisterNode registers new node factory with given name and version
// If node with given name and version already exists it will be overriden
// This is global function that is not thread-safe
func RegisterNode(nodeName, version string, factoryMethod func() Node) {
	nodesRegister[nodeRegisterMember{nodeName, version}] = factoryMethod
}

// GetNode returns node factory function which could be used to create new node
// If no node was found nil will be returned
// This is global function that is not thread-safe
func GetNode(nodeName, version string) (func() Node, error) {
	val, ok := nodesRegister[nodeRegisterMember{nodeName, version}]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Couldn't find given node %s:%s", nodeName, version))
	}
	return val, nil
}

// GetRegisteredNodes returns list of all registered nodes names with its versions
// example output: [ node:1.0.0, second_node:1.0.0 ] etc.
func GetRegisteredNodes() []string {
	names := make([]string, 0, len(nodesRegister))
	for k := range nodesRegister {
		names = append(names, fmt.Sprintf("%s:%s", k.Name, k.Version))
	}
	return names
}
