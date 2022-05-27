package routes

import (
	"net/http"
	"server/database"
	"server/database/repositories"
	"server/globals"
	"server/logger"
	"server/utils"
)

type RegisterPayload struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
}

type RegisterResponse struct {
	Email string `json:"email"`
}

func Register(route *Route) {
	payload := RegisterPayload{}
	err := route.Request.UnmarshallBody(&payload)
	if err != nil {
		route.ResponseWriter.WriteError(http.StatusBadRequest, globals.ERROR_INVALID_PAYLOAD)
		return
	}

	user, _ := repositories.GetUserByEmail(payload.Email)
	if user != nil {
		route.ResponseWriter.WriteError(http.StatusBadRequest, globals.ERROR_EXISTENT_USER)
		return
	}

	if payload.Password != payload.PasswordConfirm {
		route.ResponseWriter.WriteError(http.StatusBadRequest, globals.ERROR_DIFFERENT_PASSWORD)
		return
	}

	if len(payload.Password) < 8 {
		route.ResponseWriter.WriteError(http.StatusBadRequest, globals.ERROR_PASSWORD_MIN_LENGTH)
		return
	}

	salt := utils.SaltFromPassword(payload.Password)

	//TODO send verification email
	dto := repositories.UserDTO{
		Email:     payload.Email,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Credentials: repositories.UserCredentials{
			Hash: utils.HashPassword(payload.Password + salt),
			Salt: salt,
		},
		DTO: database.DTO{
			CollectionName: "Users",
		},
		Verified: false,
		Active:   true,
	}

	_, err = dto.Save(dto)

	if err != nil {
		go logger.LogError("Error saving new user: "+err.Error(), "Routes.Register", globals.ERROR_DATABASE_ERROR)
		route.ResponseWriter.WriteErrorMessage(http.StatusInternalServerError, globals.ERROR_DATABASE_ERROR, err.Error())
		return
	}

	route.Response.Data = RegisterResponse{
		Email: payload.Email,
	}

	route.WriteData()
}
