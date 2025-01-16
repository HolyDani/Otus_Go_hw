package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		url          string
		method       string
		body         []byte
		expectedCode int
		expectedBody string
	}{
		{
			url:          "/",
			method:       http.MethodGet,
			expectedCode: http.StatusOK,
			expectedBody: "This is GET request!",
		},
		{
			url:          "/",
			method:       http.MethodPost,
			body:         []byte("test body"),
			expectedCode: http.StatusOK,
			expectedBody: "Receive POST request with body test body\n",
		},
		{
			url:          "/",
			method:       http.MethodPut,
			body:         []byte("test body"),
			expectedCode: http.StatusMethodNotAllowed,
			expectedBody: "Method not allowed\n",
		},
	}
	for _, test := range tests {
		req := httptest.NewRequest(test.method, test.url, bytes.NewBuffer(test.body))
		rr := httptest.NewRecorder()

		handler(rr, req)

		res := rr.Result()
		if res.StatusCode != test.expectedCode {
			t.Errorf("expected status %d, got %d", test.expectedCode, res.StatusCode)
		}

		body := rr.Body.String()
		if body != test.expectedBody {
			t.Errorf("expected body %s, got %s", test.expectedBody, res.Body)
		}
	}
}
