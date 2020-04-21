package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	checked := make(map[string]bool)
	channel := make(chan result)

	for _, url := range urls {
		go func(url string){
			channel <- result{url, wc(url)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <- channel
		checked[result.string] = result.bool
	}

	return checked
}