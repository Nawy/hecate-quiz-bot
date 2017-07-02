package main

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"regexp"
)

var nameRegexp *regexp.Regexp

func InitHelper() {
	res, _ := regexp.Compile("[a-zA-ZА-Я-а-яЁё0-9]{1,20}")
	nameRegexp = res
}

func getUpdateParam(update *tgbotapi.Update) (int, int64) {
	if update.Message != nil {
		return update.Message.From.ID, update.Message.Chat.ID
	} else {
		return update.CallbackQuery.From.ID, update.CallbackQuery.Message.Chat.ID
	}
}

func logMessage(update *tgbotapi.Update) {
	if update.Message != nil {
		log.Printf("# [%s] -> %s", update.Message.From.ID, update.Message.Text)
	}
}

func sendMessage(update *tgbotapi.Update, text string, chatID int64) *tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatID, EmojiReplace(text))
	return &msg
}

func sendMessageWithMarkup(update *tgbotapi.Update, text string, chatID int64, replyMarkup interface{}) *tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatID, EmojiReplace(text))
	msg.ReplyMarkup = replyMarkup
	return &msg
}

func isRightAnswer(userAnswer string, answers []string) bool {

	for _, answer := range answers {
		if answer == userAnswer {
			return true
		}
	}
	return false
}

