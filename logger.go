package logger

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type loggerOptions struct {
	LogFileIsActive bool   `json:"logFileIsActive"` // If you need write to file
	LogFileName     string `json:"logFileName"`     // Filename that you want to write the logs
	LogLevel        string `json:"logLevel"`        // Log level that you want write
}

var logOptions loggerOptions

var (
	Debug *log.Logger

	// Info Default level to write output
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
)

// private variables to set output defined on setLevel caller from setLoggerOptions
var (
	debugLoggerOutput   *os.File
	infoLoggerOutput    *os.File
	warningLoggerOutput *os.File
	errorLoggerOutput   *os.File
)

// set level to write output
func setLevel() {
	logOptions.LogLevel = os.Getenv("LOG_LEVEL")

	switch strings.ToLower(logOptions.LogLevel) {
	case "debug", "debugger":
		if !logOptions.LogFileIsActive {
			debugLoggerOutput = os.Stdout
			infoLoggerOutput = os.Stdout
			warningLoggerOutput = os.Stdout
			errorLoggerOutput = os.Stderr
		}

		Debug = log.New(debugLoggerOutput, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
		Info = log.New(infoLoggerOutput, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		Warn = log.New(warningLoggerOutput, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
		Error = log.New(errorLoggerOutput, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		break
	case "info", "information":
		if !logOptions.LogFileIsActive {
			infoLoggerOutput = os.Stdout
			warningLoggerOutput = os.Stdout
			errorLoggerOutput = os.Stderr
		}

		Debug = log.New(os.Stdin, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
		Info = log.New(infoLoggerOutput, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		Warn = log.New(warningLoggerOutput, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
		Error = log.New(errorLoggerOutput, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		break
	case "warn", "warning":
		if !logOptions.LogFileIsActive {
			warningLoggerOutput = os.Stdout
			errorLoggerOutput = os.Stderr
		}

		Debug = log.New(os.Stdin, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
		Info = log.New(os.Stdin, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		Warn = log.New(warningLoggerOutput, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
		Error = log.New(errorLoggerOutput, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		break
	case "err", "error":
		if !logOptions.LogFileIsActive {
			errorLoggerOutput = os.Stderr
		}

		Debug = log.New(os.Stdin, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
		Info = log.New(os.Stdin, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		Warn = log.New(os.Stdin, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
		Error = log.New(errorLoggerOutput, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		break
	default:
		if !logOptions.LogFileIsActive {
			debugLoggerOutput = os.Stdin
			infoLoggerOutput = os.Stdout
			warningLoggerOutput = os.Stdout
			errorLoggerOutput = os.Stderr
		}

		Debug = log.New(debugLoggerOutput, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
		Info = log.New(infoLoggerOutput, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		Warn = log.New(warningLoggerOutput, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
		Error = log.New(errorLoggerOutput, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		break
	}
}

// set logger options from the file configuration
func setLoggerOptions() {
	// if the LogFileIsActive is true, write to the specific file
	// the filename is defined on config.json
	if logOptions.LogFileIsActive {
		output, err := os.OpenFile(logOptions.LogFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}

		debugLoggerOutput = output
		infoLoggerOutput = output
		warningLoggerOutput = output
		errorLoggerOutput = output
	}

	setLevel()
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func init() {
	supportedFileExtensions := []string{
		"txt",
		"log",
	}

	var err error
	logOptions.LogFileIsActive, err = strconv.ParseBool(os.Getenv("LOG_FILE_IS_ACTIVE"))
	if err != nil {
		logOptions.LogFileIsActive = false
	}

	logOptions.LogFileName = os.Getenv("LOG_FILE_NAME")
	if len(logOptions.LogFileName) == 0 {
		logOptions.LogFileName = "logger.log"
	} else {
		elementsInFileName := strings.Split(logOptions.LogFileName, ".")
		if len(elementsInFileName) > 2 || len(elementsInFileName) < 2 {
			log.Println("The file name", logOptions.LogFileName, " is not a correct format, use similar to this: logger.log")
			log.Println("The file name has changed to logger.log")
			logOptions.LogFileName = "logger.log"
		} else {
			if !contains(supportedFileExtensions, elementsInFileName[1]) {
				log.Println("The file extension is not supported, so extension has changed to .log")
				logOptions.LogFileName = elementsInFileName[1] + ".log"
			}
		}
	}

	setLoggerOptions()
}
