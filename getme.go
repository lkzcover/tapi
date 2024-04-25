package tapi

import (
	"github.com/lkzcover/tapi/handlers"
)

type GetMeStruct struct {
	OK     bool `json:"ok"`
	Result struct {
		ID                      uint64 `json:"id"`
		IsBot                   bool   `json:"is_bot"`
		FirstName               string `json:"first_name"`
		Username                string `json:"username"`
		CanJoinGroups           bool   `json:"can_join_groups"`
		CanReadAllGroupMessages bool   `json:"сan_read_all_group_messages"`
		SupportsInlineQueries   bool   `json:"supports_inline_queries"`
	} `json:"result"`
}

func (obj *Engine) GetMe() (*GetMeStruct, error) {

	var getMe GetMeStruct

	err := handlers.GetRequest(obj.telegramApiURL+obj.telegramBotToken+obj.telegramEnvironment+"/getMe", &getMe)

	return &getMe, err
}
