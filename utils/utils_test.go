package utils_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danilotorchio/goexpert/multithreading/utils"
)

func TestRequestData_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test data"))
	}))
	defer server.Close()

	data, err := utils.RequestData(server.URL)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expected := []byte("test data")
	if string(data) != string(expected) {
		t.Errorf("unexpected response data, got: %s, want: %s", string(data), string(expected))
	}
}

func TestRequestData_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("server error"))
	}))
	defer server.Close()

	_, err := utils.RequestData(server.URL)
	if err == nil {
		t.Errorf("expected an error, got nil")
	}

	expected := "error: server error"
	if err.Error() != expected {
		t.Errorf("unexpected error message, got: %s, want: %s", err.Error(), expected)
	}
}
