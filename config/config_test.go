package config

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTestConfig(configPath string) error {
	testFileString := `
---

regex_actions:
  normal:
    regex: ".*(норма|norma).*"
  gopstop:
    regex: ".*(gop|гоп).*"

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
	assert.Nil(t, err)

	config, err := GetConfig(testConfigPath)
	assert.Nil(t, err)

	// Test token from makefile
	myToken := "myLittleTestToken"
	assert.Equal(t, myToken, config.TgToken)

	// Test words count
	assert.Equal(t, 2, len(config.RegexActions))

	// Test regex compile on "normal" cathch
	assert.Contains(t, config.RegexActions, "normal")
	assert.Contains(t, config.RegexActions, "gopstop")
	hitStr := "Да это же норма, епт"
	missStr := "Нет, это уже пиздец"
	assert.True(t, config.RegexActions["normal"].Regex.MatchString(hitStr))
	assert.False(t, config.RegexActions["normal"].Regex.MatchString(missStr))
}
