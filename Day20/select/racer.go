package racer

import (
	"fmt"
	"net/http"
	"time"
)

var (
	TenSecondsTimeout = 10 * time.Second
)

func Racer(urlA, urlB string) (string, error) {
	return ConfigurableRacer(urlA, urlB, TenSecondsTimeout)
}

func ConfigurableRacer(urlA, urlB string, timeout time.Duration) (string, error) {
	select {
	case <-ping(urlA):
		return urlA, nil

	case <-ping(urlB):
		return urlB, nil

	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", urlA, urlB)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
