package goflow

import (
	"encoding/json"
	"fmt"
	"github.com/twinj/uuid"
	"log"
	"sync/atomic"
)

type Connector struct {
	ID         string
	FromID     string
	FromOutput string
	ToID       string
	ToInput    string
}

type Graph struct {
	Nodes      map[string]Node      `json:"nodes"`
	Connectors map[string]Connector `json:"connectors"`
	errorChan  chan error           `json:"-"`
	stop       atomic.Value         `json:"-"`
}

func NewGraph() *Graph {
	nodes := make(map[string]Node)
	connectors := make(map[string]Connector)
	errorChan := make(chan error)
	stop := atomic.Value{}
	stop.Store(false)
	return &Graph{nodes, connectors, errorChan, stop}
}

func (graph *Graph) AddNode(node Node) string {
	id := uuid.NewV4().String()
	node.SetID(id)
	log.Printf("Adding node: %s \n", node.ToJSONString())
	graph.Nodes[node.GetID()] = node

	return id
}

func (graph *Graph) AddNodeFromRegister(nodeName, nodeVersion string) string {
	nodeFactoryFunc, err := GetNode(nodeName, nodeVersion)
	if err != nil {
		panic(err)
	}
	id := uuid.NewV4().String()
	newNode := nodeFactoryFunc()
	newNode.SetID(id)
	log.Printf("Adding node: %s \n", newNode.ToJSONString())
	graph.Nodes[newNode.GetID()] = newNode

	return id
}

func (graph *Graph) GetNodesCount() int {
	return len(graph.Nodes)
}

func (graph *Graph) Connect(fromId, outputName, toId, inputName string) {
	connID := uuid.NewV4().String()
	graph.Nodes[fromId].GetOutputs()[outputName].ConnectWith(graph.Nodes[toId].GetInputs()[inputName])
	graph.Connectors[connID] = Connector{connID, fromId, outputName, toId, inputName}
}

func (graph *Graph) StartAsync() {
	graph.stop.Store(false)
	for key := range graph.Nodes {
		go graph.startNode(graph.Nodes[key])
	}
}

func (graph *Graph) StopAsync() {
	graph.stop.Store(true)
}

func (graph *Graph) IsRunning() bool {
	return graph.stop.Load().(bool)
}

func (graph *Graph) ToJSONString() []byte {
	jsonString, err := json.Marshal(graph)
	if err != nil {
		panic(err)
	}
	return jsonString
}

func (graph *Graph) stopNode(node Node) {
	outputs := node.GetOutputs()
	for key := range outputs {
		for i := range outputs[key].Connectors {
			close(outputs[key].Connectors[i])
		}
	}
}

func (graph *Graph) startNode(node Node) {
	var processingError error
	log.Println("Started node: ", node.GetID())
	for {
		// stop processing loop
		if graph.stop.Load().(bool) == true {
			graph.stopNode(node)
			break
		}

		inputs := node.GetInputs()
		for key := range inputs {
			input := inputs[key]
			if input.IsOptional {
				if input.Connector == nil {
					continue
				}
				// if value is optional it might be set to null nothing was provided to read
				select {
				case value := <-input.Connector:
					input.Value = value
				default:
					input.Value = nil
				}
			} else {
				if input.Connector == nil {
					panic(fmt.Sprintf("Required input '%s' is not connected at node %s", input.Name, node.GetID()))
				}
				// if value is non optional, block processing until value is ready to read
				input.Value = <-input.Connector
			}
		}

		// some nodes might wait for inputs to arrive, double check for stop condition then
		if graph.stop.Load().(bool) == true {
			graph.stopNode(node)
			break
		}

		// run node processing function, if error occurred stop whole graph
		processingError = node.Process()
		if processingError != nil {
			graph.stop.Store(true)
		}
	}
	log.Println("Stopped node: ", node.GetID())
}
