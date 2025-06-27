.PHONY: build

build:
	go build -v ./cmd/api/

up:
	migrate -path migrations -database "postgres://localhost:5432/restapi?sslmode=disable&user=postgres&password=2909" up

down:
	migrate -path migrations -database "postgres://localhost:5432/restapi?sslmode=disable&user=postgres&password=2909" down

run_toml:
	./api.exe -path configs/api.toml -type toml 
run_env:
	./api.exe -path configs/.env -type env
.DEFAULT_GOAL := build