package request

import (
	"encoding/json"
	"net/http"
)

type Request struct {
	http.Request
}

func NewRequest(r *http.Request) *Request {
	return &Request{
		Request: *r,
	}
}

func (r *Request) GetFormValue(field string) string {
	return r.Form.Get(field)
}

func (r *Request) UnmarshallBody(payload interface{}) error {
	return json.NewDecoder(r.Body).Decode(&payload)
}
