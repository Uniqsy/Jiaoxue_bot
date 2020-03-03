package main

import (
	"fmt",
	"qq-bot-api"
)

func main() {
	bot, err := qqbotapi.NewBotAPI("MyCoolqHttpToken", "http://localhost:5700", "CQHTTP_SECRET")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	u := qqbotapi.NewWebhook("/webhook_endpoint")
	u.PreloadUserInfo = true

	// Use WebHook as event method
	updates := bot.ListenForWebhook(u)
	// Or if you love WebSocket Reverse
	// updates := bot.ListenForWebSocket(u)
	go http.ListenAndServe("0.0.0.0:8443", nil)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.String(), update.Message.Text)

		bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, update.Message.Text)
	}
}