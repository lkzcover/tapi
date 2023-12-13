package tapi

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	InternalLibError = errors.New("internal lib error")
)

// Error - error message from Telegram
type Error struct {
	Ok          bool   `json:"ok"`
	ErrorCode   uint16 `json:"error_code"`
	Description string `json:"description"`
}

func parseError(body []byte) error {
	var tgError Error

	err := json.Unmarshal(body, &tgError)
	if err != nil {
		return errors.Join(InternalLibError, err)
	}

	return fmt.Errorf("%d %w", tgError.ErrorCode, errors.New(tgError.Description))
}
