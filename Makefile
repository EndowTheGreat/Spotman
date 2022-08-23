build:
	@go build -o bin/spotman -ldflags "-s -w" cmd/control/main.go