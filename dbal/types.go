package dbalpackage

const (
	Field_Gitlab_id = "gitlab_id"
	Field_Yt_id     = "youtrack_id"
)

type User struct {
	Id         uint64
	TelegramId string
	GitlabId   string
	YtId       string
	Step       int
}

type TelegramBotResult struct {
	GitlabCommand string
	YtCommand     string
	IsError       bool
	Error         error
}

func (u *User) GetId() uint64 {
	return u.Id
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) AssignTgId(id string) {
	u.TelegramId = id
}
