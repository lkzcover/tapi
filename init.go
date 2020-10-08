package tapi

const defaultTelegramAPIUrl = "https://api.telegram.org/"

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
