package tapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// DeleteMessages - is implement https://core.telegram.org/bots/api#deletemessages
type DeleteMessages struct {
	ChatID     int64    `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageIDs []uint64 `json:"message_ids" // A JSON-serialized list of 1-100 identifiers of messages to delete. See deleteMessage for limitations on which messages can be deleted`
}

// DeleteMessages - deletes one or more messages in a chat or channel.
func (obj *Engine) DeleteMessages(deleteMessages DeleteMessages) (bool, error) {
	body, err := json.Marshal(deleteMessages)
	if err != nil {
		return false, fmt.Errorf("marshal deleteMessages error: %w", err)
	}

	resp, err := http.Post(obj.telegramApiURL+obj.telegramBotToken+obj.telegramEnvironment+"/deleteMessages", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false, fmt.Errorf("send deleteMessages to user error: %w", err)
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("deleteMessages to user error: %d %s", resp.StatusCode, string(body))
	}

	var resultMsg struct {
		Ok     bool `json:"ok"`
		Result bool `json:"result"`
	}

	err = json.Unmarshal(body, &resultMsg)
	if err != nil {
		return false, fmt.Errorf("unmarshal deleteMessages %s error: %w", string(body), err)
	}

	return resultMsg.Result, nil
}
