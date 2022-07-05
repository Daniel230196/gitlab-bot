package dbalpackage

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"sync"
	"time"
)

type botRepository struct {
	connection *sql.DB
}

func (r *botRepository) Connection() *sql.DB {
	return r.connection
}

func NewBotRepository() (repository botRepository) {
	repository = botRepository{}
	repository.init()
	return
}

func (r *botRepository) init() {
	var once sync.Once
	once.Do(func() {
		conn, err := sql.Open("mysql", "root:123123123@tcp(127.0.0.1:3306)/poscredit")
		if err != nil {
			// TODO
			panic(err)
		}

		r.connection = conn
	})
}

func (r *botRepository) CreateUser(user *User) {
	stmt, err := r.connection.Prepare("REPLACE INTO users (telegram_id, gitlab_id, yt_id, step) VALUES (?, ?, ?, ?)")
	defer stmt.Close()

	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	result, err := stmt.ExecContext(ctx, user.TelegramId, user.GitlabId, user.YtId, user.Step)
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	user.Id = uint64(id)
}

func (r *botRepository) UpdateGitlabIdById(gitlabId string, userId uint64) {
	var user *User = &User{}
	user.Id = userId
	user.GitlabId = gitlabId

	r.UpdateUser(user)
}

func (r *botRepository) UpdateYtIdById(ytId string, userId uint64) {
	var user *User = &User{}
	user.Id = userId
	user.YtId = ytId

	r.UpdateUser(user)
}

func (r *botRepository) FindUserByChatId(chatId int) *User {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	rows, err := r.connection.QueryContext(ctx, "SELECT * FROM users WHERE `telegram_id`=?", chatId)
	if err != nil {
		panic(err)
	}

	var user Entity = &User{}
	r.mapRowsOnEntity(rows, &user)
	return user.(*User)
}

func (r *botRepository) UpdateUser(user *User) *User {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	res, err := r.connection.QueryContext(ctx, "UPDATE users SET `yt_id`=?,`gitlab_id`=?,`step`=? WHERE `id`=?;", user.YtId, user.GitlabId, user.Step, user.Id)
	if err != nil {
		panic(err)
	}

	var uEntity Entity = &User{}
	r.mapRowsOnEntity(res, &uEntity)
	return uEntity.(*User)
}

func (r *botRepository) mapRowsOnEntity(rows *sql.Rows, entity *Entity) *Entity {
	for {
		s := reflect.ValueOf(*entity).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}

		rows.Scan(columns...)
		if !rows.Next() {
			break
		}
	}

	return entity
}

type Entity interface {
	GetId() uint64
	TableName() string
}

type EntityCollection interface {
	list() []Entity
}
