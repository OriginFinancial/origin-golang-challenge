setup:
	go mod tidy

run:
	go run ./src/main.go

build:
	mkdir -p ./dist
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/ ./src/main.go
