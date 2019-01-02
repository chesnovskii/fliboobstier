package config

import (
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/caarlos0/env"
	"github.com/chesnovsky/fliboobstier/logger"
	"gopkg.in/yaml.v2"
)

// MainConfig is root config
type MainConfig struct {
	TgToken      string               `env:"FLIBOOBSTIER_TG_TOKEN,required"`
	RegexActions map[string]WordCatch `yaml:"regex_actions"`
}

// WordCatch is a struct for each word compare type, with its own media
type WordCatch struct {
	RawRegex string         `yaml:"regex"`
	Regex    *regexp.Regexp `yaml:"omit_me_1,omitempty"`
	ID       string         `yaml:"omit_me_2,omitempty"`
}

// GetConfig is a init func, returning root config
func GetConfig(configPath string) (MainConfig, error) {
	config := MainConfig{}

	err := env.Parse(&config)
	if err != nil {
		localErr := fmt.Errorf("Cannot parse ENV:\n%v", err)
		return config, localErr
	}

	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		localErr := fmt.Errorf("Cannot read <%s>:\n%v ", configPath, err)
		return config, localErr
	}
	logger.Logger.Debugf("Got YAML config:\n%s\n", yamlFile)

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		localErr := fmt.Errorf("Cannot parse YAML:\n%s\n%v", yamlFile, err)
		return config, localErr
	}

	// Compile all the regexes here
	for wordKey, wordData := range config.RegexActions {
		logger.Logger.Debugf("Got word <%s> with data:\t%v", wordKey, wordData)
		logger.Logger.Debug("Compiling regexp")
		compiled, err := regexp.Compile(wordData.RawRegex)
		if err != nil {
			localErr := fmt.Errorf("Cannot compile regex <%s>:\t%v", wordData.RawRegex, err)
			return config, localErr
		}
		updatedWord := wordData
		updatedWord.Regex = compiled
		updatedWord.ID = wordKey
		config.RegexActions[wordKey] = updatedWord
		logger.Logger.Debugf("Compiled succesfull:\t%v", config.RegexActions[wordKey].Regex)
	}
	logger.Logger.Debugf("Got parsed config config:\n%s\n", yamlFile)
	return config, nil
}
