package tapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// SendMessage - send message to chatID https://core.telegram.org/bots/api#sendmessage
func (obj *Engine) SendMessage(chatID int64, msg MsgBody, replyMarkup ...interface{}) (*ResultMsg, error) {
	sMsg := replyMsgStruct{
		ChatID: chatID,
		Text:   msg.Text,
	}

	if len(msg.Format) != 0 {
		sMsg.ParseMode = &msg.Format
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

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, parseError(body)
	}

	var resultMsg ResultMsg

	err = json.Unmarshal(body, &resultMsg)
	if err != nil {
		return nil, err
	}

	return &resultMsg, nil
}
