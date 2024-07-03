package tapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// DeleteWebhook https://core.telegram.org/bots/api#deletewebhook
type DeleteWebhook struct {
	DropPendingUpdates bool `json:"drop_pending_updates"` // Pass True to drop all pending updates
}

// DeleteWebhookResp response for DeleteWebhook request
type DeleteWebhookResp struct {
	OK          bool   `json:"ok"`
	Result      bool   `json:"result"`
	Description string `json:"description"`
}

func (obj *Engine) DeleteWebhook() (*DeleteWebhookResp, error) {
	data := DeleteWebhook{true}
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("marshal sMsg error: %w", err)
	}

	resp, err := http.Post(obj.telegramApiURL+obj.telegramBotToken+obj.telegramEnvironment+"/setWebhook", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("send message to user error: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response status code %d", resp.StatusCode)
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result DeleteWebhookResp

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
