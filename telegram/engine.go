package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type Command struct {
	Key           string
	IsCallback    bool
	Message       Message
	CallbackQuery CallbackQuery
}

type CommandHandler interface {
	execute() BotMessage
	setCommand(command *Command)
}

type Engine struct {
	url      string
	commands map[string]CommandHandler
}

func NewEngine() Engine {
	e := Engine{}
	e.init()
	return e
}

func (e Engine) init() {
	e.url = TelegramUrl + Token
	e.commands = make(map[string]CommandHandler)
	e.commands[commandDefault] = &defaultHandler{}
	e.commands[commandBindUser] = &bindUserHandler{}
	e.commands[commandUpdateUser] = &updateUserHandler{}
	e.commands[commandStart] = &startHandler{}
	e.commands[commandAddGitlab] = &addGitlabHandler{}
	e.commands[commandAddYt] = &addYtHandler{}
}

func (e Engine) ReceiveHook(hook BotHook) BotMessage {
	command := e.prepareCommandFromHook(hook)

	handler, ok := e.commands[command.Key]
	if !ok {
		handler = e.commands[commandDefault]
	}

	handler.setCommand(&command)
	return handler.execute()
}

func (e Engine) SendMessage(message BotMessage) (status int, err error) {
	var response *http.Response

	requestData, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}

	response, err = http.Post(e.url+sendMessage, "application/json", bytes.NewBuffer(requestData))
	defer response.Body.Close()

	if err != nil {
		err = errors.New("Telegram request error:" + err.Error())
		return
	}
	if response.StatusCode > 399 || response.StatusCode < 200 {
		err = errors.New("Telegram")
	}
	status = response.StatusCode
	return
}

func (e Engine) prepareCommandFromHook(hook BotHook) Command {
	var command *Command = &Command{}
	targetMessage := hook.Message
	commandText := targetMessage.Text

	if hook.CallbackQuery.Data != "" {
		commandText = hook.CallbackQuery.Data
		command.IsCallback = true
		targetMessage = hook.CallbackQuery.Message
		command.CallbackQuery = hook.CallbackQuery
	}

	command.Message = targetMessage
	command.Key = e.parseCommandFromText(commandText)
	return *command
}

func (e Engine) parseCommandFromText(text string) string {
	command := strings.Split(text, " ")[0]
	return strings.TrimPrefix(command, "/")
}
