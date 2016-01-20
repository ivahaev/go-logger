package logger_test

import "github.com/ivahaev/go-logger"

func Example() {
	logger.SetLevel("DEBUG")

	// This is for matching with output time :)
	logger.SetTimeFormat("1999/10/17 - 10:11:10")

	logger.Info("Some string for info", 123, map[string]interface{}{"prop1": "val1", "prop2": 321})

	// Output:
	// [APP] 1999/10/17 - 10:11:10 [INFO]     logger_test.go:11  ▶  Some string for info | 123 | map[prop1:val1 prop2:321]
}
