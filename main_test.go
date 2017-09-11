package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRootHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	homeHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: wanted %v, but got %v", http.StatusOK, status)
	}
}

func TestSineHandler(t *testing.T) {
	rr := httptest.NewRecorder()

	var tt = []struct {
		name   string
		input  string
		status int
		err    bool
	}{
		{"Input=int", "5", http.StatusOK, false},
		{"Input=float", "10.0", http.StatusOK, false},
		{"Input=string", "hello", http.StatusOK, true},
		{"Input=empty", "", http.StatusOK, true},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			req, err := http.NewRequest("GET", "/sine?input="+tc.input, nil)
			if err != nil {
				t.Fatal(err)
			}

			sineHandler(rr, req)

			if status := rr.Code; status != tc.status {
				t.Errorf("handler returned wrong status: want %v, but got %v", tc.status, status)
			}

			body := rr.Body.String()
			if tc.err && !strings.Contains(body, "<p>Error:") {
				t.Errorf("handler should return error, got %s", body)
			}
			if !tc.err && !strings.Contains(body, "<p>Result:") {
				t.Errorf("handler should return result, got:\n%s", body)
			}
		})
	}
}
