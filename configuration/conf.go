package configuration

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type ConfigModel interface{}

type AllConfigurationContract interface {
	LoadAllConfiguration(string, interface{})
}

type AllConfiguration struct {
	AllConfiguration AllConfigurationContract
	Config           Config
}

type Config struct {
	AllConfiguration ConfigModel
}

const (
	APPLICATION_PROPERTIES_FILE = "./resources/Application.yml"
)

func (configuration *AllConfiguration) LoadAllConfiguration(modelOfConfiguration ConfigModel) {

	fmt.Printf("Reading all configuration from properties: %s", APPLICATION_PROPERTIES_FILE)

	file, err := ioutil.ReadFile(APPLICATION_PROPERTIES_FILE)

	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(file, &modelOfConfiguration)

	if err != nil {
		panic(err)
	}
}
