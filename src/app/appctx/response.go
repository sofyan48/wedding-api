package appctx

import (
	"github.com/orn-id/wedding-api/src/pkg/msg"
)

// Response presentation contract object
type Response struct {
	Name    string      `json:"name"`
	Message interface{} `json:"message,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Lang    string      `json:"-"`
}

// GetCode method to transform response name var to http status
func (r *Response) GetCode() int {
	return msg.GetCode(r.Name)
}

// GetMessage method to transform response name var to message detail
func (r *Response) GetMessage() string {
	return msg.Get(r.Name, r.Lang, r.Message)
}

// SetMessage method to set message response for not default
func (r *Response) SetMessage() {
	r.Message = msg.Get(r.Name, r.Lang, r.Message)
}
