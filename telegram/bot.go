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
	Token       = "BOT_TOKEN"
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
