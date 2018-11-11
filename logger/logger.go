package logger

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/Sirupsen/logrus"
)

var Logger = logrus.New()

type ContextHook struct{}

func (hook ContextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire ...
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
	Logger.Formatter.(*logrus.TextFormatter).FullTimestamp = true
	Logger.AddHook(ContextHook{})
	if os.Getenv("FLIBOOBSTIER_DEBUG") == "true" {
		Logger.SetLevel(logrus.DebugLevel)
	}
}

func Init(logFile string) {

	if logFile == "" {
		Logger.Out = os.Stdout
		fmt.Println("Init stdout logger")
	} else {
		var file, err = os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		Logger.Out = file
		fmt.Println("Init file logger")
	}
}
