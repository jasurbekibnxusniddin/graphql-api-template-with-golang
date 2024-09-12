package resolvers

import "github.com/jasurbekibnxusniddin/graphql-api-template-with-golang/storage"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type resolver struct {
	Storage storage.StorageI
}

func NewResolver(storage storage.StorageI) *resolver {
	return &resolver{Storage: storage}
}
