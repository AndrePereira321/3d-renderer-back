package routes

import (
	"net/http"
	"server/globals"
	"server/logger"
	"server/server/request"
	"server/server/response"
)

type Route struct {
	Path           string
	Method         string
	Response       *response.Response
	ResponseWriter *response.ResponseWriter
	Request        *request.Request
}

func NewRoute(w http.ResponseWriter, r *http.Request) *Route {
	return &Route{
		Path:           r.URL.Path,
		Method:         r.Method,
		Response:       response.NewResponse(),
		ResponseWriter: response.NewResponseWriter(w),
		Request:        request.NewRequest(r),
	}
}

func (r *Route) WriteData() {
	err := r.ResponseWriter.WriteJSON(r.Response)
	if err != nil {
		go logger.LogError("Error writing JSON data in route"+r.Path, "Route.WriteData", globals.ERROR_RESPONSE_JSON_CONVERSION)
		r.ResponseWriter.WriteError(http.StatusInternalServerError, globals.ERROR_RESPONSE_JSON_CONVERSION)
	}
}
