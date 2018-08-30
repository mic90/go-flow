package properties

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mic90/go-flow/property"
	"net/http"
	"strings"
)

type setValueRequest struct {
	Value      interface{} `json:"value"`
}

type setPropertyHandler struct {
	properties property.PropertyMapReaderWriter
}

func (h *setPropertyHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	propertyName := vars["propertyName"]
	if propertyName == "" {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	var request setValueRequest
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	propertyNameLower := strings.ToLower(propertyName)
	err = h.properties.Write(propertyNameLower, request.Value)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		return
	}
	timestamp, err := h.properties.GetTimestamp(propertyNameLower)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
	}
	response := valueResponse{propertyNameLower, request.Value, timestamp.Format("2006-01-02T15:04:05Z07:00")}
	err = json.NewEncoder(resp).Encode(response)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
	}
}