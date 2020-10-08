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
		Text      string `json:"text"`
		Chat      struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

type replyMsgStruct struct {
	ChatID uint64 `json:"chat_id"`
	Text   string `json:"text"`
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

func (obj *Message) Reply(replyMsg, proxyURL, tAPI, botKEY, secretKey string) error {
	//
	//sMsg := replyMsgStruct{
	//	ChatID: obj.Message.Chat.ID,
	//	Text:   replyMsg,
	//}

	//resp, err := gc.PostAesRequest(proxyURL, tAPI+botKEY+"/sendMessage", secretKey, "application/json", sMsg)
	//if err != nil {
	//	return fmt.Errorf("send message to user error: %s", err)
	//}

	//log.Println(string(resp))

	return nil
}
