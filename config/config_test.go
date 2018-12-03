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
    // Create and load config
    testConfigPath := "../.test.yml"
    err := createTestConfig(testConfigPath)
    if err != nil {
        t.Fatal(err)
    }
    config, err := GetConfig(testConfigPath)
    if err != nil {
        t.Fatal(err)
    }

    // Test token from makefile
    myToken := "myLittleTestToken"
    if config.TgToken != myToken{
        t.Fatal("TgToken mismatch")
    }

    // Test words count
    wordsCount := 2
    if len(config.Catches) != wordsCount {
        t.Fatal("Words in wordCathes mismatch")
    }

    // Test regex compile on "normal" cathch
    hitStr := "Да это же норма, епт"
    missStr := "Нет, это уже пиздец"
    hitMatch := config.Catches["normal"].Regex.MatchString(hitStr)
    missMatch := config.Catches["normal"].Regex.MatchString(missStr)
    if !(hitMatch && !missMatch) {
        t.Fatal("Cannot match \"normal\" regex to strings")
    }
}
