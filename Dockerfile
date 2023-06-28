FROM golang:alpine as builder
WORKDIR /app

RUN apk add make

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN make build

FROM alpine:latest
WORKDIR /app

COPY --from=builder /build/app .
COPY --from=builder /app/.env .

ENTRYPOINT ["./app"]