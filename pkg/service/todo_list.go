package service

import (
	"todo"
	"todo/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (t TodoListService) CreateTodoList(userId int, todoList todo.TodoList) (int, error) {
	return t.repo.CreateTodoList(userId, todoList)
}
