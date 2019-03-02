/*
	Упражнение: поисковый робот
		https://go-tour-ru-ru.appspot.com/concurrency/10
*/

package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl[]string =
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	visitedUrls := struct {
		k   map[string]bool
		mux sync.Mutex
	}{
		k: make(map[string]bool),
	}
	var wg sync.WaitGroup

	var crawler func(string, int)
	crawler = func(url string, depth int) {
		defer wg.Done()

		// This implementation doesn't do either:
		if depth <= 0 {
			return
		}

		visitedUrls.mux.Lock()
		if _, ok := visitedUrls.k[url]; ok {
			visitedUrls.mux.Unlock()
			return
		} else {
			visitedUrls.k[url] = true
			visitedUrls.mux.Unlock()
		}

		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)
		for _, u := range urls {
			wg.Add(1)
			go crawler(u, depth-1)
		}
	}
	wg.Add(1)
	crawler(url, depth)
	// Ждём пока закончит свою работу вся группа горутин
	wg.Wait()
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
