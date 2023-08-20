.PHONY: run
run:
	docker compose down --remove-orphans
	docker compose --profile postgres up --build

.PHONY: local
local:
	go run cmd/main.go

.PHONY: stop
stop:
	docker compose down --remove-orphans

.PHONY: test
test:
	go test -v -race -cover ./...

.PHONY: generate
generate:
	go generate ./...

.DEFAULT_GOAL := run