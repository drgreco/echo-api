.DEFAULT_GOAL := all
.PHONY: all test-build lint test start run clean-certs

BIN_FILE=bin/echo-api

build:
	go build -o ${BIN_FILE} main.go handlers.go

create-certs:
	mkdir --parents ssl && \
	openssl req -subj '/CN=example.com/O=echo-api/C=US' -new -newkey rsa:2048 -sha256 -days 365 -nodes -x509 -keyout ssl/server.key -out ssl/server.crt

clean-certs:
	test -d ssl && rm --recursive --force ssl

recreate-certs: clean-certs create-certs

run: create-certs
	go run .

start: create-certs build
	./${BIN_FILE}

test:
	go test -cover -tags test 

lint: 
	golint

clean:
	test -d bin && rm --recursive --force bin 

test-build: test build

all: test-build start
