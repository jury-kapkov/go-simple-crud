package service

import (
	"todo"
	"todo/pkg/repository"
)

type Authorisation interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	CreateTodoList(userId int, todoList todo.TodoList) (int, error)
}

type TodoItem interface {
}

type Service struct {
	Authorisation
	TodoList
	TodoItem
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Authorisation: NewAuthService(r.Authorisation),
		TodoList:      NewTodoListService(r.TodoList),
	}
}
