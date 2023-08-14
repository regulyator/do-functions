package main

/*
as reference used:
https://medium.com/swlh/build-a-telegram-bot-in-go-in-9-minutes-e06ad38acef1
https://github.com/fpaupier/telegrap
*/

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

const happyCherryStickerUniqId = "AgADBQADwDZPEw"
const gaySticker = "CAACAgIAAxkBAANjZNn1ccdS0m3cjMrPNnxQAAHpj9y3AAJCAANkYXEufUbkbOo3ZmgwBA"
const shameSticker = "CAACAgIAAxkBAAN_ZNn4wzhLcOx-PhIFpnxBbXTm9FcAAtcDAAJ06TMGuqdkHDQwKf4wBA"
const firedSticker = "CAACAgIAAxkBAAN9ZNn4stKLlzim_wRBlL5mJNZnIlMAAtkDAAJ06TMGvQu3rl5frtQwBA"
const telegramApiBaseUrl = "https://api.telegram.org/bot"
const telegramApiSendStickerMethod = "sendSticker"

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Id      int     `json:"message_id"`
	Text    string  `json:"text"`
	Chat    Chat    `json:"chat"`
	Sticker Sticker `json:"sticker"`
}

type Chat struct {
	Id int `json:"id"`
}

type Sticker struct {
	FileUniqId string `json:"file_unique_id"`
	FileId     string `json:"file_id"`
}

func Main(update Update) {

	if update.Message.Sticker.FileUniqId == happyCherryStickerUniqId {
		replyMap := map[int]string{
			0: gaySticker,
			1: shameSticker,
			2: firedSticker,
		}

		_ = sendReplyStickerToTelegram(update.Message.Id, update.Message.Chat.Id, replyMap[rand.Intn(3)])
	}

}

func sendReplyStickerToTelegram(messageId int, chatId int, stickerId string) error {
	var telegramApiCall = fmt.Sprintf("%s%s/%s", telegramApiBaseUrl, os.Getenv("HELLO_BOT_API_KEY"), telegramApiSendStickerMethod)
	if _, err := http.PostForm(
		telegramApiCall,
		url.Values{
			"chat_id":             {strconv.Itoa(chatId)},
			"sticker":             {stickerId},
			"reply_to_message_id": {strconv.Itoa(messageId)},
		}); err != nil {
		log.Printf("error when send sticker: %s %d", err.Error(), messageId)
		return err
	} else {
		return nil
	}
}
