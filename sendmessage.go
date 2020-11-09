package tapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (obj *Engine) SendMessage(chatID int64, msg string, replyMarkup ...interface{}) error {
	sMsg := replyMsgStruct{
		ChatID: chatID,
		Text:   msg,
	}

	if len(replyMarkup) != 0 {
		sMsg.ReplyMarkup = replyMarkup[0]
	}

	body, err := json.Marshal(sMsg)
	if err != nil {
		return fmt.Errorf("marshal sMsg error: %s", err)
	}

	resp, err := http.Post(obj.telegramApiURL+obj.telegramBotToken+"/sendMessage", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("send message to user error: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("return resp status code: %d", resp.StatusCode)
	}

	return nil
}
