package app

import (
	"bufio"
	"encoding/base64"
	"io"
	"net/http"
	"os"
)

// Фунция получения конфигурационных файлов
//
// Получение значение из файла filename
func dataFile(filename string) (string, error) {
	// Открыть файл
	fileToken, errorToken := os.Open(filename)
	if errorToken != nil {
		return "", errorToken
	}

	// Прочитать значение файла
	data := make([]byte, 512)
	n, err := fileToken.Read(data)
	if err == io.EOF { // если конец файла
		return "", errorToken
	}
	fileToken.Close() // Закрытие файла

	return string(data[:n]), nil
}

// Convert Picture in base64
func PicToBase64(filename string) (string, error) {
	imgFile, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer imgFile.Close()

	fInfo, fInfoError := imgFile.Stat()
	if fInfoError != nil {
		return "", fInfoError
	}
	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	fReader := bufio.NewReader(imgFile)
	_, ReadError := fReader.Read(buf)
	if ReadError != nil {
		return "", ReadError
	}

	return base64.StdEncoding.EncodeToString(buf), nil
}

// Скачать файл по ссылке
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
