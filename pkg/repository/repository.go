package repository

import (
	"github.com/jmoiron/sqlx"
	"todo"
)

type Authorisation interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username string, password string) (todo.User, error)
}

type TodoList interface {
	CreateTodoList(userId int, list todo.TodoList) (int, error)
}

type TodoItem interface {
}

type Repository struct {
	Authorisation
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorisation: NewAuth(db),
		TodoList:      NewTodoList(db),
	}
}
