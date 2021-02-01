//+build !test

package main

import (
    "bufio"
    "log"
    "net/http"
    "os"
    "strings"
)

var BasicAuthData = make(map[string]string)

// initialize some default configuration for net/http
func config() map[string]string {
    config := make(map[string]string)
    config["Host"]              = "localhost"
    config["Port"]              = "8443"
    config["HealthCheck"]       = "/healthz"
    config["ServerPrivateKey"]  = "ssl/server.key"
    config["ServerCertificate"] = "ssl/server.crt"
    config["TLSDisable"]        = "false"
    config["BasicAuthFile"]     = "auth.db"

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

    TLSDisable, ok := os.LookupEnv("ECHO_TLS_DISABLE")
    if ok { config["TLSDisable"] = TLSDisable }

    BasicAuthFile, ok := os.LookupEnv("ECHO_BASIC_AUTH_FILE")
    if ok { config["BasicAuthFile"] = BasicAuthFile }

    // Set BasicAuthData
    BasicAuthData = readBasicAuthData(config["BasicAuthFile"])

    return config
}

// Read basic auth file, return map
func readBasicAuthData(BasicAuthFile string) map[string]string {
    // read basic auth file
    // else error out if the basic auth file is not readable
    log.Printf("reading basic auth file: %s", BasicAuthFile)
    file, err := os.Open(BasicAuthFile)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    data := make(map[string]string)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lineSplits := strings.Split(scanner.Text(), ":")
        data[lineSplits[0]] = strings.Join(lineSplits[1:], "")
        log.Printf("adding basic auth for user: %v", lineSplits[0])
    }

    return data
}

func main () {
    config := config()

    // create handler function calls
    http.HandleFunc("/", echoRequestHandler)
    http.HandleFunc(config["HealthCheck"], healthCheckHandler)

    // launch http listener
    // check if we want TLS or not
    if config["TLSDisable"] == "true" {
        log.Printf("Creating http server at %s:%s...", config["Host"], config["Port"])
        log.Fatal(http.ListenAndServe(config["Host"]+":"+config["Port"], nil))
    } else {
        log.Printf("Creating TLS server at %s:%s...", config["Host"], config["Port"])
        log.Fatal(http.ListenAndServeTLS(config["Host"]+":"+config["Port"], config["ServerCertificate"], config["ServerPrivateKey"], nil))
    }
}
