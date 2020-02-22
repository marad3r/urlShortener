package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

func ReadConfig(c interface{}, yamlPath string) {
	readYaml(c, yamlPath)
	readEnv(c)
}

func readYaml(c interface{}, path string) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		processError("cannot read config.yaml", err)
	}
	err = yaml.Unmarshal(file, c)
	if err != nil {
		processError("cannot unmarshal yaml data", err)
	}
}

func readEnv(c interface{}) {
	err := envconfig.Process("", c)
	if err != nil {
		processError("cannot read env file", err)
	}
}

func processError(desc string, err error) {
	fmt.Printf("%s with error: %v", desc, err)
	os.Exit(2)
}
