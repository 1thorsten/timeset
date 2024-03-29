update_dependencies:
	go get -u -t ./...

build: bin
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o bin/timeset.exe ./src

bin:
	mkdir -p bin/

install: bin/timeset
	GOOS=windows GOARCH=amd64 go install

bin/timeset: build

format:
	gofmt -w -d -s *.go

.PHONY: build install format
