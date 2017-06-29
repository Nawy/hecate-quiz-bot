package main

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"io/ioutil"
)

// UserStatus for right handlers
type UserStatus string

const (
	RENAME = UserStatus("rename")
	QUESTION = UserStatus("question")
	IDLE = UserStatus("idle")
)

var handlers map[string]func(*tgbotapi.Update) *tgbotapi.MessageConfig = make(map[string]func(*tgbotapi.Update) *tgbotapi.MessageConfig)
var helpMsg string

func InitHandle() {
	handlers["rename"] = cmdChangeName
	handlers["help"] = cmdHelp

	data, err := ioutil.ReadFile(conf.Bot.Resources.Help)
	if err != nil {
		log.Fatal(err)
	}
	helpMsg = string(data)
}

func Handle(update *tgbotapi.Update) *tgbotapi.MessageConfig {
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


