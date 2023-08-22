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


.DEFAULT_GOAL := run