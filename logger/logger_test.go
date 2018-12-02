package logger

import (
	"testing"
	"io/ioutil"
	"strings"
)

// TestInit tests logger init
func TestInit(t *testing.T) {
	logPath := "../.test.log"
	Init(logPath)

	logText := "This_is_my_error"
	Logger.Error(logText)

	logFileText, err := ioutil.ReadFile(logPath)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(logFileText), logText) {
		t.Fatal("Cannot find log record on log file")
	}
}
