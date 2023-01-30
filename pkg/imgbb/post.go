package imgbb

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const URL string = "https://api.imgbb.com/1/upload"

type ImgbbUser struct {
	API_key string // API Ключ сервиса https://api.imgbb.com/
}

func NewImgbbUser(key string) *ImgbbUser {
	return &ImgbbUser{key}
}

// Стуктура ответа от загрузки на сервер
type ImgbbResponse struct {
	Data struct {
		ID         string `json:"id"`
		Title      string `json:"title"`
		URLViewer  string `json:"url_viewer"`
		URL        string `json:"url"`
		DisplayURL string `json:"display_url"`
		Width      int    `json:"width"`
		Height     int    `json:"height"`
		Size       int    `json:"size"`
		Time       int    `json:"time"`
		Expiration int    `json:"expiration"`
		Image      struct {
			Filename  string `json:"filename"`
			Name      string `json:"name"`
			Mime      string `json:"mime"`
			Extension string `json:"extension"`
			URL       string `json:"url"`
		} `json:"image"`
		Thumb struct {
			Filename  string `json:"filename"`
			Name      string `json:"name"`
			Mime      string `json:"mime"`
			Extension string `json:"extension"`
			URL       string `json:"url"`
		} `json:"thumb"`
		DeleteURL string `json:"delete_url"`
	} `json:"data"`
	Success bool `json:"success"`
	Status  int  `json:"status"`

	StatusCode int `json:"status_code"`
	Error      struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"error"`
	StatusTxt string `json:"status_txt"`
}

// Функция загрузки изображения на сервер.
func (img ImgbbUser) Upload(pictureBase64, name string) (ImgbbResponse, error) {
	// Ответ от сервера
	var imgbbRes ImgbbResponse

	// Выполнить запрос
	resp, responseError := http.PostForm(URL, url.Values{"key": {img.API_key}, "image": {pictureBase64}, "name": {name}})
	if responseError != nil {
		return ImgbbResponse{}, responseError
	}
	defer resp.Body.Close()

	//Считываем ответ запроса
	body, bodyRead := io.ReadAll(resp.Body)
	if bodyRead != nil {
		return ImgbbResponse{}, bodyRead
	}

	//fmt.Println(string(body))

	// Распарсить данные
	responseErrorUnmarshal := json.Unmarshal(body, &imgbbRes)
	if responseErrorUnmarshal != nil {
		return ImgbbResponse{}, responseErrorUnmarshal
	}

	return imgbbRes, responseError
}
