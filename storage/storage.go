package storage

import "github.com/jasurbekibnxusniddin/graphql-api-template-with-golang/storage/postgres"

type StorageI interface {
	GetTodoRepo() postgres.TodoRepoI
}

type storage struct {
	todoRepo postgres.TodoRepoI
}

func NewStorage() StorageI {
	return &storage{
		todoRepo: postgres.NewTodoRepo(),
	}
}

func (s *storage) GetTodoRepo() postgres.TodoRepoI {
	return s.todoRepo
}
