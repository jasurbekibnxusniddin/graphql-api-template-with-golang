package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jasurbekibnxusniddin/graphql-api-template-with-golang/gql/resolvers"
	"github.com/jasurbekibnxusniddin/graphql-api-template-with-golang/gql/schema"
)

func main() {

	srv := handler.NewDefaultServer(schema.NewExecutableSchema(schema.Config{Resolvers: &resolvers.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", "8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
