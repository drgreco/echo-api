# echo-api

A simple api server that echoes back whatever string a client sent through an API. Written in golang

## Usage

Post `application/x-www-form-urlencoded` data to the endpoint with a key of `echo` and any string for the value, and you will get a response of the value posted.

```
❯ curl localhost:8080/ -d "echo=hello"                                                                                             ─╯
hello
```

### Build

```
❯ go build
```

### Run

```
❯ ./echo-api
2021/01/30 15:58:55 Creating server at localhost:8080...
```


To make requests, you can do some thing like this:
```
❯ curl localhost:8080 -d "echo=hello" 
```

You can also hit the `healthz` endpoint
```
❯ curl localhost:8080/healthz                                                                                                      ─╯
OK
```

### Testing Changes

Run `go test`

```
❯ go test
2021/01/30 17:49:05  - invalid_http_method: GET
2021/01/30 17:49:05  - 415: unsupported_media_type: application/json
2021/01/30 17:49:05  - 400: bad_request
2021/01/30 17:49:05  - 400: bad_request
2021/01/30 17:49:05  - echo: success
PASS
ok  	_/home/greco/src/drgreco/echo-api	0.003s
```

Additional testing can be found [here](test/README.md)

#### Configuration

echo-api accepts three different environment variables
 - `ECHO_HOST` for the ip to listen on. defaults to `localhost`
 - `ECHO_PORT` for the port to listen on. defaults to `8080`
 - `ECHO_HEALTHCHECK` where to listen for a healthcheck. requires leading `/`. defaults to `/healthz`

## Features

  - [x] Communicate with multiple clients simultaneously
  - [x] Create Unit Tests

### Extra

  - [ ] makefile to easily build and demonstrate your server capabilities
  - [ ] Documentation
  - [ ] SSL
  - [ ] authentication
