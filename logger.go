package logger

import (
    "runtime"
    "fmt"
    "path/filepath"
    "time"
    "github.com/davecgh/go-spew/spew"
    "errors"
)

const (
    FG_BOLD_BLACK    string = "\x1b[30;1m"
    FG_BOLD_RED      string = "\x1b[31;1m"
    FG_BOLD_GREEN    string = "\x1b[32;1m"
    FG_BOLD_YELLOW   string = "\x1b[33;1m"
    FG_BOLD_BLUE     string = "\x1b[34;1m"
    FG_BOLD_MAGENTA  string = "\x1b[35;1m"
    FG_BOLD_CYAN     string = "\x1b[36;1m"
    FG_BOLD_WHITE    string = "\x1b[37;1m"

    FG_BLACK         string = "\x1b[30m"
    FG_RED           string = "\x1b[31m"
    FG_GREEN         string = "\x1b[32m"
    FG_YELLOW        string = "\x1b[33m"
    FG_BLUE          string = "\x1b[34m"
    FG_MAGENTA       string = "\x1b[35m"
    FG_CYAN          string = "\x1b[36m"
    FG_WHITE         string = "\x1b[37m"

    FG_NORMAL        string = "\x1b[0m"
)

var logLevels = map[string]int{
    "DEBUG": 6,
    "INFO": 5,
    "NOTICE": 4,
    "WARN": 3,
    "ERROR": 2,
    "CRIT": 1,
}

var LogLevel string = "DEBUG"

type setLogger struct {
    level   string
    out     string
}

func Debug(v ...interface{}) {
    if logLevels[LogLevel] == 6 {
        writeLog("DEBUG", v)
    }
}

func Info(v ...interface{}) {
    if logLevels[LogLevel] >=5 {
        writeLog("INFO", v)
    }
}

func Notice(v ...interface{}) {
    if logLevels[LogLevel] >= 4 {
        writeLog("NOTICE", v)
    }
}

func Warn(v ...interface{}) {
    if logLevels[LogLevel] >= 3 {
        writeLog("WARNING", v)
    }
}

func Error(v ...interface{}) {
    if logLevels[LogLevel] >= 2 {
        writeLog("ERROR", v)
    }
}

func Crit(v ...interface{}) {
    if logLevels[LogLevel] >= 2 {
        writeLog("CRITICAL", v)
    }
}

func SetLevel(level string) error {
    if level == "DEBUG" || level == "INFO" || level == "NOTICE" || level == "WARM" || level == "ERROR" || level == "CRIT" {
        LogLevel = level
        return nil
    }
    return errors.New("Invalid log level. Set level: " + LogLevel)
}

func writeLog(level string, v []interface{}) {
    var color, extra string
    switch level {
        case "DEBUG": color = FG_CYAN; extra = "   "
        case "INFO": color = FG_WHITE; extra = "    "
        case "NOTICE": color = FG_GREEN; extra = "  "
        case "WARNING": color = FG_YELLOW; extra = " "
        case "ERROR": color = FG_RED; extra = "   "
        case "CRITICAL": color = FG_BOLD_RED; extra = ""
    }

    _, file, line, _ := runtime.Caller(2)
    fmt.Print(color, "[APP] ", time.Now().Format("2006/01/02 - 15:04:05"), " [", level, "] ", extra, filepath.Base(file), ":", line, "  ▶  ")
    for i, value := range v {
        if level == "DEBUG" {
            value = spew.Sdump(value)
        }
        fmt.Printf("%+v", value)
        if i < len(v) - 1 {
            fmt.Print(" | ")
        }
    }
    fmt.Print(FG_NORMAL, "\n")
}