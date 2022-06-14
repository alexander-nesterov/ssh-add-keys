package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

var (
	configFlag = flag.String("config", "private-keys.yml", "Path")
)

type Config struct {
	Keys []struct {
		Path     string `yaml:"path"`
		Password string `yaml:"passphrase"`
	} `yaml:"keys"`
}

func LoadConfig(filename string) (*Config, error) {
	out, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf(err.Error())
	}
	config := &Config{}
	err = yaml.Unmarshal(out, config)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return config, nil
}

func main() {
	config, err := LoadConfig(*configFlag)
	if err != nil {
		log.Fatalf(err.Error())
	}
	for _, key := range config.Keys {
		fmt.Println("File:", key.Path)
	}
}
