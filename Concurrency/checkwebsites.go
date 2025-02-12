package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultsChannel := make(chan result)
	for _, url := range urls {
		go func(url string) {
			resultsChannel <- result{url, wc(url)}
		}(url)
	}
	for i := 0; i < len(urls); i++ {
		r := <-resultsChannel
		results[r.string] = r.bool
	}
	return results
}
