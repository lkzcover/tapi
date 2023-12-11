package tapi

type MessageList struct {
	OK     bool      `json:"ok"`
	Result []Message `json:"result"`
}

// ResultMsg reply message from telegram api
type ResultMsg struct {
	Ok     bool `json:"ok"`
	Result struct {
		MessageId int `json:"message_id"`
		From      struct {
			Id        int64  `json:"id"`
			IsBot     bool   `json:"is_bot"`
			FirstName string `json:"first_name"`
			Username  string `json:"username"`
		} `json:"from"`
		Chat struct {
			Id        int    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Username  string `json:"username"`
			Type      string `json:"type"`
		} `json:"chat"`
		Date int    `json:"date"`
		Text string `json:"text"`
	} `json:"result"`
}

type Message struct {
	UpdateID int64 `json:"update_id"`

	Message MessageStruct `json:"message"`

	CallbackQuery *struct {
		From struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"from"`
		Message MessageStruct `json:"message"`
		Data    *string       `json:"data,omitempty"`
	} `json:"callback_query,omitempty"`
}

type MessageStruct struct {
	MessageID uint64 `json:"message_id"`
	From      struct {
		ID       int64  `json:"id"`
		Username string `json:"username"`
	} `json:"from"`
	Text string `json:"text"`
	Chat struct {
		ID int64 `json:"id"`
	} `json:"chat"`
}

type replyMsgStruct struct {
	ChatID           int64       `json:"chat_id"`
	MessageID        *uint64     `json:"message_id"`
	Text             string      `json:"text"`
	ReplyToMessageID *uint64     `json:"reply_to_message_id,omitempty"`
	ReplyMarkup      interface{} `json:"reply_markup,omitempty"`
}

// ReplyKeyboardMarkup https://core.telegram.org/bots/api#replykeyboardmarkup
type ReplyKeyboardMarkup struct {
	Keyboard        [][]KeyboardButton `json:"keyboard"`
	ResizeKeyboard  bool               `json:"resize_keyboard"`
	OneTimeKeyboard bool               `json:"one_time_keyboard"`
	Selective       bool               `json:"selective"`
}

// KeyboardButton https://core.telegram.org/bots/api#keyboardbutton
type KeyboardButton struct {
	Text string `json:"text"`
}

// InlineKeyboardMarkup InlineKeyboardMarkup
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

// InlineKeyboardButton https://core.telegram.org/bots/api#inlinekeyboardbutton
type InlineKeyboardButton struct {
	Text string `json:"text"`
	// Url  *string `json:"url,omitempty"` // Optional. HTTP or tg:// url to be opened when button is pressed
	// login_url
	CallbackData *string `json:"callback_data,omitempty"`
	// switch_inline_query
	// switch_inline_query_current_chat
	// callback_game
	// pay
}
