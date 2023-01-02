.PHONY: build

# Use the command `make run` in the terminal to run the application directly
run:
	go run ./cmd/main.go

# Use the command `make build` in the terminal to build the application
build:
	go build -o ./bin/gofiber-app ./cmd/main.go

# Use the command `make run-build` in the terminal to run the build of the application
run-build:
	./bin/gofiber-app

# Use the command `make build-linux` in the terminal to build the application for Linux OS
build-linux:
	GOOS=linux GOARCH=amd64 go build -o ./bin/gofiber-app ./cmd/main.go

# Use the command `make build-macos` in the terminal to build the application for MacOS
build-macos:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/gofiber-app ./cmd/main.go

# Use the command `make build-windows` in the terminal to build the application for Windows OS
build-windows:
	GOOS=windows GOARCH=amd64 go build -o ./bin/gofiber-app.exe ./cmd/main.go