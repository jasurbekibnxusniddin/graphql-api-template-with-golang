add-tool:
	mkdir -p pkgs/tools && printf '//go:build tools\npackage tools\nimport (\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > pkgs/tools/tools.go

