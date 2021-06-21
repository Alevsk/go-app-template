package pkg

import (
	"log"
	"os"
)

var infoLog = log.New(os.Stdout, "I: ", log.LstdFlags)
var errorLog = log.New(os.Stdout, "E: ", log.LstdFlags)

func logInfo(msg string, data ...interface{}) {
	infoLog.Printf(msg+"\n", data...)
}

func logError(msg string, data ...interface{}) {
	errorLog.Printf(msg+"\n", data...)
}

// globally changeable logger styles
var (
	LogInfo  = logInfo
	LogError = logError
)
