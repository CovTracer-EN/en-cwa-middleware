#!make

name = en-cwa

upx := $(shell which upx)

run:
	go run main.go

build:
	go build -ldflags "-s -w" -o ${name} .
ifdef upx
	upx ${name}
endif

test:
	go test -v ./...
