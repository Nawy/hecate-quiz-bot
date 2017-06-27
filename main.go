package hecate_quiz_bot

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"os"
)

type AppConfigYAML struct {
	Bot struct {
		Name  string `yaml:"name"`
		Debug bool `yaml:"debug"`
		Networks struct {
			Token string `yaml:"token"`
			Host  string `yaml:"host"`
			Port  string `yaml:"port"`
		}`yaml:"networks"`
	} `yaml:"bot"`
}

var conf AppConfigYAML = AppConfigYAML{}

func initConfig() {
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

func getPathFromArgs() string {
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		if args[i] == "-c" {
			return args[i+1]
		}
	}
	panic("cannot read config from -c option")
}

func main() {
	initConfig()

	bot, err := tgbotapi.NewBotAPI(conf.Bot.Networks.Token)

	bot.Debug = conf.Bot.Debug

	if err != nil {
		log.Fatal(err)
	}
}
