package app

import (
	"net/http"
	"testing"

	"github.com/RB-PRO/SignatureLetter/pkg/imgbb"
)

func TestRun(t *testing.T) { // go test ./internal/app/LoadTestImage_test.go
	// Получить API key из файла "imgbb"
	imgbbAPI_key := "f4255d312b02bbe6e72b78d706b77655"

	// Создаём экземпляр клиента imgbb
	imgb := imgbb.NewImgbbUser(imgbbAPI_key)

	// Переводим локальную картинку в base64
	picBase64, baseError := imgbb.PicToBase64("test.jpg")
	if baseError != nil {
		//log.Println(baseError)
		t.Error(baseError)
	}

	// Загружаем картинку на сервер
	respUpload, uploadError := imgb.Upload(picBase64, "test.jpg")
	if uploadError != nil {
		//log.Println(uploadError)
		t.Error(uploadError)
	}
	if respUpload.Status != http.StatusOK { // проверка на отрицательный ответ
		//log.Println("Ошибка:", respUpload.Error.Message)
		t.Error("Ошибка:", respUpload.Error.Message)
	}
	t.Log("Ссылка:", respUpload.Data.Image.URL)
}
