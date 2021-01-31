package main

import (
    "log"
    "net/http"
    "os"
)

const (
    HostDefault = "localhost"
    PortDefault = "8080"
    HealthCheckDefault = "/healthz"
)

func main () {
    // get env variables if set
    Host, ok := os.LookupEnv("ECHO_HOST")
    if !ok { Host = HostDefault }

    Port, ok := os.LookupEnv("ECHO_PORT")
    if !ok { Port = PortDefault }

    HealthCheck, ok := os.LookupEnv("ECHO_HEALTHCHECK")
    if !ok { HealthCheck = HealthCheckDefault }

    // create handler function calls
    http.HandleFunc("/", echoRequestHandler)
    http.HandleFunc(HealthCheck, healthCheckHandler)

    // launch http listener
    log.Printf("Creating server at %s:%s...", Host, Port)
    log.Fatal(http.ListenAndServe(Host+":"+Port, nil))
}
