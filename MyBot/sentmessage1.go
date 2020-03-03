package main

import (
	"log"
	"net/http"
	"qqbotapi"
)

func main() {
	bot, err := qqbotapi.NewBotAPI("123456789ab", "http://192.168.0.1:5700", "CQHTTP_SECRET")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	u := qqbotapi.NewWebhook("/webhook_endpoint")
	u.PreloadUserInfo = true

	bot.ListenForWebhookSync(u, func(update qqbotapi.Update) interface{} {

		log.Printf("[%s] %s", update.Message.From.String(), update.Message.Text)

		return map[string]interface{}{
			"reply": update.Message.Text,
		}
	})

	// Send a text-img message
	s := bot.NewMessage(10000000, "group").
		At("1232332333").
		Text("嘤嘤嘤").
		NewLine().
		FaceByName("调皮").
		Text("这是一个测试").
		ImageBase64("img.jpg").
		Send()

	// Withdraw that message
	if s.Err == nil {
		bot.DeleteMessage(s.Result.MessageID)
	}

	// Send a stand-alone message (No need to call Send())
	bot.NewMessage(10000000, "private").
		Dice()

	http.ListenAndServe("0.0.0.0:8443", nil)
}
