ifneq (,$(wildcard .env))
	include .env
	export
endif

MIGRATIONS_DIR ?= ./migrations
DATABASE_URL := postgres://$(DATABASE_USER):$(DATABASE_PASSWORD)@$(DATABASE_HOST):$(DATABASE_PORT)/$(DATABASE_NAME)?sslmode=disable
PG_CONTAINER_NAME := postgres-dev
TEST_DB_URL := postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable

.PHONY: go
go:
	@go run cmd/app/main.go

.PHONY: pg
pg:
	@docker run --name=postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d --rm postgres

.PHONY: psql
psql:
	docker exec -it $(PG_CONTAINER_NAME) psql -U $(DATABASE_USER) -d $(DATABASE_NAME)

.PHONY: migrate_up
migrate_up:
	migrate -path $(MIGRATIONS_DIR) -database '$(DATABASE_URL)' up

.PHONY: migrate_down
migrate_down:
	migrate -path $(MIGRATIONS_DIR) -database '$(DATABASE_URL)' down