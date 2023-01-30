package app

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartBot() {
	token, ErrorFile := dataFile("token")
	if ErrorFile != nil {
		log.Fatal(ErrorFile)
	}
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Произошла авторизация %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			if !update.Message.IsCommand() { // ignore any non-command Messages

				// Проверка на наличие фотографии и текста
				signaturePeople, signError := InputStrToPeople(update.Message.Text)
				if signError == nil {
					fmt.Printf("%#v", signaturePeople)
				} else {
					bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, signError.Error()))
				}
				continue
			}
		}

		switch update.Message.Command() {
		case "start":
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, `Привет! Я бот, который поможет тебе собрать подпись. Необходимо передать:
			- Имя Фамилия
			- Должность
			- e-mail
			- Компания
			- Адрес
			- Ник в телеграм
			- Номер на вотсап
			> Фото человека`))

		case "example":
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "sss"))
			continue
		default:
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Я не знаю такую команду\nПопробуй /start"))
			continue
		}
	}
}
