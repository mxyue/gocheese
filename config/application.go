package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

var Settings map[string]string



func loadYml() {
	var filename string
	var ok bool
	// filename, err := os.Executable()
	filename, err := filepath.Abs(filepath.Dir(os.Args[0]))
	errPanic(err)
	defer func() {
		_, filename, _, ok = runtime.Caller(0)
		if !ok {
			panic("No caller information")
		}
		readFile(filename)
	}()
	readFile(filename)
}

func readFile(filename string) {
	exPath := path.Dir(filename)
	setting_file, _ := filepath.Abs(exPath + "/settings.yml")
	yamlFile, err := ioutil.ReadFile(setting_file)
	err = yaml.Unmarshal(yamlFile, &Settings)
	errPanic(err)
}

func errPanic(err error) {
	if err != nil {
		panic(err)
	}
}