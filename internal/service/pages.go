package pages

import (
	"github.com/avkosme/golang-api-boilerplate/internal/config"
	"github.com/avkosme/golang-api-boilerplate/internal/delivery/http/telegram"
	pages "github.com/avkosme/golang-api-boilerplate/internal/repository"
)

func Manage(command string) (result bool) {

	result = false

	switch command {
	case "/pages":
		pages.FindAll()

		telegram := new(telegram.Teleram)
		if config.ModeDev != true {
			telegram.Send()
		}

		result = true

	}

	return
}
