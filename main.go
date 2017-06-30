package main

import (
	"log"

	"gopkg.in/telegram-bot-api.v4"
)

func main() {
	InitConfig()
	defer InitLogger().Close()
	defer InitStorage().Close()
	InitHandle()
	LoadMessages()
	LoadGames()

	bot, err := tgbotapi.NewBotAPI(conf.Bot.Token)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = conf.Bot.Debug

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil && update.CallbackQuery == nil {
			continue
		}

		bot.Send(Handle(&update))
	}
}
