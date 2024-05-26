package main

import (
	"log"
	"os"
)

var (
	infoLog  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime)
)

func logInfo(message string) {
	infoLog.Println(message)
}

func logError(message string, err error) {
	errorLog.Printf("%s: %v", message, err)
}
