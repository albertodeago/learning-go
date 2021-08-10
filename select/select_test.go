package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func createServer(t time.Duration) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(t)
		w.WriteHeader(http.StatusOK)
	}))

	return server
}

func TestRacer(t *testing.T) {

	// t.Run("returns the fastest server", func(t *testing.T) {
	// 	slowServer := createServer(20 * time.Millisecond)
	// 	fastServer := createServer(0 * time.Millisecond)

	// 	defer slowServer.Close()
	// 	defer fastServer.Close()

	// 	slowURL := slowServer.URL
	// 	fastURL := fastServer.URL

	// 	want := fastURL
	// 	got, err := Racer(slowURL, fastURL)

	// 	if err != nil {
	// 		t.Fatalf("did not expect an error but got one %v", err)
	// 	}

	// 	if got != want {
	// 		t.Errorf("got %q, want %q", got, want)
	// 	}
	// })

	t.Run("returns an error if a server doesn't respond within the timeout", func(t *testing.T) {
		server := createServer(25 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}
