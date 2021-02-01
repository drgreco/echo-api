# echo-api

A simple api server that echoes back whatever string a client sent through an API. Written in golang

## Usage

Post `application/x-www-form-urlencoded` data to the endpoint with a key of `echo` and any string for the value, and you will get a response of the value posted.

```
❯ curl https://localhost:8443/ --insecure --user "user:test" --data "echo=hello" 
hello
```

You can also hit the `healthz` endpoint
```
❯ curl https://localhost:8443/healthz --insecure
OK
```

Note that the `healthz` endpoint does not require basic auth.

### Build

```
❯ go build
```

### Run

```
❯ ./echo-api
2021/01/30 15:58:55 Creating server at localhost:8443...
```

## Testing Changes

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

Additional testing can be found [here](testing/README.md)

### Makefile targets

`Makefile` targets:
| argument       | description |
|----------------|----------------|
| build          | creates executable at `bin/echo-api`         |
| create-certs   | creates certificates for testing in `ssl/`   |
| clean-certs    | deletes directory `ssl/`                     |
| recreate-certs | runs `clean-certs` then `create-certs`       |
| run            | runs without creating executable: `go run .` |
| start          | runs `build`, then executes `bin/echo-api`   |
| test           | runs tests: `go test -cover -tags test `     |
| lint           | runs golint                                  |
| clean          | deletes directory `bin/`                     |
| test-build     | runs the arguments `test` then `build`       |
| all            | runs the arguments `test-build` then `start  |


Defaults to `all`

#### Configuration

echo-api accepts three different environment variables
 - `ECHO_HOST` for the ip to listen on. defaults to `localhost`
 - `ECHO_PORT` for the port to listen on. defaults to `8443`
 - `ECHO_HEALTHCHECK` where to listen for a healthcheck. requires leading `/`. defaults to `/healthz`
 - `ECHO_SERVERPRIVATEKEY` path to tls private key. defaults to `ssl/server.key`
 - `ECHO_SERVERCERTIFICATE` path to tls certificate. defaults to `ssl/server.crt`
 - `ECHO_TLS_DISABLE` disables TLS. Only accepts strings `true` or `false`. defaults to `false`.
 - `ECHO_BASIC_AUTH_FILE` file to find basic auth information. defaults to `auth.db`.

#### Basic Auth

Basic Auth is handled by default in a file named `auth.db`. This file looks like this:
```
user:dGVzdA==
bob:bGV0bWVpbg==
```

The format is `[user]:[base64encoded password]

## Features

  - [x] Communicate with multiple clients simultaneously
  - [x] Create Unit Tests

### Extra

  - [x] makefile to easily build and demonstrate your server capabilities
  - [x] Documentation
  - [x] SSL
  - [x] authentication
