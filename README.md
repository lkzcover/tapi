# tapi - simple golang package for the Telegram Bot API

This package may use for simple integration your application to the [Telegram API](https://core.telegram.org/bots/api)

For version v2.2.0 implement next API methods:
* [getMe](https://core.telegram.org/bots/api#getme) - method for test connection to Telegram API
* [getUpdates](https://core.telegram.org/bots/api#getupdates) - method for get new message from Telegram users
* [sendMessage](https://core.telegram.org/bots/api#sendmessage) - method for send message from bot to Telegram users or groups
* [editMessageText](https://core.telegram.org/bots/api#editmessagetext) - method for edit message from your's bot* [deleteMessages](https://core.telegram.org/bots/api#deletemessages) - method for delete messages in a chat or channel
* [setWebhook](https://core.telegram.org/bots/api#setwebhook) - method for get new message from Telegram by webhook
* [getWebhookInfo](https://core.telegram.org/bots/api#getwebhookinfo) - method for get information about set a webhook connection
* [deleteWebhook](https://core.telegram.org/bots/api#deletewebhook) - method for disabled Webhook connection
* [answerInlineQuery](https://core.telegram.org/bots/api#answerinlinequery) - method for send answers to an inline query

tapi support next Telegeam API object by default:
* [replykeyboardmarkup](https://core.telegram.org/bots/api#replykeyboardmarkup) - this object represents a custom keyboard with reply options
* [keyboardbutton](https://core.telegram.org/bots/api#keyboardbutton) - object represents one button of the reply keyboard
* [inlinekeyboardbutton](https://core.telegram.org/bots/api#inlinekeyboardbutton) - object represents one button of an inline keyboard


### Example

For connect to telegram api use:

```go
tgConn := tapi.Init("token from BotFather")
```

For check connect use

```go
getMsg, err := tgConn.GetMe()
```

For get new messages

```go
getMsgs, err := tgConn.GetUpdates()
```

Or you can use webhook

```go
// set connection
resp, err := tgConn.SetWebhook(tapi.SetWebhook{
		URL: <address to to your server for get new message from Telegram>,
		Certificate: <your public key certificate if you used self-signed SSL certifucate>,
	})

// get webhook connection info
resp, err := tgConn.GetWebhookInfo()

// disabled webhook connection
resp, err := tgConn.DeleteWebhook()
```

After setup connection you can send message to chat in Telegram
```go
// simple message
resp, err := tgConn.SendMessage(<chat_id>, tapi.MsgParams{Text: <message text>})

// message with formating
resp, err := tgConn.SendMessage(<chat_id>, tapi.MsgParams{Text: <message text>, Format: tapi.MsgFormatMarkDown})

// if Telegram return error "migrate_chat_id" you can use additional parameter: MigrateToChatID = true for auto resend message to new chat_id
resp, err := tgConn.SendMessage(<chat_id>, tapi.MsgParams{Text: <message text>, MigrateToChatID: true, Format: tapi.MsgFormatMarkDown})
```


