package main

import (
  "fmt"
  "sync"
  "time"
)

type Fetcher interface {
  // Fetch returns the body of URL and
  // a slice of URLs found on that page.
  Fetch(url string) (body string, urls []string, err error)
}

type FetchStatus struct {
  fetched map[string]bool
  mux sync.Mutex
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, fs *FetchStatus) {
  // TODO: Fetch URLs in parallel.
  // TODO: Don't fetch the same URL twice.
  // This implementation doesn't do either:
  if depth <= 0 {
    return
  }

  fs.mux.Lock()
  defer fs.mux.Unlock()

  if _, ok := fs.fetched[url]; ok {
    return
  }

  body, urls, err := fetcher.Fetch(url)
  fs.fetched[url] = true

  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Printf("found: %s %q\n", url, body)
  for _, u := range urls {
    go Crawl(u, depth-1, fetcher, fs)
  }
  return
}

func main() {
  fs := FetchStatus{fetched: make(map[string]bool)}
  Crawl("https://golang.org/", 4, fetcher, &fs)
  time.Sleep(1000 * time.Millisecond)
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
