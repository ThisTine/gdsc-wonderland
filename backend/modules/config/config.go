package config

import (
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"

	"backend/modules"
	"backend/types/module"
	"backend/utils/text"
)

func Init() {
	conf := new(module.Config)

	// Load configuration to struct
	yml, err := os.ReadFile("config.yaml")
	if err != nil {
		logrus.Fatal("UNABLE TO LOAD CONFIGURATION FILE")
	}
	if err := yaml.Unmarshal(yml, conf); err != nil {
		logrus.Fatal("UNABLE TO PARSE YAML CONFIGURATION FILE")
	}

	// Validate configurations
	if err := text.Validator.Struct(conf); err != nil {
		logrus.Fatal("INVALID CONFIGURATION: " + err.Error())
	}

	// Apply log level configuration
	logrus.SetLevel(logrus.Level(conf.LogLevel))
	spew.Config = spew.ConfigState{Indent: " "}

	mod.Conf = conf
}
