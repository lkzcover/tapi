package tapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// TODO X-Telegram-Bot-Api-Secret-Token
// SetWebhook https://core.telegram.org/bots/api#setwebhook
type SetWebhook struct {
	URL                string   `json:"url"`
	Certificate        *os.File `json:"certificate"`
	IpAddress          string   `json:"ip_address"`
	MaxConnections     uint8    `json:"max_connections"`
	AllowedUpdates     []string `json:"allowedUpdates"`
	DropPendingUpdates bool     `json:"dropPendingUpdates"`
	SecretToken        string   `json:"secretToken"`
}

// SetWebhookResp response for SetWebhook request
type SetWebhookResp struct {
	OK          bool   `json:"ok"`
	Result      bool   `json:"result"`
	Description string `json:"description"`
}

func (obj *Engine) SetWebhook(setup SetWebhook) (*SetWebhookResp, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	if err := w.WriteField("url", setup.URL); err != nil {
		return nil, err
	}

	if setup.Certificate != nil {
		fw, err := w.CreateFormFile("certificate", filepath.Base(setup.Certificate.Name()))
		if err != nil {
			return nil, err
		}
		if _, err = io.Copy(fw, setup.Certificate); err != nil {
			return nil, err
		}
	}

	if setup.IpAddress != "" {
		if err := w.WriteField("ip_address", setup.IpAddress); err != nil {
			return nil, err
		}
	}

	if setup.MaxConnections != 0 {
		if err := w.WriteField("max_connections", strconv.FormatUint(uint64(setup.MaxConnections), 10)); err != nil {
			return nil, err
		}
	}

	if len(setup.AllowedUpdates) != 0 {
		allowedUpdates, err := json.Marshal(setup.AllowedUpdates)
		if err != nil {
			return nil, err
		}
		if err := w.WriteField("allowed_updates", string(allowedUpdates)); err != nil {
			return nil, err
		}
	}

	if setup.DropPendingUpdates {
		if err := w.WriteField("drop_pending_updates", "true"); err != nil {
			return nil, err
		}
	}

	if setup.SecretToken != "" {
		if err := w.WriteField("secret_token", setup.SecretToken); err != nil {
			return nil, err
		}
	}

	if err := w.Close(); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", obj.telegramApiURL+obj.telegramBotToken+obj.telegramEnvironment+"/setWebhook", &b)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result SetWebhookResp

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
