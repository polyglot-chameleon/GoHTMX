package main

import (
	"math"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

var test_rw *httptest.ResponseRecorder

func init() {
	addRoutes()
	test_rw = httptest.NewRecorder()
}

func testRequest(t *testing.T, method string, path string, contentLength int) {
	req, err := http.NewRequest(method, path, nil)

	if err != nil {
		t.Fatal(err)
	}

	http.DefaultServeMux.ServeHTTP(test_rw, req)

	if test_rw.Code != http.StatusOK {
		t.Fatalf("%v != %v", test_rw.Code, http.StatusOK)
	}

	header := test_rw.Header()

	isContentLength, _ := strconv.Atoi(header.Get("Content-Length"))

	if math.Abs(float64(isContentLength-contentLength)) > 10 {
		t.Fatalf("%v != %v", isContentLength, contentLength)
	}

}

func TestGetRoot(t *testing.T) {
	testRequest(t, http.MethodGet, "/", 510)
	testRequest(t, http.MethodGet, "/htmx.min.js", 49567)
}

func TestGetCreateForm(t *testing.T) {
	testRequest(t, http.MethodGet, "/form/create", 304)
}

func TestGetData(t *testing.T) {
	testRequest(t, http.MethodGet, "/data", 304)
}

func TestPOSTData(t *testing.T) {
	testRequest(t, http.MethodPost, "/data", 304)
}
