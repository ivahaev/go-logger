package logger

import (
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"path/filepath"
	"runtime"
	"time"
)

const (
	FG_BOLD_BLACK   string = "\x1b[30;1m"
	FG_BOLD_RED     string = "\x1b[31;1m"
	FG_BOLD_GREEN   string = "\x1b[32;1m"
	FG_BOLD_YELLOW  string = "\x1b[33;1m"
	FG_BOLD_BLUE    string = "\x1b[34;1m"
	FG_BOLD_MAGENTA string = "\x1b[35;1m"
	FG_BOLD_CYAN    string = "\x1b[36;1m"
	FG_BOLD_WHITE   string = "\x1b[37;1m"
	FG_BLACK        string = "\x1b[30m"
	FG_RED          string = "\x1b[31m"
	FG_GREEN        string = "\x1b[32m"
	FG_YELLOW       string = "\x1b[33m"
	FG_BLUE         string = "\x1b[34m"
	FG_MAGENTA      string = "\x1b[35m"
	FG_CYAN         string = "\x1b[36m"
	FG_WHITE        string = "\x1b[37m"

	FG_NORMAL string = "\x1b[0m"
)

type logMessage struct {
	v    []interface{}
	file string
	line int
}

var logLevels = map[string]int{
	"DEBUG":  6,
	"INFO":   5,
	"NOTICE": 4,
	"WARN":   3,
	"ERROR":  2,
	"CRIT":   1,
}

var LogLevel string = "DEBUG"

var logDebugChan = make(chan logMessage)
var logInfoChan = make(chan logMessage)
var logNoticeChan = make(chan logMessage)
var logWarnChan = make(chan logMessage)
var logErrorChan = make(chan logMessage)
var logCritChan = make(chan logMessage)

type setLogger struct {
	level string
	out   string
}

func Debug(v ...interface{}) {
	if logLevels[LogLevel] == 6 {
		file, line := getFileAndLine()
		logDebugChan <- logMessage{v, file, line}
	}
}

func Info(v ...interface{}) {
	if logLevels[LogLevel] >= 5 {
		file, line := getFileAndLine()
		logInfoChan <- logMessage{v, file, line}
	}
}

func Notice(v ...interface{}) {
	if logLevels[LogLevel] >= 4 {
		file, line := getFileAndLine()
		logNoticeChan <- logMessage{v, file, line}
	}
}

func Warn(v ...interface{}) {
	if logLevels[LogLevel] >= 3 {
		file, line := getFileAndLine()
		logWarnChan <- logMessage{v, file, line}
	}
}

func Error(v ...interface{}) {
	if logLevels[LogLevel] >= 2 {
		file, line := getFileAndLine()
		logErrorChan <- logMessage{v, file, line}
	}
}

func Crit(v ...interface{}) {
	if logLevels[LogLevel] >= 2 {
		file, line := getFileAndLine()
		logCritChan <- logMessage{v, file, line}
	}
}

func SetLevel(level string) error {
	if level == "DEBUG" || level == "INFO" || level == "NOTICE" || level == "WARM" || level == "ERROR" || level == "CRIT" {
		LogLevel = level
		return nil
	}
	return errors.New("Invalid log level. Set level: " + LogLevel)
}

func getFileAndLine() (string, int) {
	_, file, line, _ := runtime.Caller(2)
	return filepath.Base(file), line
}

func writeLog() {
	for {
		var (
			v                   logMessage
			color, level, extra string
			debug               bool
		)
		select {
		case v = <-logDebugChan:
			color = FG_CYAN
			level = "DEBUG"
			extra = "   "
			debug = true
		case v = <-logInfoChan:
			color = FG_WHITE
			level = "INFO"
			extra = "    "
		case v = <-logNoticeChan:
			color = FG_GREEN
			level = "NOTICE"
			extra = "  "
		case v = <-logWarnChan:
			color = FG_YELLOW
			level = "WARNING"
			extra = " "
		case v = <-logErrorChan:
			color = FG_RED
			level = "ERROR"
			extra = "   "
		case v = <-logCritChan:
			color = FG_BOLD_RED
			level = "CRITICAL"
			extra = ""
		}

		fmt.Print(color, "[APP] ", time.Now().Format("2006/01/02 - 15:04:05"), " [", level, "] ", extra, v.file, ":", v.line, "  ▶  ")
		for i, value := range v.v {
			if debug {
				value = spew.Sdump(value)
			}
			fmt.Printf("%+v", value)
			if i < len(v.v)-1 {
				fmt.Print(" | ")
			}
		}
		fmt.Print(FG_NORMAL, "\n")
	}
}

func init() {
	go writeLog()
}