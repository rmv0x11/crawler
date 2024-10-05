run:
	go mod tidy && go mod download && \
	GIN_MODE=debug CGO_ENABLED=0 go run ./cmd
.PHONY: run

test: ### run test
	go test -v -cover -race ./internal/...
.PHONY: test