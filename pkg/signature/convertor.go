package signature

// Структура входных данных
type People struct {
	Name     string // - Имя Фамилия
	Working  string // - Должность
	Email    string // - e-mail
	Company  string // - Компания
	Adres    string // - Адрес
	Image    string // - Ссылка на фото
	Telegram string // - Ник в телеграм
	Whatsapp string // - Номер на вотсап (79269755457 ) // https://wa.me/79269755457?text=%D0%97%D0%B4%D1%80%D0%B0%D0%B2%D1%81%D1%82%D0%B2%D1%83%D0%B9%D1%82%D0%B5%2C%20%5B%D0%98%D0%BC%D1%8F%20%D0%A4%D0%B0%D0%BC%D0%B8%D0%BB%D0%B8%D1%8F%5D!%20%D0%9F%D0%B8%D1%88%D1%83%20%D0%92%D0%B0%D0%BC%20%D0%BF%D0%BE%20%D0%BF%D0%BE%D0%B2%D0%BE%D0%B4%D1%83%20%5B%D0%BF%D0%BE%D0%B2%D0%BE%D0%B4%5D.
}

// #Конвертировать входные данные в html
//
// ###Входные данные:
// - Имя Фамилия
// - Должность
// - e-mail
// - Компания
// - Адресс
// - Ссылка на фото
// - Ссылка на телеграм
// - Номер на вотсап (79269755457 ) //https://wa.me/79269755457?text=%D0%97%D0%B4%D1%80%D0%B0%D0%B2%D1%81%D1%82%D0%B2%D1%83%D0%B9%D1%82%D0%B5%2C%20%5B%D0%98%D0%BC%D1%8F%20%D0%A4%D0%B0%D0%BC%D0%B8%D0%BB%D0%B8%D1%8F%5D!%20%D0%9F%D0%B8%D1%88%D1%83%20%D0%92%D0%B0%D0%BC%20%D0%BF%D0%BE%20%D0%BF%D0%BE%D0%B2%D0%BE%D0%B4%D1%83%20%5B%D0%BF%D0%BE%D0%B2%D0%BE%D0%B4%5D.
func (people People) Сonvertor() string {
	return "<" + people.Name + "\n" + people.Email + ">"
}
