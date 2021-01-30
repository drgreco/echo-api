package main

import (
    "fmt"
    "os"
    "log"
    "net/http"
)

const (
    HostDefault = "localhost"
    PortDefault = "8080"
    HealthCheckDefault = "/healthz"
)

// Handle POST echo
func echoRequest(w http.ResponseWriter, r *http.Request) {
    // if this isn't a POST, return error and disregard
    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        fmt.Fprintf(w, "invalid_http_method: only POST accepted")
        log.Printf("%v - invalid_http_method: %v", r.RemoteAddr, r.Method)
        return
    }

    // if this is not the right content-type, throw an error
    // content-type: application/x-www-form-urlencoded
    if r.Header.Get("Content-type") != "application/x-www-form-urlencoded" {
        w.WriteHeader(http.StatusUnsupportedMediaType)
        fmt.Fprintf(w, "unsupported_media_type: only application/x-www-form-urlencoded accepted")
        log.Printf("%v - %v: unsupported_media_type: %v", r.RemoteAddr, http.StatusUnsupportedMediaType, r.Header.Get("Content-type"))
        return
    }

    // parse form data
    r.ParseForm()

    // return value of echo if it is not empty
    // else return error
    echo := r.Form.Get("echo")
    if len(echo) > 0 {
        log.Printf("%v - echo: %v", r.RemoteAddr, echo)
        fmt.Fprintf(w, echo)
    } else {
        log.Printf("%v - %v: bad_request", r.RemoteAddr, http.StatusBadRequest)
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("echo not found or empty"))
    }
}

// healthCheck should just return 200: OK, no logging
func healthCheck(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "OK")
}


func main () {
    // get env variables if set
    Host, ok := os.LookupEnv("ECHO_HOST")
    if !ok { Host = HostDefault }

    Port, ok := os.LookupEnv("ECHO_PORT")
    if !ok { Port = PortDefault }

    HealthCheck, ok := os.LookupEnv("ECHO_HEALTHCHECK")
    if !ok { HealthCheck = HealthCheckDefault }

    // create handle function
    http.HandleFunc("/", echoRequest)
    http.HandleFunc(HealthCheck, healthCheck)

    // launch http listener
    log.Printf("Creating server at %s:%s...", Host, Port)
    log.Fatal(http.ListenAndServe(Host+":"+Port, nil))
}
