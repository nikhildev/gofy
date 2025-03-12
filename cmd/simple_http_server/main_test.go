package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMains(t *testing.T) {
	t.Log("TestMain")
}

func TestHealthEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := http.NewServeMux()
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "OK"
	actual, err := io.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(actual) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", string(actual), expected)
	}
}

func TestEchoEndpoint(t *testing.T) {
	message := "hello"
	req, err := http.NewRequest("GET", "/echo/"+message, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := http.NewServeMux()
	router.HandleFunc("/echo/", func(w http.ResponseWriter, r *http.Request) {
		message := r.URL.Path[len("/echo/"):]
		w.Write([]byte("Echo: " + message))
	})

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Echo: " + message
	actual, err := io.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(actual) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", string(actual), expected)
	}
}
