package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RB-PRO/SignatureLetter/pkg/imgbb"
	"github.com/RB-PRO/SignatureLetter/pkg/signature"
)

func Run() {
	pep := signature.People{
		Name:    "Name",
		Working: "Working",
		Email:   "Email",
		//Company:  "Company",
		Adres:    "Adres",
		Image:    "Image",
		Telegram: "Telegram",
		Whatsapp: "Whatsapp",
	}
	pepHtml, errConvertor := pep.Сonvertor()
	if errConvertor != nil {
		log.Fatal(errConvertor)
	}
	fmt.Println(pepHtml)

	// Получить API key из файла "imgbb"
	imgbbAPI_key, imgbb_error := dataFile("imgbbKey")
	if imgbb_error != nil {
		log.Fatal(imgbb_error)
	}

	// Создаём экземпляр клиента imgbb
	imgb := imgbb.NewImgbbUser(imgbbAPI_key)

	// Переводим локальную картинку в base64
	picBase64, baseError := PicToBase64("roma.jpeg")
	if baseError != nil {
		log.Fatal(baseError)
	}

	// Загружаем картинку на сервер
	respUpload, uploadError := imgb.Upload(picBase64, "Roma")
	if imgbb_error != nil {
		log.Fatal(uploadError)
	}
	if respUpload.Status != http.StatusOK { // проверка на отрицательный ответ
		log.Println("Ошибка:", respUpload.Error.Message)
	}
}
