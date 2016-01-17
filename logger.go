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
	fgBoldBlack   string = "\x1b[30;1m"
	fbBoldRed     string = "\x1b[31;1m"
	fgBoldGreen   string = "\x1b[32;1m"
	fgBoldYellow  string = "\x1b[33;1m"
	fgBoldBlue    string = "\x1b[34;1m"
	fgBoldMagenta string = "\x1b[35;1m"
	fgBoldCyan    string = "\x1b[36;1m"
	fgBoldWhite   string = "\x1b[37;1m"
	fbBlack       string = "\x1b[30m"
	fgRed         string = "\x1b[31m"
	fgGreen       string = "\x1b[32m"
	fgYellow      string = "\x1b[33m"
	fgBlue        string = "\x1b[34m"
	fgMagenta     string = "\x1b[35m"
	fgCyan        string = "\x1b[36m"
	fgWhite       string = "\x1b[37m"
	fgNormal      string = "\x1b[0m"
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

// Debug logs provided arguments to console with extra info.
// Works only when level sets to DEBUG (default)
func Debug(v ...interface{}) {
	if logLevel == 6 {
		file, line := getFileAndLine()
		writeLog(&logMessage{v, file, line, 6})
	}
}

// Info logs provided arguments to console when level is INFO or DEBUG.
func Info(v ...interface{}) {
	if logLevel >= 5 {
		file, line := getFileAndLine()
		writeLog(&logMessage{v, file, line, 5})
	}
}

// Notice logs provided arguments to console when level is NOTICE, INFO or DEBUG.
func Notice(v ...interface{}) {
	if logLevel >= 4 {
		file, line := getFileAndLine()
		writeLog(&logMessage{v, file, line, 4})
	}
}

// Warn logs provided arguments to console when level is WARN, NOTICE, INFO or DEBUG.
func Warn(v ...interface{}) {
	if logLevel >= 3 {
		file, line := getFileAndLine()
		writeLog(&logMessage{v, file, line, 3})
	}
}

// Error logs provided arguments to console when level is ERROR, WARN, NOTICE, INFO or DEBUG.
func Error(v ...interface{}) {
	if logLevel >= 2 {
		file, line := getFileAndLine()
		writeLog(&logMessage{v, file, line, 2})
	}
}

// Crit logs provided arguments to console when level is CRIT, ERROR, WARN, NOTICE, INFO or DEBUG.
func Crit(v ...interface{}) {
	if logLevel >= 2 {
		file, line := getFileAndLine()
		writeLog(&logMessage{v, file, line, 1})
	}
}

// SetLevel sets level of logging.
// level can be "CRIT", 'ERROR', 'WARN', "NOTICE", "INFO" or "DEBUG"
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
		color = fgCyan
		level = "DEBUG"
		extra = "   "
	case info:
		color = fgNormal
		level = "INFO"
		extra = "    "
	case notice:
		color = fgGreen
		level = "NOTICE"
		extra = "  "
	case warn:
		color = fgYellow
		level = "WARNING"
		extra = " "
	case err:
		color = fgRed
		level = "ERROR"
		extra = "   "
	case crit:
		color = fbBoldRed
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
	out += fmt.Sprint(fgNormal)
	fmt.Println(out)
}

func init() {
	scs = &spew.ConfigState{Indent: "\t", SortKeys: true}
}
