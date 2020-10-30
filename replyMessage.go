package tapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (obj *Engine) Reply(baseMsg *Message, replyMsg string, replyMarkup ...interface{}) error {

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
		return fmt.Errorf("marshal sMsg error: %s", err)
	}

	log.Println(string(body))

	// TODO допилить handlers
	resp, err := http.Post(obj.telegramApiURL+obj.telegramBotToken+"/sendMessage", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("send message to user error: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("return resp status code: %d", resp.StatusCode)
	}

	return nil
}
