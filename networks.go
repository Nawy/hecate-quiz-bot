package main

func GetWebhookURL(url URL) string {
	host := url.Host
	port := url.Port
	return host + ":" + port + "/" + conf.Bot.Networks.Token
}
