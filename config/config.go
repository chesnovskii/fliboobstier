package config

import (
    "fmt"
    "io/ioutil"

    "github.com/caarlos0/env"
    "github.com/chesnovsky/fliboobstier/logger"
    "gopkg.in/yaml.v2"
)

type MainConfig struct {
    TgToken string               `env:"FLIBOOBSTIER_TG_TOKEN,required"`
    Catches map[string]WordCatch `yaml:"wordCathes"`
}

type WordCatch struct {
    Regex    string   `yaml:"regex"`
    Images   []string `yaml:"images"`
    Stickers []string `yaml:"stickers"`
    Gifs     []string `yaml:"gifs"`
}

func GetConfig() (error, MainConfig) {
    config := MainConfig{}

    err := env.Parse(&config)
    if err != nil {
        err = fmt.Errorf("Cannot parse ENV:\n%v\n", err)
    }

    yamlFile, err := ioutil.ReadFile("config.yml")
    if err != nil {
        err = fmt.Errorf("Cannot read config.yml:\n%v ", err)
        return err, config
    }
    logger.Logger.Debugf("Got YAML config:\n%s\n", yamlFile)

    err = yaml.Unmarshal(yamlFile, &config)
    if err != nil {
        err = fmt.Errorf("Cannot parse YAML:\n%s\n%v\n", yamlFile, err)
        return err, config
    }

    // if config.Regex == nil {
    //     err = fmt.Errorf("Some of cathes is missing <regex> field. Config is invalid.")
    //     return err, config
    // }

    logger.Logger.Debugf("Got parsed config config:\n%s\n", yamlFile)
    return nil, config
}
