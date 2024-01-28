package multithreaded

import (
	"fmt"
	"sync"
	"webcrawler/utils"
)

var (
	urlMap = map[string]bool{}
	lock   = sync.RWMutex{}
)

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func crawl(wg *sync.WaitGroup, url string, depth int, fetcher utils.Fetcher) {
	defer wg.Done()
	if depth <= 0 {
		return
	}

	// Writer's lock as we are updating urlMap
	lock.Lock()
	urlMap[url] = true
	lock.Unlock()

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {

		// Reader's lock as we are simply reading urlMap
		lock.RLock()
		exists := urlMap[u]
		lock.RUnlock()

		if !exists {
			wg.Add(1)
			go crawl(wg, u, depth-1, fetcher)
		}
	}
}

func Run() {
	var wg sync.WaitGroup
	wg.Add(1)
	go crawl(&wg, "https://golang.org/", 4, utils.FakeFetcher)
	wg.Wait()
}
