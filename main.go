package main

import (
	"encoding/json"
	"fmt"
	dbalpackage "github.com/pos-gitlab/dbal"
	"github.com/pos-gitlab/gitlab"
	"github.com/pos-gitlab/telegram"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/tg-hook", handleTgHook)
	mux.HandleFunc("/gitlab-hook", handleGitlabHook)
	http.ListenAndServe(":8080", mux)
}

var TgEngine *telegram.Engine = telegram.NewEngine()
var GitLabEngine gitlab.Engine = gitlab.NewEngine()

func handleTgHook(w http.ResponseWriter, r *http.Request) {
	hook := &telegram.BotHook{}
	json.NewDecoder(r.Body).Decode(hook)

	message := TgEngine.ReceiveHook(*hook)
	TgEngine.SendMessage(message)
}

func handleGitlabHook(w http.ResponseWriter, r *http.Request) {
	hook := &gitlab.GitlabHook{}
	json.NewDecoder(r.Body).Decode(hook)
	result := GitLabEngine.RecieveHook(*hook)

	if result.IsSend() {
		message := telegram.BotMessage{}
		message.ChatId = result.TelegramChatId
		message.Text = result.Text
		message.ParseMode = "markdown"
		fmt.Fprint(os.Stderr, message.ChatId)
		fmt.Fprint(os.Stderr, "\n"+message.Text)
		fmt.Fprint(os.Stderr, "\n"+message.ParseMode)

		TgEngine.SendMessage(message)
	}
}

var rep = dbalpackage.NewBotRepository()

func requestWithTimeOut(t int64, endpoint string) {

}
