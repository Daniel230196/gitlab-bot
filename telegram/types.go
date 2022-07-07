package telegram

type BotMessage struct {
	ChatId      int         `json:"chat_id"`
	Text        string      `json:"text"`
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
	ParseMode   string      `json:"parse_mode,omitempty"`
}

type ReplyMarkup struct {
	Keyboard        [][]InlineKeyboard `json:"inline_keyboard,omitempty"`
	ResizeKeyboard  bool               `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard bool               `json:"one_time_keyboard,omitempty"`
}

type InlineKeyboard struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
}

type BotHook struct {
	Message       Message       `json:"message"`
	MessageId     int           `json:"message_id"`
	CallbackQuery CallbackQuery `json:"callback_query"`
	From          User          `json:"from"`
}

type CallbackQuery struct {
	Message Message `json:"message"`
	Text    string  `json:"text"`
	From    User    `json:"from"`
	Data    string  `json:"data"`
}

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type Message struct {
	Date string `json:"date"`
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

type Chat struct {
	LastName  string `json:"last_name"`
	Id        int    `json:"id"`
	FIrstName string `json:"f_irst_name"`
	UserName  string `json:"user_name"`
}

type InlineQuery struct {
	Id        int    `json:"id"`
	From      User   `json:"from"`
	Query     string `json:"query"`
	Offset    string `json:"offset"`
	Chat_type string `json:"chat_type"`
}
