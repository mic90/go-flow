package goflow

import (
	"encoding/json"
	"errors"
	"github.com/mic90/go-flow/port"
	"github.com/mic90/go-flow/property"
	"log"
	"reflect"
	"sync"
	"sync/atomic"
	"time"

	"github.com/twinj/uuid"
)

type GraphStatusReader interface {
	IsRunning() bool
	ErrorString() string
}

type GraphNodesReader interface {
	GetNodesCount() int
	GetNodesJSON() []byte
	GetConnectorsJSON() []byte
}

type GraphReader interface {
	GraphStatusReader
	GraphNodesReader
}

type Graph struct {
	Nodes      map[string]Node      `json:"nodes"`
	Connectors map[string]port.Connector `json:"connectors"`
	Stop       atomic.Value         `json:"-"`
	Error      atomic.Value         `json:"-"`
	WaitGroup  sync.WaitGroup       `json:"-"`
}

func NewGraph() *Graph {
	nodes := make(map[string]Node)
	connectors := make(map[string]port.Connector)
	stop := atomic.Value{}
	stop.Store(false)
	err := atomic.Value{}
	var wg sync.WaitGroup
	return &Graph{nodes, connectors, stop, err, wg}
}

func (graph *Graph) AddNode(node Node) string {
	id := uuid.NewV4().String()
	node.SetID(id)
	log.Printf("Adding node: %s | %s | %s\n", node.GetName(), node.GetVersion(), id)
	graph.Nodes[node.GetID()] = node

	return id
}

func (graph *Graph) AddNodeFromRegister(nodeName, nodeVersion string, parameters...interface{}) string {
	nodeFactoryFunc, err := GetNode(nodeName, nodeVersion)
	if err != nil {
		panic(err)
	}
	id := uuid.NewV4().String()
	newNode := nodeFactoryFunc(parameters...)
	newNode.SetID(id)
	log.Printf("Adding node: %s | %s | %s\n", newNode.GetName(), newNode.GetVersion(), id)
	graph.Nodes[newNode.GetID()] = newNode

	return id
}

func (graph *Graph) GetNodesCount() int {
	return len(graph.Nodes)
}

func (graph *Graph) Connect(fromId, outputName, toId, inputName string) {
	fromPort, err := graph.getNodeOutputPort(fromId, outputName)
	if err != nil {
		panic(err)
	}
	toPort, err := graph.getNodeInputPort(toId, inputName)
	if err != nil {
		panic(err)
	}
	connector := port.NewPortConnector(fromPort, toPort)
	graph.Nodes[toId].AddConnector(connector.ID)
	graph.Connectors[connector.ID] = connector
}

func (graph *Graph) ConnectProperty(property property.PropertyReader, toId, inputName string) {
	toPort, err := graph.getNodeInputPort(toId, inputName)
	if err != nil {
		panic(err)
	}
	connector := port.NewPropertyConnector(property, toPort)
	graph.Nodes[toId].AddConnector(connector.ID)
	graph.Connectors[connector.ID] = connector
}

func (graph *Graph) StartAsync() {
	graph.Stop.Store(false)
	for key := range graph.Nodes {
		graph.WaitGroup.Add(1)
		go graph.startNode(graph.Nodes[key])
	}
}

func (graph *Graph) StopAndWait() {
	graph.Stop.Store(true)
	graph.WaitGroup.Wait()
}

func (graph *Graph) StopAsync() {
	graph.Stop.Store(true)
}

func (graph *Graph) IsRunning() bool {
	return !graph.Stop.Load().(bool)
}

func (graph *Graph) ErrorString() string {
	return graph.Error.Load().(error).Error()
}

func (graph *Graph) GetNodesJSON() []byte {
	jsonString, err := json.Marshal(graph.Nodes)
	if err != nil {
		panic(err)
	}
	return jsonString
}

func (graph *Graph) GetConnectorsJSON() []byte {
	jsonString, err := json.Marshal(graph.Connectors)
	if err != nil {
		panic(err)
	}
	return jsonString
}

func (graph *Graph) setError(err error) {
	log.Println("Graph error occured:", err)
	graph.Error.Store(err)
	graph.Stop.Store(true)
}

func (graph *Graph) startNode(node Node) {
	defer graph.WaitGroup.Done()

	setupError := node.Setup()
	if setupError != nil {
		log.Println("Couldn't start node", node.GetName())
		graph.setError(setupError)
		return
	}

	log.Println("Started node:", node.GetName(), node.GetID())
	for {
		// stop processing loop
		if graph.Stop.Load().(bool) == true {
			break
		}

		for _, connID := range node.GetConnectors() {
			conn := graph.Connectors[connID]

			// if inputs is required to be new before processing is run
			// wait for the writer node to write some new value to its output port
			if conn.IsInputBlocking() {
				for {
					if conn.IsOutputNew() {
						break
					}
					if graph.Stop.Load().(bool) == true {
						break
					}
					time.Sleep(1 * time.Millisecond)
				}
			}

			err := conn.Trigger()
			if err != nil {
				graph.setError(err)
				break
			}
		}

		// stop processing loop, double check in case output->input write went wrong
		if graph.Stop.Load().(bool) == true {
			break
		}

		// run node processing function, if error occurred stop whole graph
		processingError := node.Process()
		if processingError != nil {
			graph.setError(processingError)
		}
	}
	log.Println("Stopped node:", node.GetName(), node.GetID())
}

func (graph *Graph) getNodeInputPort(nodeId, portName string) (port.PortReader, error) {
	readerType := reflect.TypeOf((*port.PortReader)(nil)).Elem()
	node := graph.Nodes[nodeId]

	// read node ports into map by name
	nodeValue := reflect.ValueOf(node).Elem()
	for i := 0; i < nodeValue.NumField(); i++ {
		currentPortName := nodeValue.Type().Field(i).Name
		if currentPortName != portName {
			continue
		}
		isReader := nodeValue.Field(i).Type().Implements(readerType)
		if isReader {
			return nodeValue.Field(i).Interface().(port.PortReader), nil
		}
	}
	return nil, errors.New("unable to find input port")
}

func (graph *Graph) getNodeOutputPort(nodeId, portName string) (port.PortWriter, error) {
	writerType := reflect.TypeOf((*port.PortWriter)(nil)).Elem()
	node := graph.Nodes[nodeId]

	// read node ports into map by name
	nodeValue := reflect.ValueOf(node).Elem()
	for i := 0; i < nodeValue.NumField(); i++ {
		currentPortName := nodeValue.Type().Field(i).Name
		if currentPortName != portName {
			continue
		}
		isWriter := nodeValue.Field(i).Type().Implements(writerType)
		if isWriter {
			return nodeValue.Field(i).Interface().(port.PortWriter), nil
		}
		return nil, errors.New("given port is not an output port")
	}
	return nil, errors.New("unable to find output port")
}
