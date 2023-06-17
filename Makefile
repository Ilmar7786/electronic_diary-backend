APP_BIN = build/app
ARGS = $(filter-out $@,$(MAKECMDGOALS))

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

.PHONY: generate-migration
generate-migration:
	migrate create -ext sql -dir migrations -seq $(name)