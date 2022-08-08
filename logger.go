package logger

import (
	"errors"
	"io"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type Level uint32

var AllLevels = []Level{
	PanicLevel,
	FatalLevel,
	ErrorLevel,
	WarnLevel,
	InfoLevel,
	DebugLevel,
	TraceLevel,
}

const (
	PanicLevel Level = iota
	FatalLevel       = 1
	ErrorLevel       = 2
	WarnLevel        = 3
	InfoLevel        = 4
	DebugLevel       = 5
	TraceLevel       = 6
)

type LoggerInfo struct {
	ProgramName string
	FileName    string
	Level       string
	LogToFile   bool
	PrettyPrint bool
	JSONOutput  bool
}

type Logger struct {
	logger *logrus.Entry
}

type ILogger interface {
	Debug(...interface{})
	Info(...interface{})
	Warning(...interface{})
	Error(...interface{})
	Panic(...interface{})
	Fatal(...interface{})
}

// Private function to get
func getLevel(level string) Level {
	level = strings.ToLower(level)
	switch strings.ToLower(level) {
	case "panic":
		return PanicLevel
	case "fatal":
		return FatalLevel
	case "error":
		return ErrorLevel
	case "warn", "warning":
		return WarnLevel
	case "info":
		return InfoLevel
	case "debug":
		return DebugLevel
	case "trace":
		return TraceLevel
	default:
		return DebugLevel
	}
}

func logToFile(fileName string) (io.Writer, error) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return file, errors.New("Can't write in " + fileName)
	}

	return file, nil
}

func NewLogger(loggerInfo LoggerInfo) ILogger {

	var baseLogger = logrus.New()

	if len(loggerInfo.ProgramName) == 0 {
		currentPathExe, _ := os.Executable()
		loggerInfo.ProgramName = filepath.Base(currentPathExe)
	}

	logger := baseLogger.WithFields(logrus.Fields{
		"program": loggerInfo.ProgramName,
	})

	logger.Logger.SetLevel(logrus.Level(getLevel(loggerInfo.Level)))

	if loggerInfo.LogToFile {
		writer, err := logToFile(loggerInfo.FileName)
		if err != nil {
			logger.Logger.Errorln(err)
		} else {
			logger.Logger.Out = writer
		}
	}

	if loggerInfo.JSONOutput {
		logger.Logger.Formatter = &logrus.JSONFormatter{
			PrettyPrint: loggerInfo.PrettyPrint,
		}
	}

	return &Logger{
		logger: logger,
	}
}

func (l *Logger) Debug(args ...interface{}) {
	_, fileName, numberLine, _ := runtime.Caller(1)
	l.logger = l.logger.WithField("location", path.Base(fileName)+":"+strconv.Itoa(numberLine))
	l.logger.Debug(args...)
}

func (l *Logger) Info(args ...interface{}) {
	_, fileName, numberLine, _ := runtime.Caller(1)
	l.logger = l.logger.WithField("location", path.Base(fileName)+":"+strconv.Itoa(numberLine))
	l.logger.Info(args...)
}

func (l *Logger) Warning(args ...interface{}) {
	_, fileName, numberLine, _ := runtime.Caller(1)
	l.logger = l.logger.WithField("location", path.Base(fileName)+":"+strconv.Itoa(numberLine))
	l.logger.Warn(args...)
}

func (l *Logger) Error(args ...interface{}) {
	_, fileName, numberLine, _ := runtime.Caller(1)
	l.logger = l.logger.WithField("location", path.Base(fileName)+":"+strconv.Itoa(numberLine))
	l.logger.Error(args...)
}

func (l *Logger) Panic(args ...interface{}) {
	_, fileName, numberLine, _ := runtime.Caller(1)
	l.logger = l.logger.WithField("location", path.Base(fileName)+":"+strconv.Itoa(numberLine))
	l.logger.Panic(args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	_, fileName, numberLine, _ := runtime.Caller(1)
	l.logger = l.logger.WithField("location", path.Base(fileName)+":"+strconv.Itoa(numberLine))
	l.logger.Fatalln(args...)
}
