BINARY_NAME=steam-update-notifier

.PHONY: test

build:
	go build -o bin/${BINARY_NAME} cmd/steam-update-notifier/main.go
clean:
	go clean
	rm bin/${BINARY_NAME}*
test:
	go test -v -coverpkg ./... ./test/...
release:
	GOOS=darwin GOARCH=arm64 go build -o bin/${BINARY_NAME}-darwin-arm64 cmd/steam-update-notifier/main.go
	GOOS=linux GOARCH=arm64 go build -o bin/${BINARY_NAME}-linux-arm64 cmd/steam-update-notifier/main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/${BINARY_NAME}-darwin-amd64 cmd/steam-update-notifier/main.go
	GOOS=linux GOARCH=amd64 go build -o bin/${BINARY_NAME}-linux-amd64 cmd/steam-update-notifier/main.go
	GOOS=windows GOARCH=amd64 go build -o bin/${BINARY_NAME}-windows-amd64 cmd/steam-update-notifier/main.go
	GOOS=windows GOARCH=386 go build -o bin/${BINARY_NAME}-windows-386 cmd/steam-update-notifier/main.go
