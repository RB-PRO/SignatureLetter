package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RB-PRO/SignatureLetter/pkg/imgbb"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartBot() {
	// Получить API key из файла "imgbb"
	imgbbAPI_key, imgbb_error := dataFile("imgbbKey")
	if imgbb_error != nil {
		log.Fatal(imgbb_error)
	}

	// Создаём экземпляр клиента imgbb
	imgb := imgbb.NewImgbbUser(imgbbAPI_key)

	//log.Printf("%#v\n", imgb)

	// ***

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

		// Игнорируем НЕкоманды
		if !update.Message.IsCommand() {

			// Проверка наличия текста в сообщении
			if update.Message.Caption == "" {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Не вижу текста.\nНужно отправить фотографию вместе с текстом."))
				continue
			}

			// Проверка наличия фото в сообщении
			if update.Message.Photo == nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Не вижу фото.\nНужно отправить фотографию вместе с текстом."))
				continue
			}

			// Получить ссылку на сообщение
			DirectURL, fileError := bot.GetFileDirectURL(update.Message.Photo[1].FileID)
			if fileError != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, fileError.Error()))
				continue
			}
			//fmt.Println(DirectURL)

			// Получить ссылку на сообщение
			//fmt.Printf("%#v", update.Message)

			// Проверка на наличие фотографии и текста
			signaturePeople, signError := InputStrToPeople(update.Message.Caption)
			if signError != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, signError.Error()))
				continue
			}

			// Получаем ссылку на сервере на фотографию
			DownFileError := DownloadFile("main.png", DirectURL)
			if signError != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, DownFileError.Error()))
				continue
			}

			// Переводим локальную картинку в base64
			picBase64, baseError := PicToBase64("main.png")
			if baseError != nil {
				log.Println("Ошибка:", baseError)
			}

			// Загружаем картинку на сервер
			respUpload, uploadError := imgb.Upload(picBase64, signaturePeople.Name) // с imgb
			//respUpload, uploadError := imgb.Upload(picBase64, "main") // с телеграма
			if imgbb_error != nil {
				log.Println("Ошибка:", uploadError.Error())
			}
			fmt.Printf("%#v", respUpload)
			if respUpload.Status != http.StatusOK { // проверка на отрицательный ответ
				log.Println("Ошибка:", respUpload.Error.Message)
			}

			signaturePeople.Image = respUpload.Data.Image.URL // Ссылка с сервиса картинок
			signaturePeople.Image = DirectURL                 // Ссылка с телеграма

			log.Printf("Пользователь %s забил структуру: %#v\n", update.Message.Chat.UserName, signaturePeople)

			// Преобразовать входные данные в исходный html
			htmlEmail, errorSingatureParse := signaturePeople.Сonvertor()
			if errorSingatureParse != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, errorSingatureParse.Error()))
				continue
			}

			/*
				// Сохранить в html
				errorSave := signature.SaveHTML(htmlEmail, signaturePeople.Name)
				if errorSave != nil {
					bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, errorSave.Error()))
				}

				// Прочитать сохранённый html
				b, errorReadFile := signature.ReadHTML(signaturePeople.Name)
				if errorReadFile != nil {
					bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, errorReadFile.Error()))
				}
			*/

			bot.Send(tgbotapi.NewDocument(update.Message.Chat.ID, tgbotapi.FileBytes{Name: signaturePeople.Name + ".txt", Bytes: htmlEmail}))
			continue
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
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Это прмиер сообщения, которое я пойму. Пришли мне подобное. Я жду картинку с 6 строками.\nВот в таком формате Вам необходимо отправить мне данные:"))

			// Отправить картинку
			photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("roma.jpeg"))
			msgM, err := bot.Send(photo)
			if err != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, err.Error()))
				continue
			}

			// Редактировать сообщение, добавив текст
			msgME := tgbotapi.NewEditMessageCaption(update.Message.Chat.ID, msgM.MessageID, `Роман Блинов
Студент
romanblinov2013@yandex.ru
Россия, Москва, 2-я Бауманская улица, 5, стр. 1
rb_pro
79269755457`)
			_, err = bot.Send(msgME)
			if err != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, err.Error()))
				continue
			}
			continue

		case "using":
			// Отправить картинку
			photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("pp1.png"))
			msgM, err := bot.Send(photo)
			if err != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, err.Error()))
				continue
			}
			// Редактировать сообщение, добавив текст
			msgME := tgbotapi.NewEditMessageCaption(update.Message.Chat.ID, msgM.MessageID, "Пример того, как нужно загружать подпись на Яндекс почту.\nПереходим в настройки подписей(Все настройки > Личные данные)")
			_, err = bot.Send(msgME)
			if err != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, err.Error()))
				continue
			}

			// Отправить картинку
			photo = tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("pp2.png"))
			msgM, err = bot.Send(photo)
			if err != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, err.Error()))
				continue
			}
			// Редактировать сообщение, добавив текст
			msgME = tgbotapi.NewEditMessageCaption(update.Message.Chat.ID, msgM.MessageID, "Правой кнопкой по форме \"Ваши подписи\" > \"Исследовать элемеент\"")
			_, err = bot.Send(msgME)
			if err != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, err.Error()))
				continue
			}

			// Отправить картинку
			photo = tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("pp3.png"))
			msgM, err = bot.Send(photo)
			if err != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, err.Error()))
				continue
			}
			// Редактировать сообщение, добавив текст
			msgME = tgbotapi.NewEditMessageCaption(update.Message.Chat.ID, msgM.MessageID, "Находим тег <br>, правой кнопкой мыши \"Редактировать как HTML\" и вставляем файл, который Вам прислал ранее я :)")
			_, err = bot.Send(msgME)
			if err != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, err.Error()))
				continue
			}
			continue

		default:
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Я не знаю такую команду\nПопробуй /start"))
			continue
		}
	}
}
