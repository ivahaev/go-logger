package logger_test

import "github.com/ivahaev/go-logger"

func ExampleInfo() {
	// This is for matching with output time :)
	logger.SetTimeFormat("1999/10/17 - 10:11:10")

	logger.SetLevel("INFO")

	logger.Info("Some string for info", 123, []interface{}{"val1", 321})

	// Output:
	// [APP] 1999/10/17 - 10:11:10 [INFO]     logger-example_test.go:11  ▶  Some string for info | 123 | [val1 321]

	logger.SetLevel("NOTICE")
}

func ExampleInfoLowLevel() {
	logger.SetLevel("NOTICE")

	// This is for matching with output time :)
	logger.SetTimeFormat("1999/10/17 - 10:11:10")

	// We set level to NOTICE so no log will output
	logger.Info("Some string for info", 123, []interface{}{"val1", 321})

	// Output:
}
