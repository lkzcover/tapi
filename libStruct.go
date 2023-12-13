package tapi

const (
	MsgFormatMarkDown = Formatting("MarkdownV2")
)

// Formatting defining text style
type Formatting string

// MsgBody struct for configuration message style
type MsgBody struct {
	Text   string     // message text
	Format Formatting // format message
}
