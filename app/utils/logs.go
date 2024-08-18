package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

var cli struct {
	logFile string
}

func InfoLog(msg string) {
	log.Printf("[INFO] %s", msg)
}

func ErrorLog(msg string) {
	log.Printf("[ERROR] %s", msg)
}

func ResponseLog(msg string) {
	log.Printf("[RESPONSE] %s", msg)
}

func doesFileExist(logFileName string) bool {
	_, error := os.Stat(logFileName)

	// check if error is "file not exists"
	if os.IsNotExist(error) {
		fmt.Printf("%v file does not exist\n", logFileName)
		return true
	}
	return true
}

func SetupLogs() {
	currentTime := time.Now().Format("01-02-2006")
	logFileName := "./logs/" + currentTime + ".log"

	if doesFileExist(logFileName) {
		return
	}

	file, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	// Set config for logger
	log.SetOutput(&lumberjack.Logger{
		Filename:   logFileName,
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     1,     //days
		Compress:   false, // disabled by default
	})
}
