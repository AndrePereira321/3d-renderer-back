package routes

import (
	"net/http"
	"server/database/repositories"
	"server/globals"
)

type IsConnectedPayload struct {
	SessionCode string `json:"sessionCode"`
}

type IsConnectedResponse struct {
	Active bool `json:"active"`
}

func IsConnected(route *Route) {
	payload := IsConnectedPayload{}
	err := route.Request.UnmarshallBody(&payload)
	if err != nil {
		route.ResponseWriter.WriteError(http.StatusBadRequest, globals.ERROR_INVALID_PAYLOAD)
		return
	}

	active, err := repositories.IsActiveSession(payload.SessionCode)
	if err != nil {
		route.ResponseWriter.WriteErrorMessage(http.StatusInternalServerError, globals.ERROR_DATABASE_ERROR, "Error retrieving user session: "+err.Error())
		return
	}

	response := IsConnectedResponse{
		Active: active,
	}

	route.Response.Data = response
	route.WriteData()
}
