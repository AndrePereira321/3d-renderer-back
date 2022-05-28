package repositories

import (
	"server/database"
	"time"
)

type LogMessageDTO struct {
	database.DTO
	Level      string `bson:"level"`
	Message    string `bson:"message"`
	Detail     string `bson:"detail"`
	Location   string `bson:"location"`
	Time       string `bson:"time"`
	StackTrace string `bson:"stackTrace"`
}

func NewLogMessageDTO(logLevel, msg, location, detail string) *LogMessageDTO {
	return &LogMessageDTO{
		DTO: database.DTO{
			CollectionName: "LogMessages",
		},
		Level:    logLevel,
		Message:  msg,
		Location: location,
		Detail:   detail,
		Time:     time.Now().Format(time.RFC1123),
	}
}
