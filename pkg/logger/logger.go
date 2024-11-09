package logger

import (
    "log"
    "os"
)

var (
    Logger *log.Logger
)

// InitLogger initializes the logger for the application
func InitLogger() {
    file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("Failed to open log file: %v", err)
    }
    Logger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// LogInfo logs informational messages
func LogInfo(message string) {
    if Logger != nil {
        Logger.Println(message)
    } else {
        log.Println("INFO: " + message)
    }
}