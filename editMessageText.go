package tapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// EditMessageText method for edit bot message https://core.telegram.org/bots/api#editmessagetext
func (obj *Engine) EditMessageText(chatID int64, msgID uint64, msg string, replyMarkup ...interface{}) error {
	sMsg := replyMsgStruct{
		ChatID:    chatID,
		MessageID: &msgID,
		Text:      msg,
	}

	if len(replyMarkup) != 0 {
		sMsg.ReplyMarkup = replyMarkup[0]
	}

	body, err := json.Marshal(sMsg)
	if err != nil {
		return fmt.Errorf("marshal sMsg error: %s", err)
	}

	resp, err := http.Post(obj.telegramApiURL+obj.telegramBotToken+"/editMessageText", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("send message to user error: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("return resp status code: %d", resp.StatusCode)
	}

	return nil
}
