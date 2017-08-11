package logger

import (
	js "encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
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
	json
)

var (
	logLevels = map[string]int{
		"JSON":   7,
		"DEBUG":  6,
		"INFO":   5,
		"NOTICE": 4,
		"WARN":   3,
		"ERROR":  2,
		"CRIT":   1,
	}
	logLevel   = 6
	timeFormat = "2006/01/02 - 15:04:05"
	scs        = spew.NewDefaultConfig()
	mutex      = &sync.Mutex{}
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

// JSON logs provided arguments to console with json.MarshalIndent each arguments.
// Works only when level sets to DEBUG (default)
func JSON(v ...interface{}) {
	if logLevel == debug {
		writeLog(createMessage(v, json))
	}
}

// Debug logs provided arguments to console with extra info.
// Works only when level sets to DEBUG (default)
func Debug(v ...interface{}) {
	if logLevel == debug {
		writeLog(createMessage(v, debug))
	}
}

// Info logs provided arguments to console when level is INFO or DEBUG.
func Info(v ...interface{}) {
	if logLevel >= info {
		writeLog(createMessage(v, 5))
	}
}

// Notice logs provided arguments to console when level is NOTICE, INFO or DEBUG.
func Notice(v ...interface{}) {
	if logLevel >= notice {
		writeLog(createMessage(v, 4))
	}
}

// Warn logs provided arguments to console when level is WARN, NOTICE, INFO or DEBUG.
func Warn(v ...interface{}) {
	if logLevel >= warn {
		writeLog(createMessage(v, 3))
	}
}

// Error logs provided arguments to console when level is ERROR, WARN, NOTICE, INFO or DEBUG.
func Error(v ...interface{}) {
	if logLevel >= err {
		writeLog(createMessage(v, 2))
	}
}

// Crit logs provided arguments to console when level is CRIT, ERROR, WARN, NOTICE, INFO or DEBUG.
func Crit(v ...interface{}) {
	writeLog(createMessage(v, 1))
}

// Debugf logs provided arguments to console with extra info.
// Works only when level sets to DEBUG (default)
func Debugf(format string, v ...interface{}) {
	if logLevel == debug {
		writeLog(createFormattedMessage(debug, format, v))
	}
}

// Infof logs provided arguments to console when level is INFO or DEBUG.
func Infof(format string, v ...interface{}) {
	if logLevel >= info {
		writeLog(createFormattedMessage(5, format, v))
	}
}

// Noticef logs provided arguments to console when level is NOTICE, INFO or DEBUG.
func Noticef(format string, v ...interface{}) {
	if logLevel >= notice {
		writeLog(createFormattedMessage(4, format, v))
	}
}

// Warnf logs provided arguments to console when level is WARN, NOTICE, INFO or DEBUG.
func Warnf(format string, v ...interface{}) {
	if logLevel >= warn {
		writeLog(createFormattedMessage(3, format, v))
	}
}

// Errorf logs provided arguments to console when level is ERROR, WARN, NOTICE, INFO or DEBUG.
func Errorf(format string, v ...interface{}) {
	if logLevel >= err {
		writeLog(createFormattedMessage(2, format, v))
	}
}

// Critf logs provided arguments to console when level is CRIT, ERROR, WARN, NOTICE, INFO or DEBUG.
func Critf(format string, v ...interface{}) {
	writeLog(createFormattedMessage(1, format, v))
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

// SetTimeFormat sets string format for time.Time.Format() method
// Default is "2006/01/02 - 15:04:05"
func SetTimeFormat(format string) {
	timeFormat = format
}

func createMessage(v []interface{}, level int) *logMessage {
	file, line := getFileAndLine()
	return &logMessage{v, file, line, level}
}

func createFormattedMessage(level int, format string, v []interface{}) *logMessage {
	file, line := getFileAndLine()
	message := fmt.Sprintf(format, v...)
	v = make([]interface{}, 1)
	v[0] = message
	return &logMessage{v, file, line, level}
}

func getFileAndLine() (string, int) {
	_, file, line, _ := runtime.Caller(3)
	return filepath.Base(file), line
}

func writeLog(v *logMessage) {
	// Mutex used just for queueing messages
	mutex.Lock()
	fmt.Println(createLogString(v))
	mutex.Unlock()
}

func createLogString(v *logMessage) string {
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
		color = ""
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

	now := time.Now().Format(timeFormat)
	out := fmt.Sprint(color, "[APP] ", now, " [", level, "] ", extra, v.file, ":", v.line, "  ▶  ")
	for i, value := range v.v {
		switch v.level {
		case debug:
			value = scs.Sdump(value)
		case json:
			val, err := js.MarshalIndent(value, "", "  ")
			if err != nil {
				value = err.Error()
			} else {
				value = string(val)
			}
		}

		out += fmt.Sprintf("%+v", value)
		if v.level != debug && i < len(v.v)-1 {
			out += fmt.Sprint(" | ")
		}
	}
	if v.level != info {
		out += fmt.Sprint(fgNormal)
	}

	return out
}

func init() {
	scs = &spew.ConfigState{Indent: "  ", SortKeys: true}
}
