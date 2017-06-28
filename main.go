package main

import (
	"log"

	"gopkg.in/telegram-bot-api.v4"
)

// func main() {
// 	InitConfig()
// 	InitLogger()
//
// 	bot, err := tgbotapi.NewBotAPI(conf.Bot.Networks.Token)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	bot.Debug = conf.Bot.Debug
//
// 	webhookURL := GetWebhookURL(conf.Bot.Networks.External)
// 	webhook := tgbotapi.NewWebhookWithCert(webhookURL, conf.Bot.Networks.Ssl.Cert)
//
// 	_, err = bot.SetWebhook(webhook)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	updates := bot.ListenForWebhook("/" + conf.Bot.Networks.Token)
//
// 	listenURL := GetWebhookURL(conf.Bot.Networks.Internal)
// 	go http.ListenAndServeTLS(listenURL, conf.Bot.Networks.Ssl.Cert, conf.Bot.Networks.Ssl.Key, nil)
//
// 	for update := range updates {
// 		log.Printf("%+v\n", update)
// 	}
// }

func main() {
	InitConfig()
	defer InitLogger().Close()

	bot, err := tgbotapi.NewBotAPI(conf.Bot.Networks.Token)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = conf.Bot.Debug

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		bot.Send(handler(update))
	}
}
