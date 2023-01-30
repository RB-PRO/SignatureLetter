package app

import (
	"fmt"

	"github.com/RB-PRO/SignatureLetter/pkg/signature"
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
}
