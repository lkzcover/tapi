package tapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// SendMessage - send message to chatID https://core.telegram.org/bots/api#sendmessage
func (obj *Engine) SendMessage(chatID int64, msg string, replyMarkup ...interface{}) (*Message, error) {
	sMsg := replyMsgStruct{
		ChatID: chatID,
		Text:   msg,
	}

	if len(replyMarkup) != 0 {
		sMsg.ReplyMarkup = replyMarkup[0]
	}

	body, err := json.Marshal(sMsg)
	if err != nil {
		return nil, fmt.Errorf("marshal sMsg error: %w", err)
	}

	resp, err := http.Post(obj.telegramApiURL+obj.telegramBotToken+"/sendMessage", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("send message to user error: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("return resp status code: %d", resp.StatusCode)
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rplyMsg Message

	err = json.Unmarshal(body, &rplyMsg)
	if err != nil {
		return nil, err
	}

	return &rplyMsg, nil
}
