package tapi

const (
	MsgFormatMarkDown = Formatting("MarkdownV2")
)

// Formatting defining text style
type Formatting string

// MsgParams struct for message parameters
type MsgParams struct {
	Text            string     // message text
	Format          Formatting // format message
	MigrateToChatID bool       // set true for automatic redirection message to new chat id
}
