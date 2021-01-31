.DEFAULT_GOAL := all

BIN_FILE=bin/echo-api

build:
	go build -o ${BIN_FILE} main.go handlers.go

run: 
	go run .

start: build
	./${BIN_FILE}

test:
	go test -cover -tags test 

lint: 
	golint

clean:
	test -f ${BIN_FILE} && rm ${BIN_FILE}

test-build: test build

all: test-build start
