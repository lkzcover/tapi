package tapi

import (
	"fmt"
	"strconv"

	"github.com/lkzcover/tapi/handlers"
)

var offset int64

func init() {
	offset = -1
}

type MessageList struct {
	OK     bool      `json:"ok"`
	Result []Message `json:"result"`
}

type Message struct {
	UpdateID int64 `json:"update_id"`
	Message  struct {
		MessageID uint64 `json:"message_id"`
		From      struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"from"`
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

func (obj *Engine) GetUpdates() ([]Message, error) {

	var messageList MessageList

	err := handlers.GetRequest(obj.telegramApiURL+obj.telegramBotToken+"/getUpdates?offset="+strconv.FormatInt(offset, 10), &messageList)
	if err != nil {
		return nil, err
	}

	if !messageList.OK {
		return nil, fmt.Errorf("get update message not OK. Resp: %+v", messageList)
	}

	// Читаем следующее сообщение
	if len(messageList.Result) != 0 {
		offset = messageList.Result[len(messageList.Result)-1].UpdateID + 1
	}

	return messageList.Result, nil
}
