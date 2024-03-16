package logger

import (
	"log"
	"os"
)

func LogFatalExitIf(err error) {
	if err != nil {
		log.Fatalf(err.Error())
		os.Exit(1)
	}
}
