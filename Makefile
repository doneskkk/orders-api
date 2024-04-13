build:
	@go build -o bin/donesk cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/donesk