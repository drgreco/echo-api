package main

import (
    "encoding/base64"
    "net/http"
    "net/http/httptest"
    "net/url"
    "strings"
    "testing"
)

var BasicAuthData = make(map[string]string)

func TestEchoRequestHandler(t *testing.T) {
    // create parameters for testing
    tests := map[string]struct {
        method      string
        contentType string
        dataKey     string
        dataValue   string
        statusCode  int
        err         error
    }{
        "WrongMethod": {"GET", "application/x-www-form-urlencoded", "echo", "success",  http.StatusMethodNotAllowed, nil},
        "WrongContentType": {"POST", "application/json", "echo", "success", http.StatusUnsupportedMediaType, nil},
        "WrongData": {"POST", "application/x-www-form-urlencoded", "dontecho", "success", http.StatusBadRequest, nil},
        "MissingData": {"POST", "application/x-www-form-urlencoded", "echo", "", http.StatusBadRequest, nil},
        "GoodTest": {"POST", "application/x-www-form-urlencoded", "echo", "success", http.StatusOK, nil},
    }

    testAuth := map[string]struct {
        username string
        password string
    }{
        "GoodCredentials": {"user", "test"},
        "BadCredentials": {"notauthorized", "nopass"},
    }

    // create mock BasicAuthData
    BasicAuthData[testAuth["GoodCredentials"].username] = base64.StdEncoding.EncodeToString([]byte(testAuth["GoodCredentials"].password))

    // create loop for tests with good auth
    for name, test := range tests {
        t.Run(name, func(t *testing.T) {
            credentials := testAuth["GoodCredentials"]
            data := url.Values{}
            data.Set(test.dataKey, test.dataValue)

            // set up valid request
            req, err := http.NewRequest(test.method, "/", strings.NewReader(data.Encode()))
            req.Header.Add("Content-Type", test.contentType)
            // use good credentails
            req.SetBasicAuth(credentials.username, credentials.password)
            if err != nil {
                t.Fatal(err)
            }

            rr := httptest.NewRecorder()
            handler := http.HandlerFunc(echoRequestHandler)
            handler.ServeHTTP(rr, req)

            // verify status
            if status := rr.Code; status != test.statusCode {
                t.Errorf("handler returned worng status code: got %v want %v", status, test.statusCode)
            }

            // verify body if status is http.StatusOK
            if status := rr.Code; status == http.StatusOK {
                expected := test.dataValue
                if rr.Body.String() != expected {
                    t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
                }
            }
        })
    }

    // test with bad auth
    t.Run("BadAuth", func(t *testing.T) {
        test := tests["GoodTest"]
        credentials := testAuth["BadCredentials"]
        data := url.Values{}
        data.Set(test.dataKey, test.dataValue)

         // set up a good request
         req, err := http.NewRequest(test.method, "/", strings.NewReader(data.Encode()))
         req.Header.Add("Content-Type", test.contentType)
         // use bad credentails
         req.SetBasicAuth(credentials.username, credentials.password)
         if err != nil {
             t.Fatal(err)
         }

         rr := httptest.NewRecorder()
         handler := http.HandlerFunc(echoRequestHandler)
         handler.ServeHTTP(rr, req)

         // verify status
         if status := rr.Code; status != http.StatusUnauthorized {
             t.Errorf("handler returned worng status code: got %v want %v", status, http.StatusUnauthorized)
         }
    })
}

func TestHealthCheckHandler(t *testing.T) {
    // set up request
    req, err := http.NewRequest("GET", "/healthz", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(healthCheckHandler)

    handler.ServeHTTP(rr, req)

    // verify status
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned worng status code: got %v want %v", status, http.StatusOK)
    }

    // verify body
    expected := `OK`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }
}
