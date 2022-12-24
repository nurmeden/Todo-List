package service

import "todo/pkg/repository"

type (
	Authorization interface{}
	TodoList      interface{}
	TodoItem      interface{}
)

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service { // конструктор
	return &Service{}
}
