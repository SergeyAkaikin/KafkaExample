build-producer:
	go build -tags musl -o bin/producer ./cmd/producer/main.go

build-consumer:
	go build -tags musl -o bin/consumer ./cmd/consumer/main.go