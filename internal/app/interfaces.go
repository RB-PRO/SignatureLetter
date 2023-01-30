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
		return signature.People{}, errors.New("неверный ввод")
	}
	var peop signature.People
	peop.Name = strs[0]     // - Имя Фамилия
	peop.Working = strs[1]  // - Должность
	peop.Email = strs[2]    // - e-mail
	peop.Company = strs[3]  // - Компания
	peop.Adres = strs[4]    // - Адрес
	peop.Telegram = strs[5] // - Ник в телеграм
	peop.Whatsapp = strs[6] // - Номер на вотсап (79269755457 )
	return peop, nil
}
