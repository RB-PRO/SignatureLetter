package app

import (
	"fmt"
	"log"

	"github.com/RB-PRO/SignatureLetter/pkg/signature"
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

		log.Println(update.Message.Chat.UserName, "-", update.Message.Text, ">", update.Message.Caption)
		//log.Printf("%#v\n", update.Message)

		if !update.Message.IsCommand() { // ignore any non-command Messages
			if update.Message.Caption != "" {
				if update.Message.Photo != nil {
					//DirectURL, fileError := bot.GetFileDirectURL(update.Message.Photo[0].FileID)
					DirectURL, fileError := bot.GetFileDirectURL(update.Message.Photo[0].FileID)
					//bot.GetFile()
					if fileError == nil {
						fmt.Println(DirectURL)

						// Проверка на наличие фотографии и текста
						signaturePeople, signError := InputStrToPeople(update.Message.Caption)
						if signError == nil {
							log.Printf("Пользователь %s забил структуру: %#v\n", update.Message.Chat.UserName, signaturePeople)
							htmlEmail, errorSingatureParse := signaturePeople.Сonvertor()
							if errorSingatureParse == nil {
								// Сохранить в html
								errorSave := signature.SaveHTML(htmlEmail, signaturePeople.Name)
								if errorSave != nil {
									bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, errorSave.Error()))
								}

								// ПРочитать сохранённый html
								b, errorReadFile := signature.ReadHTML(signaturePeople.Name)
								if errorReadFile != nil {
									bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, errorReadFile.Error()))
								}

								bot.Send(tgbotapi.NewDocument(update.Message.Chat.ID, tgbotapi.FileBytes{Name: signaturePeople.Name + ".txt", Bytes: b}))
							} else {
								bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, errorSingatureParse.Error()))
								continue
							}
						} else {
							bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, signError.Error()))
							continue
						}
					} else {
						bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, fileError.Error()))
					}
					continue
				} else {
					bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Не вижу фотографию.\nНужно отправить фотографию вместе с текстом."))
					continue
				}
			} else {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Не вижу текста.\nНужно отправить фотографию вместе с текстом."))
				continue
			}
		}

		switch update.Message.Command() {
		case "start":
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, `Привет! Я бот, который поможет тебе собрать подпись. Необходимо передать:
			- Имя Фамилия
			- Должность
			- e-mail
			- Адрес
			- Ник в телеграм
			- Номер на вотсап
			> Фото человека`))
		case "example":
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Вот в таком формате Вам необходимо отправить мне данные:"))

			//bot.Send(tgbotapi.NewDocument(update.Message.Chat.ID, tgbotapi.FileBytes{Name: "roma.txt", Bytes: b}))
			bot.Send(tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.RequestFileData{}))
			/*`Роман Блинов
			Студент
			romanblinov2013@yandex.ru
			Россия, Москва, 2-я Бауманская улица, 5, стр. 1
			rb_pro
			79269755457`*/
			continue
		default:
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Я не знаю такую команду\nПопробуй /start"))
			continue
		}
	}
}
