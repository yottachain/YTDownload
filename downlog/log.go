package downlog

import (
	"log"
	"os"
)

func GetLog(s string) *log.Logger {
	fileName := s + ".log"
	logFile, _ := os.Create(fileName)
	//logFile, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	debugLog := log.New(logFile, "", log.Ldate|log.Ltime)
	return debugLog
}
