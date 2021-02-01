//+build !test

package main

import (
    "log"
    "net/http"
    "os"
)

// initialize some default configuration for net/http
func config() (map[string]string) {
    config := make(map[string]string)
    config["Host"]              = "localhost"
    config["Port"]              = "8443"
    config["HealthCheck"]       = "/healthz"
    config["ServerPrivateKey"]  = "ssl/server.key"
    config["ServerCertificate"] = "ssl/server.crt"
    config["TlsDisable"]        = "false"

    // get env variables if set
    Host, ok := os.LookupEnv("ECHO_HOST")
    if ok { config["Host"] = Host }

    Port, ok := os.LookupEnv("ECHO_PORT")
    if ok { config["Port"] = Port }

    HealthCheck, ok := os.LookupEnv("ECHO_HEALTHCHECK")
    if ok { config["HealthCheck"] = HealthCheck }

    ServerPrivateKey, ok := os.LookupEnv("ECHO_SERVERPRIVATEKEY")
    if ok { config["ServerPrivateKey"] = ServerPrivateKey }

    ServerCertificate, ok := os.LookupEnv("ECHO_SERVERCERTIFICATE")
    if ok { config["ServerCertificate"] = ServerCertificate }

    TlsDisable, ok := os.LookupEnv("ECHO_TLS_DISABLE")
    if ok { config["TlsDisable"] = TlsDisable}

    return config
}

func main () {
    config := config()

    // create handler function calls
    http.HandleFunc("/", echoRequestHandler)
    http.HandleFunc(config["HealthCheck"], healthCheckHandler)

    // launch http listener
    // check if we want TLS or not
    if config["TlsDisable"] == "true" {
        log.Printf("Creating http server at %s:%s...", config["Host"], config["Port"])
        log.Fatal(http.ListenAndServe(config["Host"]+":"+config["Port"], nil))
    } else {
        log.Printf("Creating TLS server at %s:%s...", config["Host"], config["Port"])
        log.Fatal(http.ListenAndServeTLS(config["Host"]+":"+config["Port"], config["ServerCertificate"], config["ServerPrivateKey"], nil))
    }
}
