package app

import (
	"errors"
	"strings"

	"github.com/RB-PRO/SignatureLetter/pkg/signature"
)

// Перевести входную строку в структуру
func InputStrToPeople(str string) (signature.People, error) {
	strs := strings.Split(str, "\n")
	if len(strs) != 7 {
		return signature.People{}, errors.New("неверный ввод данных")
	}
	var peop signature.People
	peop.Name = strs[0]    // - Имя Фамилия
	peop.Working = strs[1] // - Должность
	peop.Email = strs[2]   // - e-mail
	//peop.Company = strs[3]  // - Компания
	peop.Adres = strs[3]    // - Адрес
	peop.Telegram = strs[4] // - Ник в телеграм
	peop.Whatsapp = strs[5] // - Номер на вотсап (79269755457 )
	peop.Site = strs[6]     // - Ссылка на сайт
	return peop, nil
}
