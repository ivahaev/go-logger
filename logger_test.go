package logger_test

import "github.com/ivahaev/go-logger"

func Example() {
	logger.SetLevel("DEBUG")

	// This is for matching with output time :)
	logger.SetTimeFormat("1999/10/17 - 10:11:10")

	logger.Debug("Some string for debug", 123, map[string]interface{}{"prop1": "val1", "prop2": 321})
	logger.Info("Some string for info", 123, map[string]interface{}{"prop1": "val1", "prop2": 321})
	logger.Notice("Some string for debug", 123, map[string]interface{}{"prop1": "val1", "prop2": 321})
	logger.Warn("Some string for warning", 123, map[string]interface{}{"prop1": "val1", "prop2": 321})
	logger.Error("Some string for error", 123, map[string]interface{}{"prop1": "val1", "prop2": 321})
	logger.Crit("Some string for critical", 123, map[string]interface{}{"prop1": "val1", "prop2": 321})

	// Output:

	// [APP] 1999/10/17 - 10:11:10 [DEBUG]    logger_test.go:11  ▶  (string) (len=21) "Some string for debug"
	// (int) 123
	// (map[string]interface {}) (len=2) {
	// 	(string) (len=5) "prop1": (string) (len=4) "val1",
	// 	(string) (len=5) "prop2": (int) 321
	// }
	//
	// [APP] 1999/10/17 - 10:11:10 [INFO]     logger_test.go:12  ▶  Some string for info | 123 | map[prop1:val1 prop2:321]
	// [APP] 1999/10/17 - 10:11:10 [NOTICE]   logger_test.go:13  ▶  Some string for debug | 123 | map[prop1:val1 prop2:321]
	// [APP] 1999/10/17 - 10:11:10 [WARNING]  logger_test.go:14  ▶  Some string for warning | 123 | map[prop1:val1 prop2:321]
	// [APP] 1999/10/17 - 10:11:10 [ERROR]    logger_test.go:15  ▶  Some string for error | 123 | map[prop1:val1 prop2:321]
	// [APP] 1999/10/17 - 10:11:10 [CRITICAL] logger_test.go:16  ▶  Some string for critical | 123 | map[prop1:val1 prop2:321]
}
