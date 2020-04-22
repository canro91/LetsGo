package racer

import (
	"net/http"
	"time"
)

func Racer(urlA, urlB string) string {
	durationA := duration(urlA)
	durationB := duration(urlB)

	if durationA < durationB {
		return urlA
	}
	return urlB
}

func duration(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
