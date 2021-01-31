package main

import (
    "net/http"
    "net/http/httptest"
    "net/url"
    "strings"
    "testing"
)

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

    // create loop for tests
    for name, test := range tests {
        t.Run(name, func(t *testing.T) {
            data := url.Values{}
            data.Set(test.dataKey, test.dataValue)

            // set up valid request
            req, err := http.NewRequest(test.method, "/", strings.NewReader(data.Encode()))
            req.Header.Add("Content-Type", test.contentType)
            if err != nil {
                t.Fatal(err)
            }

            rr := httptest.NewRecorder()
            handler := http.HandlerFunc(echoRequestHandler)

            handler.ServeHTTP(rr, req)

            // verify status
            if status := rr.Code; status != test.statusCode {
                t.Errorf("handler returned worng status gode: got %v want %v", status, test.statusCode)
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
        t.Errorf("handler returned worng status gode: got %v want %v", status, http.StatusOK)
    }

    // verify body
    expected := `OK`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }
}
