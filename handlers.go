package main

import (
    "fmt"
    "log"
    "net/http"
)

// Handle POST echo
func echoRequestHandler(w http.ResponseWriter, r *http.Request) {
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

// HealthCheck should just return 200: OK, no logging
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "OK")
}
