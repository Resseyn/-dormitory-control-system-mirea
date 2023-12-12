package loggers

import (
	"log"
	"os"
)

var ErrorLogger *log.Logger

func init() {
	file, err := os.OpenFile("/-dormitory-control-system-mirea/logs/errorLogs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	ErrorLogger = log.New(file, "", log.LstdFlags)
}