package tapi

const defaultTelegramAPIUrl = "https://api.telegram.org/"

// TODO переименовать в v1.0.0
type Engine struct {
	telegramApiURL   string
	telegramBotToken string
}

func Init(token string, url ...string) Engine {

	tapi := new(Engine)

	tapi.telegramBotToken = token

	if len(url) > 0 {
		tapi.telegramApiURL = url[0]

		return *tapi
	}

	tapi.telegramApiURL = defaultTelegramAPIUrl

	return *tapi

}
