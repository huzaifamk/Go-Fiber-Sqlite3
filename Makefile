.PHONY: build

run:
	go run ./cmd/main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o ./bin/main ./cmd/main.go

build-macos:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/main ./cmd/main.go

build-windows:
	GOOS=windows GOARCH=amd64 go build -o ./bin/main.exe ./cmd/main.go
