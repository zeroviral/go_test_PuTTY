package utils

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var logDir = "resources/logs"
var errorLogdir = "resources/logs/errors.log"
var infoLogDir = "resources/logs/info.log"

var (
	LogTrace   *log.Logger
	LogInfo    *log.Logger
	LogWarning *log.Logger
	LogError   *log.Logger
)

func Init() {

	// Logger logFile output stuff.
	if _, err := os.Stat("resources/logs"); os.IsNotExist(err) {
		os.Mkdir(logDir, 0700)
	}
	logErrorFile, err := os.OpenFile(errorLogdir, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	logInfoFile, err := os.OpenFile(infoLogDir, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	LogTrace = log.New(ioutil.Discard,
		"TRACE: ",
		log.Ldate|log.Ltime)

	LogInfo = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime)

	LogWarning = log.New(os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime)

	LogError = log.New(os.Stdout,
		"ERROR: ",
		log.Ldate|log.Ltime)

	// The errors will go to the log file for review on failures.
	// TODO: Add a rotating file handler to handle max size overflow later...
	// Output log errors to both STDOUT and file for review.
	multiWriterERROR := io.MultiWriter(os.Stdout, logErrorFile)
	multiWriterINFO := io.MultiWriter(os.Stdout, logInfoFile)

	LogError.SetOutput(multiWriterERROR)
	LogInfo.SetOutput(multiWriterINFO)
}
