all: tidy test-color lint

tidy:
    go mod tidy

test:
    go test -race -failfast ./...

test-color:
    # go install github.com/haunt98/go-test-color@latest
    go-test-color -race -failfast ./...

coverage:
    go test -coverprofile=coverage.out ./...

coverage-cli: coverage
    go tool cover -func=coverage.out

coverage-html: coverage
    go tool cover -html=coverage.out

lint:
    golangci-lint run --fix ./...
    modernize -fix -test ./...
