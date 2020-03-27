package util

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Conf struct {
	QueryDirectory string `yaml:"query_directory"`
	Port           string `yaml:"port"`
}

func (c *Conf) GetConf() *Conf {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}

	if err = yaml.Unmarshal(yamlFile, c); err != nil {
		fmt.Println(err.Error())
	}
	return c
}
