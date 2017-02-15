package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

var Settings map[string]string

func loadYml() {
	filename, _ := filepath.Abs("./settings.yml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &Settings)
	if err != nil {
		panic(err)
	}
}
