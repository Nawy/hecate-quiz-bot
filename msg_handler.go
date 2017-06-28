package main

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"io/ioutil"
)

var handlers map[string]func(*tgbotapi.Update) *tgbotapi.MessageConfig = make(map[string]func(*tgbotapi.Update) *tgbotapi.MessageConfig)
var helpMsg string

func InitHandle() {
	handlers["rename"] = cmdChangeName
	handlers["help"] = cmdHelp

	data, err := ioutil.ReadFile(conf.Bot.Commands.Help)
	if err != nil {
		log.Fatal(err)
	}
	helpMsg = string(data)
}

func Handle(update *tgbotapi.Update) *tgbotapi.MessageConfig {
	log.Println("Command:", update.Message.Command())
	handler := handlers[update.Message.Command()]
	if handler != nil {
		return handler(update)
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "И что ты имел ввиду?")
	return &msg
}

func cmdChangeName(update *tgbotapi.Update) *tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Имя значит хочешь поменять?")
	return &msg
}

func cmdHelp(update *tgbotapi.Update) *tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, helpMsg)
	return &msg
}


