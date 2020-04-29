package context

import (
	"time"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubStore struct {
	response string
	cancelled bool
	t *testing.T
}

func (s *StubStore) Fetch() string{
	time.Sleep(100 * time.Millisecond)

	return s.response
}

func (s *StubStore) Cancel() {
	s.cancelled = true
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

		store.assertWasNotCancelled()
	})

	t.Run("Cancel request", func(t *testing.T){
		data := "Hello, World!"
		store := &StubStore{response: data, t: t}
		server := Server(store)
	
		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5 * time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()
	
		server.ServeHTTP(response, request)
	
		store.assertWasCancelled()
	})
}

func (s *StubStore) assertWasCancelled() {
	s.t.Helper()
	if !s.cancelled {
		s.t.Errorf("store was not told to cancel")
	}
}

func (s *StubStore) assertWasNotCancelled() {
	s.t.Helper()
	if s.cancelled {
		s.t.Errorf("store was told to cancel")
	}
}