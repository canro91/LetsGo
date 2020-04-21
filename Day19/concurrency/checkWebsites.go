package concurrency

// import (
// 	"time"
// )

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	checked := make(map[string]bool)

	for _, url := range urls {
		go func(){
			checked[url] = wc(url)
		}()
	}

	// time.Sleep(2*time.Second)

	return checked
}