package main

/*
as reference used:
https://medium.com/swlh/build-a-telegram-bot-in-go-in-9-minutes-e06ad38acef1
https://github.com/fpaupier/telegrap
*/

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

const happyCherrySticker = "CAACAgIAAxkBAAMGZNfUqkaAMTh87Ji642z6mHvLnnQAAgUAA8A2TxP5al-agmtNdTAE"
const telegramApiBaseUrl = "https://api.telegram.org/bot"
const telegramApiSendStickerMethod = "sendSticker"

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}

type Chat struct {
	Id int `json:"id"`
}

func Main(update Update) {
	if strings.ToLower(update.Message.Text) == "yo" {
		_ = sendStickerToTelegram(update.Message.Chat.Id, happyCherrySticker)
	}
}

func sendStickerToTelegram(chatId int, stickerId string) error {
	var telegramApiCall = fmt.Sprintf("%s%s/%s", telegramApiBaseUrl, os.Getenv("HELLO_BOT_API_KEY"), telegramApiSendStickerMethod)
	if _, err := http.PostForm(
		telegramApiCall,
		url.Values{
			"chat_id": {strconv.Itoa(chatId)},
			"sticker": {stickerId},
		}); err != nil {
		log.Printf("error when send sticker: %s", err.Error())
		return err
	} else {
		return nil
	}
}
