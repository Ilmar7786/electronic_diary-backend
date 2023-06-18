include .env

APP_BIN = build/app
ARGS = $(filter-out $@,$(MAKECMDGOALS))
DATABASE_URL = "postgresql://${PSQL_USER}:${PSQL_PASSWORD}@${PSQL_HOST}:${PSQL_PORT}/${PSQL_DATABASE}?sslmode=disable"

.PHONY: run
run:
	go run ./cmd/app/main.go $(ARGS)

.PHONY: lint
lint:
	golangci-lint run

.PHONY: build
build: clean $(APP_BIN)

$(APP_BIN):
	go build -o $(APP_BIN) ./cmd/app/main.go

.PHONY: clean
clean:
	rm -rf ./build || true

.PHONY: swagger
swagger:
	swag init -g ./cmd/app/main.go -o ./docs

.PHONY: docker-up
docker-up:
	docker-compose up

.PHONY: migration-generate
migration-generate:
	migrate create -ext sql -dir migrations -seq $(name)

.PHONY: migration
migration:
	migrate -source file://migrations -database $(DATABASE_URL) $(ARGS)

