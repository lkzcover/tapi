# tapi - simple golang package for the Telegram Bot API

This package may use for simple integration your application to the [Telegram API](https://core.telegram.org/bots/api)

For version v1.0.1 implement next API methods:
* [getupdates](https://core.telegram.org/bots/api#getupdates) - method for get new message from Telegram users
* [sendmessage](https://core.telegram.org/bots/api#sendmessage) - method for send message from bot to Telegram users or groups
* [editmessage](https://core.telegram.org/bots/api#editmessagetext) - method for edit message from your's bot

tapi support next Telegeam API object by default:
* [replykeyboardmarkup](https://core.telegram.org/bots/api#replykeyboardmarkup) - this object represents a custom keyboard with reply options
* [keyboardbutton](https://core.telegram.org/bots/api#keyboardbutton) - object represents one button of the reply keyboard
* [inlinekeyboardbutton](https://core.telegram.org/bots/api#inlinekeyboardbutton) - object represents one button of an inline keyboard


### Example

For connect to telegram api use:

```Golang
tgConn := tapi.Init("token from BotFather")
```

For check connect use

```Golang
getMsg, err := tgConn.GetMe()
```

For get new messages

```Golang
getMsgs, err := tgConn.GetUpdates()
```