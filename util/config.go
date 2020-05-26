package util

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Conf struct {
	QueryDirectory    string `yaml:"query_directory"`
	GinServePort      string `yaml:"gin_serve_port"`
	GRPCServerAddress string `yaml:"grpc_server_address"`
}

func (c *Conf) GetConf() *Conf {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err.Error())
	}

	if err = yaml.Unmarshal(yamlFile, c); err != nil {
		panic(err.Error())
	}
	return c
}
