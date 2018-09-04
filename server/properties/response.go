package properties

import "github.com/rs/xid"

type valueResponse struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
	ID    xid.ID      `json:"id"`
}
