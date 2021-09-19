.DEFAULT_GOAL := lint



lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix

test:
	mkdir -p coverage
	go test ./... -coverprofile coverage/cover.out
	go tool cover -html=coverage/cover.out -o=coverage/cover.html

ci: lint test