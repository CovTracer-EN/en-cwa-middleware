name = en-cwa

run:
	go run main.go

build:
	go build -o ${name} main.go
	which upx && upx ${name}

test:
	go test -v ./...
