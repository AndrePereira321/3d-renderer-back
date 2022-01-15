package logger

import (
	"fmt"
	"server/database/repositories"
	"server/globals"
	"time"
)

func displayLogMessage(logDTO repositories.LogMessageDTO) {
	fmt.Println(logDTO.Time + " - [" + logDTO.Level + "] - " + logDTO.Message)
}

func DisplayMessage(level, msg string) {
	fmt.Println(time.Now().String() + " - [" + level + "] - " + msg)
}

func Debug(msg string) {
	LogDebug(msg, "", "")
}

func LogDebug(msg, location, detail string) {
	if globals.LOG_DEBUG || globals.LOG_ALL_CONSOLE {
		logDTO := repositories.NewLogMessageDTO("DEBUG", msg, location, detail)
		if globals.LOG_DEBUG || globals.LOG_ALL_CONSOLE {
			displayLogMessage(logDTO)
		}
		if globals.LOG_DEBUG && globals.LOG_PERSISTENT {
			_, err := logDTO.Save(logDTO)
			if err != nil {
				DisplayMessage("CRITICAL ERROR", "Error saving log message: "+err.Error())
			}
		}
	}
}

func Warning(msg string) {
	LogWarning(msg, "", "")
}

func LogWarning(msg, location, detail string) {
	if globals.LOG_WARNING || globals.LOG_ALL_CONSOLE {
		logDTO := repositories.NewLogMessageDTO("WARNING", msg, location, detail)
		if globals.LOG_WARNING || globals.LOG_ALL_CONSOLE {
			displayLogMessage(logDTO)
		}
		if globals.LOG_WARNING && globals.LOG_PERSISTENT {
			_, err := logDTO.Save(logDTO)
			if err != nil {
				DisplayMessage("CRITICAL ERROR", "Error saving log message: "+err.Error())
			}
		}
	}
}

func Error(msg string) {
	LogError(msg, "", "")
}

func LogError(msg, location, detail string) {
	if globals.LOG_ERROR || globals.LOG_ALL_CONSOLE {
		logDTO := repositories.NewLogMessageDTO("ERROR", msg, location, detail)
		if globals.LOG_ERROR || globals.LOG_ALL_CONSOLE {
			displayLogMessage(logDTO)
		}
		if globals.LOG_ERROR && globals.LOG_PERSISTENT {
			_, err := logDTO.Save(logDTO)
			if err != nil {
				DisplayMessage("CRITICAL ERROR", "Error saving log message: "+err.Error())
			}
		}
	}
}
