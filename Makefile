.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./hello-world/hello-world
	
build:
	GOOS=linux GOARCH=amd64 go build -o hello-world/hello-world ./hello-world

build-name:
	GOOS=linux GOARCH=amd64 go build -o hello-name/hello-name ./hello-name
build-message:
	GOOS=linux GOARCH=amd64 go build -o get-message/get-message ./get-message
	GOOS=linux GOARCH=amd64 go build -o send-message/send-message ./send-message
