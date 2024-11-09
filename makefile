#!make

migrator:
	go build -o scripts/migrator ./scripts/goose/main.go
