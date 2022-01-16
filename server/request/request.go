package request

import "net/http"

type Request struct {
	http.Request
}

func NewRequest(r *http.Request) *Request {
	return &Request{
		Request: *r,
	}
}
