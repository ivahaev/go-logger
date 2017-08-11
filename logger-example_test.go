package logger_test

import "github.com/ivahaev/go-logger"

func ExampleInfo() {
	// This is for matching with output time :)
	logger.SetTimeFormat("dummy-time")

	logger.SetLevel("INFO")

	logger.Info("Some string for info", 123, []interface{}{"val1", 321})

	// Output:
	// [APP] dummy-time [INFO]     logger-example_test.go:11  ▶  Some string for info | 123 | [val1 321]

	logger.SetLevel("NOTICE")
}

func ExampleInfoLowLevel() {
	logger.SetLevel("NOTICE")

	// We set level to NOTICE so no log will output
	logger.Info("Some string for info", 123, []interface{}{"val1", 321})

	// Output:
}

func ExampleInfoF() {
	logger.SetTimeFormat("dummy-time")

	logger.SetLevel("INFO")

	logger.Infof("Some string for %v", "info")

	// Output:
	// [APP] dummy-time [INFO]     logger-example_test.go:33  ▶  Some string for info

	logger.SetLevel("NOTICE")
}

func ExampleInfoFMultipleArguments() {
	logger.SetTimeFormat("dummy-time")

	logger.SetLevel("INFO")

	logger.Infof("Some string for info | %v | %s | %s", "123", "val1", "321")

	// Output:
	// [APP] dummy-time [INFO]     logger-example_test.go:46  ▶  Some string for info | 123 | val1 | 321

	logger.SetLevel("NOTICE")
}
