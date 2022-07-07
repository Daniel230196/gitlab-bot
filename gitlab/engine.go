package gitlab

import (
	"fmt"
	dbalpackage "github.com/pos-gitlab/dbal"
	"os"
	"strconv"
)

const (
	EventMr      = "merge_request"
	EventComment = "note"
)

type Event interface {
	prepareMessage() string
	getAuthorId() string
}

type Engine struct{}

type GitlabResult struct {
	TelegramChatId int
	Text           string
	isNeedToSend   bool
}

func NewEngine() Engine {
	return Engine{}
}

type EventAttributes struct {
	Id          int
	MrAuthor    int
	Url         string
	Source      string
	Target      string
	Title       string
	State       string
	Description string
	Type        string
}

var repository = dbalpackage.NewBotRepository()

func (e Engine) RecieveHook(hook GitlabHook) GitlabResult {
	var err error
	eventAttributes := e.convertHookDataToEventAttributes(hook)
	event := createEventByType(hook.EventType, eventAttributes)
	user := repository.FindUserByGitlabId(event.getAuthorId())
	result := GitlabResult{0, "", false}

	if event == nil {
		result.isNeedToSend = false
		return result
	}
	if user.GetId() == 0 {
		return result
	}

	result.isNeedToSend = true
	result.Text = event.prepareMessage()

	result.TelegramChatId, err = strconv.Atoi(user.TelegramId)
	fmt.Fprint(os.Stderr, result.TelegramChatId)

	if err != nil {
		// TODO
		panic(err)
	}
	return result
}

func (e Engine) convertHookDataToEventAttributes(hook GitlabHook) EventAttributes {
	var result EventAttributes
	result.Url = hook.ObjectAttributes.Url
	result.Id = hook.ObjectAttributes.Id

	switch hook.EventType {
	case EventComment:
		result.Type = hook.ObjectAttributes.NoteableType
		result.Source = hook.MergeRequest.SourceBranch
		result.Target = hook.MergeRequest.TargetBranch
		result.MrAuthor = hook.MergeRequest.AuthorId
		result.Description = hook.ObjectAttributes.Note
	case EventMr:
		result.Target = hook.ObjectAttributes.TargetBranch
		result.Source = hook.ObjectAttributes.SourceBranch
		result.Title = hook.ObjectAttributes.Title
		result.Description = hook.ObjectAttributes.Description
		result.State = hook.ObjectAttributes.State
		result.MrAuthor = hook.ObjectAttributes.AuthorId
	}

	return result
}

func (res GitlabResult) IsSend() bool {
	return res.isNeedToSend
}
