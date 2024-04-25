package tapi

import "strings"

const defaultTelegramAPIUrl = "https://api.telegram.org/"

type Engine struct {
	telegramApiURL      string
	telegramEnvironment string
	telegramBotToken    string
}

func Init(token string, params ...string) Engine {

	tapi := new(Engine)

	tapi.telegramApiURL = defaultTelegramAPIUrl
	tapi.telegramBotToken = "bot" + token

	if len(params) > 0 {
		if len(params[0]) != 0 {
			tapi.telegramApiURL = params[0]
		}

		if strings.EqualFold(params[1], "test") {
			tapi.telegramEnvironment = "/test"
		}

		return *tapi
	}

	return *tapi

}
