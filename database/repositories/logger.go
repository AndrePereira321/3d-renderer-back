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
