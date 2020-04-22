package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("Return url of first reponder", func(t *testing.T) {
		slowServer := makeServer(20 * time.Millisecond)
		fastServer := makeServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		expected := fastUrl
		got, _ := Racer(slowUrl, fastUrl)

		if expected != got {
			t.Errorf("Expected %s, but want %s", expected, got)
		}
	})

	t.Run("Return error if urls don't response after 10 sec", func(t *testing.T) {
		server := makeServer(20 * time.Millisecond)

		_, err := ConfigurableRacer(server.URL, server.URL, 10*time.Millisecond)

		if err == nil {
			t.Errorf("Expected error but was nil")
		}
	})
}

func makeServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
