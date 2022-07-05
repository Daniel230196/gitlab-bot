package telegram

import (
	"fmt"
	dbalpackage "github.com/pos-gitlab/dbal"
	"strconv"
	"strings"
)

type defaultHandler struct {
	*Command
	answer *BotMessage
	user   User
}

type bindUserHandler struct {
	defaultHandler
}

type updateUserHandler struct {
	defaultHandler
}

type startHandler struct {
	defaultHandler
}

type addGitlabHandler struct {
	defaultHandler
}

type addYtHandler struct {
	defaultHandler
}

var repository = dbalpackage.NewBotRepository()

func (h *defaultHandler) execute() BotMessage {
	h.initMessage()

	user := repository.FindUserByChatId(h.chatId())
	if user.GetId() == 0 {
		user.AssignTgId(strconv.Itoa(h.chatId()))
		user.Step = step_user_created
		repository.CreateUser(user)
	}

	h.answer.ChatId = h.chatId()
	h.answer.Text = fmt.Sprintf("POSCREDIT DEVELOPMENT BOT")

	h.processStep(user)
	return *h.answer
}

func (h *defaultHandler) setCommand(command *Command) {
	h.Command = command
}

func (h *defaultHandler) initMessage() {
	testMessage := BotMessage{}
	testMessage.ChatId = h.chatId()
	h.answer = &BotMessage{}
}
func (h *bindUserHandler) execute() BotMessage {
	h.initMessage()
	h.answer.Text = "test bind"
	return *h.answer
}

func (h *updateUserHandler) execute() BotMessage {
	h.initMessage()
	h.answer.Text = "test update"
	return *h.answer
}

func (h *startHandler) execute() BotMessage {
	h.initMessage()
	user := repository.FindUserByChatId(h.chatId())
	if user.GetId() == 0 {
		user.AssignTgId(strconv.Itoa(h.chatId()))
		user.Step = step_user_created
		repository.CreateUser(user)
	}

	h.addKeyBoardIfNeeded(user)

	h.answer.ChatId = h.chatId()
	h.answer.Text = fmt.Sprintf("Bind your Yt and Gitlab ids")

	return *h.answer
}

func (h *addYtHandler) execute() BotMessage {
	h.initMessage()
	h.processUserFieldUpdate(dbalpackage.Field_Yt_id)

	return *h.answer
}

func (h *addGitlabHandler) execute() BotMessage {
	h.initMessage()
	h.processUserFieldUpdate(dbalpackage.Field_Gitlab_id)

	return *h.answer
}

func (h *defaultHandler) processUserFieldUpdate(fieldName string) {
	user := repository.FindUserByChatId(h.chatId())
	if user.GetId() == 0 {
		h.writeDefaultAnswer()
		return
	}
	h.answer.Text = "Введите ваш " + fieldName
	h.answer.ChatId = h.chatId()

	var step int
	switch fieldName {
	case dbalpackage.Field_Gitlab_id:
		step = step_wait_gitlab_token
	case dbalpackage.Field_Yt_id:
		step = step_wait_yt_token
	}
	user.Step = step
	repository.UpdateUser(user)
}

func (h *defaultHandler) writeDefaultAnswer() {
	h.answer.Text = fmt.Sprintf("Welcome to POSCREDIT dev bot!")
}

func (h *defaultHandler) processStep(user *dbalpackage.User) {
	switch user.Step {
	case step_wait_gitlab_token:
		user.GitlabId = strings.Trim(h.Message.Text, " ")
		h.answer.Text = fmt.Sprintf("Gitlab id установлен: %v", user.GitlabId)
	case step_wait_yt_token:
		user.YtId = strings.Trim(h.Message.Text, " ")
		h.answer.Text = fmt.Sprintf("YouTrack id установлен: %v", user.YtId)
	case step_wait_updates:
	case step_user_created:
	default:
		break
	}

	user.Step = calculateUserStep(*user)
	repository.UpdateUser(user)
	h.addKeyBoardIfNeeded(user)
}

func (h *defaultHandler) addInlineButtonsRow(keyboard []InlineKeyboard) {
	line := h.answer.ReplyMarkup.Keyboard
	markUp := []InlineKeyboard{}
	for _, button := range keyboard {
		markUp = append(markUp, button)
	}
	h.answer.ReplyMarkup.Keyboard = append(line, markUp)
}

func (h *defaultHandler) addKeyBoardIfNeeded(user *dbalpackage.User) {
	var inlineButtons []InlineKeyboard
	inlineButtons = append(inlineButtons, h.prepareButtonForUserFieldUpdate(dbalpackage.Field_Gitlab_id, user.GitlabId, "/add_gitlab"))
	inlineButtons = append(inlineButtons, h.prepareButtonForUserFieldUpdate(dbalpackage.Field_Yt_id, user.YtId, "/add_yt"))
	h.addInlineButtonsRow(inlineButtons)
	h.answer.ReplyMarkup.ResizeKeyboard = true
	h.answer.ReplyMarkup.OneTimeKeyboard = true
}

func (h *defaultHandler) prepareButtonForUserFieldUpdate(fieldName string, fieldValue string, callBackData string) InlineKeyboard {
	var buttonText string
	var textPrefix string
	if fieldValue == "" {
		textPrefix = "Добавить"
	} else {
		textPrefix = "Обновить"
	}

	buttonText = textPrefix + " " + fieldName
	return InlineKeyboard{buttonText, callBackData}
}

func (h *defaultHandler) chatId() int {
	if h.Command.IsCallback {
		return h.CallbackQuery.Message.Chat.Id
	}

	return h.Message.Chat.Id
}

func calculateUserStep(user dbalpackage.User) int {
	var calculatedStep int = step_user_created
	if user.YtId == "" && user.GitlabId == "" {
		calculatedStep = step_user_created
	}
	if user.YtId == "" || user.GitlabId == "" {
		calculatedStep = step_wait_updates
	}

	return calculatedStep
}
