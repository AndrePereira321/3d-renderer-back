package routes

import (
	"net/http"
	"server/database/repositories"
	"server/globals"
	"server/logger"
	"server/utils"
)

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	Email       string `json:"email"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
}

type LoginResponse struct {
	User        LoginUserResponse `json:"user"`
	SessionCode string            `json:"sessionCode"`
}

func Login(route *Route) {
	payload := LoginPayload{}
	err := route.Request.UnmarshallBody(&payload)
	if err != nil {
		route.ResponseWriter.WriteError(http.StatusBadRequest, globals.ERROR_INVALID_PAYLOAD)
		return
	}

	//TODO Retrieve activated account
	user, err := repositories.GetUserByEmail(payload.Email)
	if err != nil {
		route.ResponseWriter.WriteError(http.StatusForbidden, globals.ERROR_INEXISTENT_USER)
		return
	}

	if utils.HashPassword(payload.Password+user.Credentials.Salt) != user.Credentials.Hash {
		route.ResponseWriter.WriteError(http.StatusBadRequest, globals.ERROR_INVALID_PASSWORD)
		return
	}

	err = repositories.DisableUserSessions(payload.Email)
	if err != nil {
		go logger.LogError("Error disabling user sessions: "+err.Error(), "Routes.Login", globals.ERROR_DATABASE_ERROR)
		route.ResponseWriter.WriteErrorMessage(http.StatusInternalServerError, globals.ERROR_DATABASE_ERROR, "Error disabling user sessions: "+err.Error())
		return
	}

	sessionDTO := repositories.NewSessionDTOFill(payload.Email)
	_, err = sessionDTO.Save(sessionDTO)
	if err != nil {
		go logger.LogError("Error saving new session: "+err.Error(), "Routes.Login", globals.ERROR_DATABASE_ERROR)
		route.ResponseWriter.WriteErrorMessage(http.StatusBadRequest, globals.ERROR_DATABASE_ERROR, "Error saving new session: "+err.Error())
		return
	}

	response := LoginResponse{
		User: LoginUserResponse{
			Email:       user.Email,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			PhoneNumber: user.PhoneNumber,
		},
		SessionCode: sessionDTO.SessionCode,
	}

	route.Response.Data = response
	route.WriteData()
}
