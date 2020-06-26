package utils

import (
	"io/ioutil"
	"log"
	"os"
)

var logDir = "resources/info.log"

var (
	LogTrace   *log.Logger
	LogInfo    *log.Logger
	LogWarning *log.Logger
	LogError   *log.Logger
)

func init() {

	// Logger logFile output stuff.
	logFile, err := os.OpenFile(logDir, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	LogTrace = log.New(ioutil.Discard,
		"TRACE: ",
		log.Ldate|log.Ltime)

	LogInfo = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime)

	LogWarning = log.New(os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime)

	LogError = log.New(os.Stderr,
		"ERROR: ",
		log.Ldate|log.Ltime)

	// The errors will go to the log file for review on failures.
	// TODO: Add a rotating file handler to handle max size overflow later...
	LogError.SetOutput(logFile)
}
