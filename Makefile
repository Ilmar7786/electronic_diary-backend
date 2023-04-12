APP_BIN = build/app

.PHONY: run
run:
	go run ./cmd/app/main.go

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
	swag init -g ./app/main.go -o ./app/docs

.PHONY: docker-up
docker-up:
	docker-compose up
