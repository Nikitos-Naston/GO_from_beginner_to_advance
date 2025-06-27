.PHONY: build

build:
	go build -v ./cmd/api/

up:
	migrate -path migrations -database "postgres://localhost:5432/restapi?sslmode=disable&user=postgres&password=2909" up

down:
	migrate -path migrations -database "postgres://localhost:5432/restapi?sslmode=disable&user=postgres&password=2909" down
.DEFAULT_GOAL := build