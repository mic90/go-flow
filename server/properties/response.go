package properties

type valueResponse struct {
	Name       string      `json:"name"`
	Value      interface{} `json:"value"`
	LastChange string      `json:"lastchange"`
}