package routes

import (
	"net/http"
	"server/database/repositories"
	"server/globals"
	"server/logger"
)

type DisconnectPayload struct {
	SessionCode string `json:"sessionCode"`
}

type DisconnectResponse struct {
	Status bool `json:"status"`
}

func Disconnect(route *Route) {
	payload := DisconnectPayload{}
	err := route.Request.UnmarshallBody(&payload)
	if err != nil {
		route.ResponseWriter.WriteError(http.StatusBadRequest, globals.ERROR_INVALID_PAYLOAD)
		return
	}

	err = repositories.DisableUserSession(payload.SessionCode)
	if err != nil {
		go logger.LogError("Error disabling user session: "+err.Error(), "Routes.Disconnect", globals.ERROR_DATABASE_ERROR)
		route.ResponseWriter.WriteErrorMessage(http.StatusInternalServerError, globals.ERROR_DATABASE_ERROR, "Error disabling user session: "+err.Error())
		return
	}

	route.Response.Data = DisconnectResponse{
		Status: true,
	}
	route.WriteData()
}
