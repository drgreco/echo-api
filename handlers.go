package main

import (
    "encoding/base64"
    "fmt"
    "log"
    "net/http"
)

// Handle POST echo
func echoRequestHandler(w http.ResponseWriter, r *http.Request) {
    // First thing - check basic auth
    username, password, ok := r.BasicAuth()
    if ! ok || BasicAuthData[username] != base64.StdEncoding.EncodeToString([]byte(password)) {
        w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
        w.WriteHeader(http.StatusUnauthorized)
        w.Write([]byte("status_unauthorized"))
        log.Printf("%v - %v: status_unauthorized", r.RemoteAddr, http.StatusUnauthorized)
        return
    }

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
        log.Printf("%v - %v - echo: %v", r.RemoteAddr, username, echo)
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
