package tapi

type replyMsgStruct struct {
	ChatID           int64       `json:"chat_id"`
	Text             string      `json:"text"`
	ReplyToMessageID *uint64     `json:"reply_to_message_id,omitempty"`
	ReplyMarkup      interface{} `json:"reply_markup,omitempty"`
}

//ReplyKeyboardMarkup https://core.telegram.org/bots/api#replykeyboardmarkup
type ReplyKeyboardMarkup struct {
	Keyboard        [][]KeyboardButton `json:"keyboard"`
	ResizeKeyboard  bool               `json:"resize_keyboard"`
	OneTimeKeyboard bool               `json:"one_time_keyboard"`
	Selective       bool               `json:"selective"`
}

//KeyboardButton https://core.telegram.org/bots/api#keyboardbutton
type KeyboardButton struct {
	Text string `json:"text"`
}
