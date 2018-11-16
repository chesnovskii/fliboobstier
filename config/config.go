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
    TgToken string               `env:"FLIBOOBSTIER_TG_TOKEN,required"`
    Catches map[string]WordCatch `yaml:"wordCathes"`
}

// WordCatch is a struct for each word compare type, with its own media
type WordCatch struct {
    RawRegex string        `yaml:"regex"`
    Regex    regexp.Regexp `yaml:"fuck,omitempty"`
    Images   []string      `yaml:"images"`
    Stickers []string      `yaml:"stickers"`
    Gifs     []string      `yaml:"gifs"`
}

// GetConfig is a init func, returning root config
func GetConfig() (MainConfig, error) {
    config := MainConfig{}

    err := env.Parse(&config)
    if err != nil {
        localErr := fmt.Errorf("Cannot parse ENV:\n%v", err)
        return config, localErr
    }

    yamlFile, err := ioutil.ReadFile("config.yml")
    if err != nil {
        localErr := fmt.Errorf("Cannot read config.yml:\n%v ", err)
        return config, localErr
    }
    logger.Logger.Debugf("Got YAML config:\n%s\n", yamlFile)

    err = yaml.Unmarshal(yamlFile, &config)
    if err != nil {
        localErr := fmt.Errorf("Cannot parse YAML:\n%s\n%v", yamlFile, err)
        return config, localErr
    }

    // Compile all the regexes here

    logger.Logger.Debugf("Got parsed config config:\n%s\n", yamlFile)
    return config, nil
}
