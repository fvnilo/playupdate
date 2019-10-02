.DEFAULT_GOAL := build_all

build_all: build_macosx build_linux build_windows

build_macosx:
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/mac/playupdate ./cmd/playupdate/main.go 

build_linux:
	GOOS=linux go build -ldflags="-s -w" -o ./bin/linux/playupdate ./cmd/playupdate/main.go

build_windows:
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/windows/playupdate.exe ./cmd/playupdate/main.go

