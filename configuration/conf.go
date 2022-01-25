package configuration

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		Url string `yaml:"url"`
	}
}

const (
	APPLICATION_PROPERTIES_FILE = "./resources/Application.yml"
)

func LoadAllConfiguration(path string) Config {

	var file []byte
	var err error

	if len(path) > 0 {
		fmt.Printf("reading configuration file from: %s", path)
		file, err = ioutil.ReadFile(path)
	} else {
		fmt.Printf("reading configuration file from: %s", APPLICATION_PROPERTIES_FILE)
		file, err = ioutil.ReadFile(APPLICATION_PROPERTIES_FILE)
	}

	if err != nil {
		file, err = ioutil.ReadFile("." + APPLICATION_PROPERTIES_FILE)
		if err != nil {
			panic(err)
		}
	}

	config := Config{}

	err = yaml.Unmarshal(file, &config)

	if err != nil {
		panic(err)
	}
	return config
}
