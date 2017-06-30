package main

import (
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type URL struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type AppConfigYAML struct {
	Bot struct {
		Name     string `yaml:"name"`
		Debug    bool   `yaml:"debug"`
		Log      string `yaml:"log"`
		Token 	 string `yaml:"token"`
		Aerospike struct {
			Host string `yaml:"host"`
			Port int `yaml:"port"`
			Namespace string `yaml:"namespace"`
		} `yaml:"aerospike"`
		Resources struct {
			Images string `yaml:"images"`
			Help string `yaml:"help"`
			Games string `yaml:"games"`
			Messages string `yaml:"messages"`
			Database string `yaml:"database"`
		} `yaml:"resources"`
	} `yaml:"bot"`
}

var conf AppConfigYAML = AppConfigYAML{}

func InitConfig() {
	configPath := getPathFromArgs()
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic("Not found config.yaml")
	}

	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		panic("Cannot unmashal config")
	}
}

func InitLogger() *os.File {
	f, err := os.OpenFile(conf.Bot.Log, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetOutput(f)
	log.Println("Log started")
	return f
}

func getPathFromArgs() string {
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		if args[i] == "-c" {
			return args[i+1]
		}
	}
	panic("cannot read config from -c option")
}
