gql-tool:
	printf '//go:build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go

gqlgen:
	go run github.com/99designs/gqlgen generate

