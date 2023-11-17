package tapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Reply simple method for reply message. Used sendMessage method of Telegram api https://core.telegram.org/bots/api#sendmessage
func (obj *Engine) Reply(baseMsg *Message, replyMsg string, replyMarkup ...interface{}) (*Message, error) {

	sMsg := replyMsgStruct{
		ChatID:           baseMsg.Message.Chat.ID,
		Text:             replyMsg,
		ReplyToMessageID: &baseMsg.Message.MessageID,
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
