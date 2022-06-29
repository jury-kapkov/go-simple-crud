package service

import "todo/pkg/repository"

type Authorisation interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorisation
	TodoList
	TodoItem
}

func NewService(r *repository.Repository) *Service {
	return &Service{}
}
