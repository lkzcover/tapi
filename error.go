package tapi

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	InternalLibError = errors.New("internal lib error")
	MigrateChatID    = errors.New("migrate_chat_id")
)

// {"ok":false,"error_code":400,"description":"Bad Request: group chat was upgraded to a supergroup chat","parameters":{"migrate_to_chat_id":-1002037472066}}

// Error - error message from Telegram
type Error struct {
	Ok          bool        `json:"ok"`
	ErrorCode   uint16      `json:"error_code"`
	Description string      `json:"description"`
	Parameters  *Parameters `json:"parameters"`
}

type Parameters struct {
	MigrateToChatID int64 `json:"migrate_to_chat_id"`
}

func parseError(body []byte) (*Error, error) {
	var tgError Error

	err := json.Unmarshal(body, &tgError)
	if err != nil {
		return nil, fmt.Errorf("%s %w", InternalLibError, err)
	}

	if tgError.ErrorCode == 400 {
		if tgError.Parameters != nil && tgError.Parameters.MigrateToChatID != 0 {
			return &tgError, MigrateChatID
		}
	}

	return &tgError, fmt.Errorf("%d %w", tgError.ErrorCode, errors.New(tgError.Description))
}
