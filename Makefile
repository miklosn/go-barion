build:
	go build -o bin/main cmd/main.go

run:
	go run cmd/main.go

generate:
	go generate ./...

test:
	go test ./...

