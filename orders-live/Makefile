run:
	go run cmd/main.go

test:
	go test ./... -cover -coverprofile=coverage.out

gen-mocks:
	mockery --dir=./internal --all

# tidies dependencies
tidy-deps:
	go mod tidy

# vendors dependencies
deps:
	go mod tidy
	go mod vendor

# formats project with go's style guidelines
fmt:
	gofmt -w -s ./src

# checks the project for possible errors
check:
	staticcheck ./...
