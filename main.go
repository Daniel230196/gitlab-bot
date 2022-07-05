package main

import (
	"encoding/json"
	dbalpackage "github.com/pos-gitlab/dbal"
	"github.com/pos-gitlab/gitlab"
	"github.com/pos-gitlab/telegram"
	"net/http"
)

var TgEngine telegram.Engine = telegram.NewEngine()
var GitLabEngine gitlab.Engine = gitlab.NewEngine()

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/tg-hook", handleTgHook)
	mux.HandleFunc("/gitlab-hook", handleGitlabHook)
	http.ListenAndServe(":8080", mux)
}

func handleTgHook(w http.ResponseWriter, r *http.Request) {
	hook := &telegram.BotHook{}
	json.NewDecoder(r.Body).Decode(hook)

	TgEngine.ReceiveHook(*hook)
}

func handleGitlabHook(w http.ResponseWriter, r *http.Request) {
	hook := &gitlab.GitlabHook{}
	json.NewDecoder(r.Body).Decode(hook)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(hook.ObjectAttributes.State))
}

var rep = dbalpackage.NewBotRepository()

func requestWithTimeOut(t int64, endpoint string) {

}
