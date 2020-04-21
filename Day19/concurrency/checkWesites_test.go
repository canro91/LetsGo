package concurrency

import (
	"time"
	"reflect"
	"testing"
)

func fakeWebsiteChecker(url string) bool {
	if url == "https://any-test-website.com" {
		return false
	}
	return true
}

func TestWebsites(t *testing.T){
	websites := []string {
		"https://www.reddit.com/",
		"https://www.wikipedia.org/",
		"https://any-test-website.com",
	}

	expected := map[string]bool {
		"https://www.reddit.com/" : true,
		"https://www.wikipedia.org/": true,
		"https://any-test-website.com": false,
	}

	checked := CheckWebsites(fakeWebsiteChecker, websites)
	if !reflect.DeepEqual(expected, checked) {
		t.Fatalf("Expected %v, but was %v", expected, checked)
	}
}

func slowWebsiteChecker(url string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B){
	urls := make([]string, 100)
	for i := 0; i < 100; i++ {
		urls[i] = "any url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowWebsiteChecker, urls)
	}
}