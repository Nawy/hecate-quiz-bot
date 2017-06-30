package main

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"io/ioutil"
	"strconv"
)

// UserStatus for right handlers
type UserStatus string

const (
	RENAME = UserStatus("rename")
	QUESTION = UserStatus("question")
	CHOOSE_GAME = UserStatus("choose_game")
	IDLE = UserStatus("idle")
)

var handlers map[string]func(*tgbotapi.Update) *tgbotapi.MessageConfig = make(map[string]func(*tgbotapi.Update) *tgbotapi.MessageConfig)
var helpMsg string

func InitHandle() {
	handlers["rename"] = cmdChangeName
	handlers["games"] = cmdGames
	handlers["help"] = cmdHelp

	data, err := ioutil.ReadFile(conf.Bot.Resources.Help)
	if err != nil {
		log.Fatal(err)
	}
	helpMsg = string(data)
}

func Handle(update *tgbotapi.Update) *tgbotapi.MessageConfig {

	if update.Message != nil {
		log.Printf("# [%s] -> %s", update.Message.From.UserName, update.Message.Text)
	}

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

func cmdGames(update *tgbotapi.Update) *tgbotapi.MessageConfig {

	var buttons [][]tgbotapi.InlineKeyboardButton = make([][]tgbotapi.InlineKeyboardButton, len(GAMES))
	for i, game := range GAMES {
		row := tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(strconv.Itoa(game.Id) + ") " + game.Name, strconv.Itoa(game.Id)),
		)
		buttons[i] = row
	}


	var gamesButtons = tgbotapi.NewInlineKeyboardMarkup(buttons...)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выбери игру:")
	msg.ReplyMarkup = gamesButtons
	//msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	return &msg
}


