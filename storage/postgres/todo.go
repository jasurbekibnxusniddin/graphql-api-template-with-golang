package postgres

import (
	"context"

	"github.com/jasurbekibnxusniddin/graphql-api-template-with-golang/models"
)

type todoRepo struct {
}

func NewTodoRepo() TodoRepoI {
	return &todoRepo{}
}

// -
func (t *todoRepo) CreateTodo(ctx context.Context, todo *models.Todo) (*models.Todo, error) {

	todos = append(todos, todo)

	return todo, nil
}

// -
func (t *todoRepo) GetTodo(ctx context.Context) ([]*models.Todo, error) {
	
	return todos, nil
}
