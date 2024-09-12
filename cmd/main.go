package main

import (
	"github.com/jasurbekibnxusniddin/graphql-api-template-with-golang/gql"
	"github.com/jasurbekibnxusniddin/graphql-api-template-with-golang/storage"
)

func main() {

	storage := storage.NewStorage()

	gql.Gql(
		gql.Options{
			Storage: storage,
		},
	)
}
