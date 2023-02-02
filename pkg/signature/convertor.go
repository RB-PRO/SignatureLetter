package signature

import (
	"os"
	"strings"
)

// Структура входных данных
type People struct {
	Name    string // - Имя Фамилия
	Working string // - Должность
	Email   string // - e-mail
	//Company  string // - Компания
	Adres    string // - Адрес
	Telegram string // - Ник в телеграм
	Whatsapp string // - Номер на вотсап (79269755457 )
	Site     string // - Ссылка на сайт
	Image    string // - Ссылка на фото
}

// # Конвертировать входные данные в html
//
// ### Входные данные:
// - Имя Фамилия
// - Должность
// - e-mail
// - Компания
// - Адресс
// - Ссылка на телеграм
// - Номер на вотсап (79269755457)
// - Ссылка на фото
func (people People) Сonvertor() ([]byte, error) {
	b, errorReadFile := ReadHTML("html") // just pass the file name
	if errorReadFile != nil {
		return []byte{}, errorReadFile
	}

	strHTML := string(b) // convert content to a 'string'

	// Замена содержимого
	strHTML = strings.ReplaceAll(strHTML, "[people.Name]", people.Name)
	strHTML = strings.ReplaceAll(strHTML, "[people.Working]", people.Working)
	strHTML = strings.ReplaceAll(strHTML, "[people.Email]", people.Email)
	//strHTML = strings.ReplaceAll(strHTML, "[people.Company]", people.Company)
	strHTML = strings.ReplaceAll(strHTML, "[people.Adres]", people.Adres)
	strHTML = strings.ReplaceAll(strHTML, "[people.Image]", people.Image)
	strHTML = strings.ReplaceAll(strHTML, "[people.Telegram]", people.Telegram)
	strHTML = strings.ReplaceAll(strHTML, "[people.Whatsapp]", people.Whatsapp)
	people.Site = strings.ReplaceAll(people.Site, "https://", "")
	people.Site = strings.ReplaceAll(people.Site, "http://", "")
	strHTML = strings.ReplaceAll(strHTML, "[people.Site]", people.Site)

	return []byte(strHTML), nil
}
func ReadHTML(falename string) (b []byte, errorByteFile error) {
	b, errorReadFile := os.ReadFile(falename + ".txt") // just pass the file name
	if errorReadFile != nil {
		return b, errorReadFile
	}
	return b, errorReadFile
}

// Сохранить файл в html
func SaveHTML(data []byte, filename string) error {

	f, err := os.Create(filename + ".txt")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return err
	}

	return nil
}
