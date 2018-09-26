package main

import (
	"fmt"
	"sync"
)

// Fetcher ...
type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
	WriteToCache(url string) error
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) { // NOTE 3: use `Fetcher` instead of `*Fetcher`
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}

	if err := fetcher.WriteToCache(url); err != nil {
		fmt.Println(err)
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	Crawl("https://golang.org/", 4, &fetcher) // NOTE 4: use `&fetcher` instead of `fetcher``
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher struct {
	data  map[string]*fakeResult
	cache *cacheResult
}

type fakeResult struct {
	body string
	urls []string
}

type cacheResult struct {
	data map[string]bool
	lock sync.Mutex
}

// NOTE 1: use `fakeFetcher` instead of `*fakeFetcher`
func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f.data[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// NOTE 2: use `*fakeFetcher` instead of `fakeFetcher`
func (f *fakeFetcher) WriteToCache(url string) error {
	f.cache.lock.Lock()
	if _, ok := f.cache.data[url]; ok {
		e := fmt.Errorf("cached %s", url)
		f.cache.lock.Unlock()
		return e
	}

	f.cache.data[url] = true
	f.cache.lock.Unlock()
	return nil
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	cache: &cacheResult{make(map[string]bool), sync.Mutex{}},
	data: map[string]*fakeResult{
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
		}},
}
