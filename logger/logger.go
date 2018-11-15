package logger

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/Sirupsen/logrus"
)

// Logger is a logger object for global access
var Logger = logrus.New()

// ContextHook is a mess i dunno how working
type ContextHook struct{}

// Levels is a mess i dunno how working
func (hook ContextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire a hook, whatever it means
func (hook ContextHook) Fire(entry *logrus.Entry) error {
	if pc, file, line, ok := runtime.Caller(8); ok {
		funcName := runtime.FuncForPC(pc).Name()
		entry.Data["source"] = fmt.Sprintf("%s:%v:%s", path.Base(file), line, path.Base(funcName))
	}

	return nil
}

func init() {
	Logger.Formatter = new(logrus.TextFormatter)
	Logger.Formatter.(*logrus.TextFormatter).DisableTimestamp = false
	Logger.Formatter.(*logrus.TextFormatter).DisableColors = false
	Logger.Formatter.(*logrus.TextFormatter).TimestampFormat = "2006-01-02 15:04:05"
	Logger.Formatter.(*logrus.TextFormatter).FullTimestamp = false
	Logger.AddHook(ContextHook{})
	if os.Getenv("FLIBOOBSTIER_DEBUG") == "true" {
		Logger.SetLevel(logrus.DebugLevel)
	}
}

// Init for a logger
func Init(logFile string) {
	Logger.Out = os.Stdout
	fmt.Println("Init stdout logger")
}
