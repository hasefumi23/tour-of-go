package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type SafeMapper struct {
	v   map[string]int
	mux sync.Mutex
}

func (m *SafeMapper) Inc(key string) {
	m.mux.Lock()
	m.v[key]++
	m.mux.Unlock()
}

func (m *SafeMapper) Value(key string) (int, bool) {
	m.mux.Lock()
	defer m.mux.Unlock()
	v, ok := m.v[key]
	return v, ok
}

func Crawl(url string, depth int, fetcher Fetcher, m SafeMapper) {
	var wg sync.WaitGroup
	var crawl func(string, int)
	crawl = func(url string, depth int) {
		if depth <= 0 {
			return
		}

		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)

		wg.Add(len(urls))
		for _, u := range urls {
			go func(u string) {
				_, ok := m.Value(u)
				// avoid dupulication
				if !ok {
					m.Inc(u)
					crawl(u, depth-1)
				}
				wg.Done()
			}(u)
		}
	}
	crawl(url, depth)
	wg.Wait()
}

func main() {
	m := SafeMapper{v: make(map[string]int)}
	url := "https://golang.org/"
	m.Inc(url)
	Crawl(url, 10, fetcher, m)
}

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
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
