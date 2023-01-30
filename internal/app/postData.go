package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RB-PRO/SignatureLetter/pkg/imgbb"
	"github.com/RB-PRO/SignatureLetter/pkg/signature"
	"github.com/polds/imgbase64"
)

func Run() {
	pep := signature.People{
		Name:     "Name",
		Working:  "Working",
		Email:    "Email",
		Company:  "Company",
		Adres:    "Adres",
		Image:    "Image",
		Telegram: "Telegram",
		Whatsapp: "Whatsapp",
	}
	pepHtml := pep.Сonvertor()
	fmt.Println(pepHtml)

	// Получить API key из файла "imgbb"
	imgbbAPI_key, imgbb_error := dataFile("imgbbKey")
	if imgbb_error != nil {
		log.Fatal(imgbb_error)
	}

	// Создаём экземпляр клиента imgbb
	imgb := imgbb.NewImgbbUser(imgbbAPI_key)
	// Переводим локальную картинку в base64
	//picBase64 := PictureToBase64("roma.jpeg")
	picBase64, baseError := imgbase64.FromLocal("roma.jpeg")
	if baseError != nil {
		log.Fatal(baseError)
	}
	//picBase64 = "iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHxgljNBAAO9TXL0Y4OHwAAAABJRU5ErkJggg=="

	// Загружаем картинку на сервер
	respUpload, uploadError := imgb.Upload(picBase64, "Roma")
	if imgbb_error != nil {
		log.Fatal(uploadError)
	}
	if respUpload.Status != http.StatusOK { // проверка на отрицательный ответ
		log.Println("Ошибка:", respUpload.Error.Message)
	}
	fmt.Println(respUpload)
}
