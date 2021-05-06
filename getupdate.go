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
