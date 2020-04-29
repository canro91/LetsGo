package context

import (
	"errors"
	"time"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubStore struct {
	response string
	t *testing.T
}

func (s *StubStore) Fetch(ctx context.Context) (string, error){
	data := make(chan string, 1)

	go func(){
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("spy cancelled")
				return

		default:
			time.Sleep(10 * time.Millisecond)
				result += string(c)

			}

		}
		data <- result
	}()

	select {
	case d := <-data:
		return d, nil

	case <-ctx.Done():
		return "", ctx.Err()
	}

}

func TestHandler(t *testing.T){
	t.Run("Return body", func(t *testing.T){
		data := "Hello, World!"
		store := &StubStore{response: data, t: t}
		server := Server(store)
	
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
	
		server.ServeHTTP(response, request)
	
		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}
	})

	t.Run("Cancel request", func(t *testing.T){
		data := "Hello, World!"
		store := &StubStore{response: data, t: t}
		server := Server(store)
	
		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5 * time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{}
	
		server.ServeHTTP(response, request)
	
		if response.written {
			t.Error("a response should not have been written")
		}
	})
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}