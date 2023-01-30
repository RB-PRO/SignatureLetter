package app

import (
	"github.com/RB-PRO/SignatureLetter/pkg/signature"
)

func Run() {
	signature.Сonvertor(signature.People{
		Name:     "Name",
		Working:  "Working",
		Email:    "Email",
		Company:  "Company",
		Adres:    "Adres",
		Image:    "Image",
		Telegram: "Telegram",
		Whatsapp: "Whatsapp",
	})
}
