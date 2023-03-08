SHELL=/bin/bash
SERVICE_NAME := $(notdir $(shell pwd))

.PHONY: download
mod:
	go mod download
	go mod tidy -compat=1.20

.PHONY: swag
swag:
	swag init -g cmd/main.go
	cp ./docs/swagger.json ./docs/$(SERVICE_NAME)-openapi.json
	cp ./docs/swagger.yaml ./docs/$(SERVICE_NAME)-openapi.yaml
	swag fmt

.PHONY: mocks
mocks:
	rm -fr ./mocks
	mockery --all --dir ./internal --disable-version-string --case snake --keeptree

.PHONY: lint
lint:
	golangci-lint -v run

.PHONY: format
format:
	goimports -w ./
	go fmt ./...

.PHONY: changelog
changelog:
	git-chglog -o CHANGELOG.md

.PHONY: docker-build
docker-build:
	docker build -t system-tweet .

.PHONY: compose-up
compose-up: docker-build
	docker-compose --env-file .env up -d

.PHONY: compose-stop
compose-stop:
	docker-compose stop

.PHONY: compose-down
compose-down:
	docker-compose down

.PHONY: compose-test
compose-test:
	docker-compose run --rm test