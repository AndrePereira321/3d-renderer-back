package repositories

import (
	"server/database"
	"time"
)

type LogMessageDTO struct {
	database.DTO
	Level      string `bson:"level" json:"level"`
	Message    string `bson:"message" json:"message"`
	Detail     string `bson:"detail" json:"detail"`
	Location   string `bson:"location" json:"location"`
	Time       string `bson:"time" json:"time"`
	StackTrace string `bson:"stackTrace" json:"stackTrace"`
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
