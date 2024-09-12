package postgres

import (
	"context"

	"github.com/jasurbekibnxusniddin/graphql-api-template-with-golang/models"
)

type TodoRepoI interface {
	CreateTodo(ctx context.Context, todo *models.Todo) (*models.Todo, error)
	GetTodo(ctx context.Context) ([]*models.Todo, error)
}
