package tapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GetWebhookInfoResp struct {
	OK     bool `json:"ok"`
	Result struct {
		URL                string `json:"url"`
		HasCustomInfo      bool   `json:"has_custom_info"`
		PendingUpdateCount uint8  `json:"pending_update_count"`
		LastErrorDate      int64  `json:"last_error_date"`
		LastErrorMessage   string `json:"last_error_message"`
		MaxConnections     int    `json:"max_connections"`
		IpAddress          string `json:"ip_address"`
	}
}

// GetWebhookInfo https://core.telegram.org/bots/api#getwebhookinfo
func (obj *Engine) GetWebhookInfo() (*GetWebhookInfoResp, error) {
	resp, err := http.Get(obj.telegramApiURL + obj.telegramBotToken + obj.telegramEnvironment + "/getWebhookInfo")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result GetWebhookInfoResp

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
