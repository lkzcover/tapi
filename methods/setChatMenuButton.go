package methods

// SetChatMenuButtonParams - implement https://core.telegram.org/bots/api#setchatmenubutton
type SetChatMenuButtonParams struct {
	ChatID     *int64       `json:"chat_id,omitempty"`     // Unique identifier for the target private chat. If not specified, default bot's menu button will be changed
	MenuButton *interface{} `json:"menu_button,omitempty"` // A JSON-serialized object for the bot's new menu button. Defaults to MenuButtonDefault
}
