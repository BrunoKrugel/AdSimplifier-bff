.PHONY: run

run:
	go build ./cmd/go-webhook
	go run ./cmd/go-webhook