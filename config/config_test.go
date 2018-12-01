package config

import (
    "testing"
    "fmt"
    "io/ioutil"
)

func createTestConfig(configPath string) error {
    testFileString := `
---

wordCathes:
  normal:
    regex: ".*(норма|norma).*"
    stickers:
      - "CAADAgADnQUAAlOx9wMjvcls38LyPwI"
  gopstop:
    regex: ".*(gop|гоп).*"
    stickers:
      - "CAADAgADXwMAAgw7AAEKTh8jAAH9Q-gAAQI"

`
    fileBytes := []byte(testFileString)
    err := ioutil.WriteFile(configPath, fileBytes, 0644)
    if err != nil {
        err = fmt.Errorf("Cannot create config <%s>:\t%v", configPath, err)
    }
    return err
}

// TestGetConfig tests config loading
func TestGetConfig(t *testing.T) {
    testConfigPath := "../.test.yml"
    err := createTestConfig(testConfigPath)
    if err != nil {
        t.Fatal(err)
    }
    _, err = GetConfig(testConfigPath)
    if err != nil {
        t.Fatal(err)
    }
}
