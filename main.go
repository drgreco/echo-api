//+build !test

package main

import (
    "log"
    "net/http"
    "os"
)

// initialize some configuration for net/http
func config() (string, string, string) {
    HostDefault        := "localhost"
    PortDefault        := "8080"
    HealthCheckDefault := "/healthz"

    // get env variables if set
    Host, ok := os.LookupEnv("ECHO_HOST")
    if !ok { Host = HostDefault }

    Port, ok := os.LookupEnv("ECHO_PORT")
    if !ok { Port = PortDefault }

    HealthCheck, ok := os.LookupEnv("ECHO_HEALTHCHECK")
    if !ok { HealthCheck = HealthCheckDefault }

    return Host, Port, HealthCheck
}

func main () {
    Host, Port, HealthCheck := config()

    // create handler function calls
    http.HandleFunc("/", echoRequestHandler)
    http.HandleFunc(HealthCheck, healthCheckHandler)

    // launch http listener
    log.Printf("Creating server at %s:%s...", Host, Port)
    log.Fatal(http.ListenAndServe(Host+":"+Port, nil))
}
