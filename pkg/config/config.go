package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Conf Config

type Smtp struct {
	Host       string `yaml:"host"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	Port       int    `yaml:"port"`
	Entryption string `yaml:"entryption"`
	MailTo     string `yaml:"mailto"`
}

type Config struct {
	Smtp Smtp `yaml:"smtp"`
}

func init() {
	config, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		panic(err)
	}
	if err = yaml.Unmarshal(config, &Conf); err != nil {
		panic(err)
	}
}
