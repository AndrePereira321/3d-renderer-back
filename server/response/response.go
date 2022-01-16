package response

import "net/http"

type Response struct {
	http.ResponseWriter
}

func NewResponse(w http.ResponseWriter) *Response {
	return &Response{
		ResponseWriter: w,
	}
}
