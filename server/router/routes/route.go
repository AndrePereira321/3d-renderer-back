package routes

import (
	"net/http"
	"server/server/request"
	"server/server/response"
)

type Route struct {
	Path     string
	Method   string
	Response *response.Response
	Request  *request.Request
}

func NewRoute(w http.ResponseWriter, r *http.Request) *Route {
	return &Route{
		Path:     r.URL.Path,
		Method:   r.Method,
		Response: response.NewResponse(w),
		Request:  request.NewRequest(r),
	}
}
