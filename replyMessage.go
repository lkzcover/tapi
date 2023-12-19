package tapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Reply simple method for reply message. Used sendMessage method of Telegram api https://core.telegram.org/bots/api#sendmessage
func (obj *Engine) Reply(baseMsg *Message, msg MsgParams, replyMarkup ...interface{}) (*ResultMsg, error) {

	sMsg := replyMsgStruct{
		ChatID:           baseMsg.Message.Chat.ID,
		Text:             msg.Text,
		ReplyToMessageID: &baseMsg.Message.MessageID,
	}

	if len(msg.Format) != 0 {
		sMsg.ParseMode = &msg.Format
	}

	if len(replyMarkup) != 0 {
		sMsg.ReplyMarkup = replyMarkup[0]
	}

	body, err := json.Marshal(sMsg)
	if err != nil {
		return nil, fmt.Errorf("marshal sMsg error: %s", err)
	}

	// TODO допилить handlers
	resp, err := http.Post(obj.telegramApiURL+obj.telegramBotToken+"/sendMessage", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("send message to user error: %s", err)
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		_, err := parseError(body)
		return nil, err
	}

	var resultMsg ResultMsg

	err = json.Unmarshal(body, &resultMsg)
	if err != nil {
		return nil, err
	}

	return &resultMsg, nil
}
