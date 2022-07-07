package telegram

const (
	step_user_created = iota
	step_wait_updates
	step_wait_gitlab_token
	step_wait_yt_token
	step_fullified
)

const (
	TelegramUrl = "https://api.telegram.org/"
	Token       = "bot5534393155:AAGQxeWphL8BJGE3PiM1MLrO1rHeqI7MFaM/"
	sendMessage = "sendMessage"

	commandBindUser   = "bind_user"
	commandUpdateUser = "update_user"
	commandDefault    = "default"
	commandStart      = "start"
	commandAddGitlab  = "add_gitlab"
	commandAddYt      = "add_yt"
)

func CreateInlineButton(text string, callbackData string) InlineKeyboard {
	return InlineKeyboard{text, callbackData}
}
