package configs

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type GlobalConfig struct {
	HttpPort int     `yaml:"http_port"`
	Env      string  `yaml:"env"`
	Command  Command `yaml:"command"`
}

type Command struct {
	Sockd string `yaml:"sockd"`
}

var Config *GlobalConfig

var c = flag.String("c", "", "use <filename> as configuration file")

func Setup() {
	flag.Parse()
	var data []byte
	var err error
	if len(*c) == 0 {
		data, err = ioutil.ReadFile("config.yaml")
	} else {
		data, err = ioutil.ReadFile(*c)
	}
	if err != nil {
		log.Fatalf("Fail to read config file %v", err.Error())
	}
	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		log.Fatalf("yaml unmarshal fail %v", err.Error())
	}
}
