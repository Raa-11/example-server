package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	s := service()
	rr := httptest.NewRecorder()
	s.ServeHTTP(rr, req)

	return rr
}

func TestPing(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/ping", nil)

	resp := executeRequest(req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected response code %d. Got %d\n", resp.Code, http.StatusOK)
	}
}
