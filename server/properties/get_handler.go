package properties

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mic90/go-flow/property"
	"net/http"
	"strings"
)

type getPropertyHandler struct {
	properties property.PropertyMapReaderWriter
}

func (h *getPropertyHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	propertyName := vars["propertyName"]
	if propertyName == "" {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	propertyNameLower := strings.ToLower(propertyName)
	value, err := h.properties.Read(propertyNameLower)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		return
	}
	id, err := h.properties.GetID(propertyNameLower)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
	}
	response := valueResponse{propertyNameLower, value, *id}
	err = json.NewEncoder(resp).Encode(response)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
	}
}
