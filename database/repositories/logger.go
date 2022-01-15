package repositories

import (
	"server/database"
	"time"
)

type LogMessageDTO struct {
	database.DTO `bson:"inline"`
	Level        string
	Message      string
	Detail       string
	Location     string
	Time         string
}

func NewLogMessageDTO(logLevel, msg, location, detail string) LogMessageDTO {
	return LogMessageDTO{
		DTO: database.DTO{
			Name: "logMessages",
		},
		Level:    logLevel,
		Message:  msg,
		Location: location,
		Detail:   detail,
		Time:     time.Now().String(),
	}
}