package logger

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
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

const (
	_ = iota
	crit
	err
	warn
	notice
	info
	debug
)

var (
	logLevels = map[string]int{
		"DEBUG":  6,
		"INFO":   5,
		"NOTICE": 4,
		"WARN":   3,
		"ERROR":  2,
		"CRIT":   1,
	}
	logLevel int = 6
	scs          = spew.NewDefaultConfig()
)

type logMessage struct {
	v     []interface{}
	file  string
	line  int
	level int
}

type setLogger struct {
	level string
	out   string
}

func Debug(v ...interface{}) {
	if logLevel == 6 {
		file, line := getFileAndLine()
		writeLog(&logMessage{v, file, line, 6})
	}
}

func Info(v ...interface{}) {
	if logLevel >= 5 {
		file, line := getFileAndLine()
		writeLog(&logMessage{v, file, line, 5})
	}
}

func Notice(v ...interface{}) {
	if logLevel >= 4 {
		file, line := getFileAndLine()
		writeLog(&logMessage{v, file, line, 4})
	}
}

func Warn(v ...interface{}) {
	if logLevel >= 3 {
		file, line := getFileAndLine()
		writeLog(&logMessage{v, file, line, 3})
	}
}

func Error(v ...interface{}) {
	if logLevel >= 2 {
		file, line := getFileAndLine()
		writeLog(&logMessage{v, file, line, 2})
	}
}

func Crit(v ...interface{}) {
	if logLevel >= 2 {
		file, line := getFileAndLine()
		writeLog(&logMessage{v, file, line, 1})
	}
}

func SetLevel(level string) error {
	level = strings.ToUpper(level)
	if numLevel, ok := logLevels[level]; ok {
		logLevel = numLevel
		return nil
	}
	return errors.New("Invalid log level: " + level)
}

func getFileAndLine() (string, int) {
	_, file, line, _ := runtime.Caller(2)
	return filepath.Base(file), line
}

func writeLog(v *logMessage) {
	var (
		color string
		level string
		extra string
	)
	switch v.level {
	case debug:
		color = FG_CYAN
		level = "DEBUG"
		extra = "   "
	case info:
		color = FG_WHITE
		level = "INFO"
		extra = "    "
	case notice:
		color = FG_GREEN
		level = "NOTICE"
		extra = "  "
	case warn:
		color = FG_YELLOW
		level = "WARNING"
		extra = " "
	case err:
		color = FG_RED
		level = "ERROR"
		extra = "   "
	case crit:
		color = FG_BOLD_RED
		level = "CRITICAL"
		extra = ""
	}

	out := fmt.Sprint(color, "[APP] ", time.Now().Format("2006/01/02 - 15:04:05"), " [", level, "] ", extra, v.file, ":", v.line, "  ▶  ")
	for i, value := range v.v {
		if v.level == debug {
			value = scs.Sdump(value)
		}
		out += fmt.Sprintf("%+v", value)
		if v.level != debug && i < len(v.v)-1 {
			out += fmt.Sprint(" | ")
		}
	}
	out += fmt.Sprint(FG_NORMAL)
	fmt.Println(out)
}

func init() {
	scs = &spew.ConfigState{Indent: "\t", SortKeys: true}
}
