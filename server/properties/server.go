package properties

import (
	"github.com/gorilla/mux"
	"github.com/mic90/go-flow/property"
	"net/http"
	"strconv"
)

type PropertiesServer struct {
	Port       uint64
	properties property.PropertyMapReaderWriter
}

func NewPropertiesServer(port uint64, properties property.PropertyMapReaderWriter) *PropertiesServer {
	return &PropertiesServer{port, properties}
}

func (server *PropertiesServer) StartAsync() {
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/property/{propertyName}", &getPropertyHandler{server.properties}).Methods("GET")
	router.Handle("/property/{propertyName}", &setPropertyHandler{server.properties}).Methods("POST")

	address := ":" + strconv.FormatUint(server.Port, 10)
	go http.ListenAndServe(address, router)
}
