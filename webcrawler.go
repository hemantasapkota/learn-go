package main

import (
  "fmt"
  )

type Fetcher interface {
  Fetch(url string, ch chan fakeResult)
}

func Crawl(url string, depth int, fetcher Fetcher, ch chan fakeResult) {

  if depth <= 0 {
    return
  }

  go fetcher.Fetch(url, ch)

  v := <-ch
  fmt.Println(v.body)

  for _,u := range v.urls {
    Crawl(u, depth - 1, fetcher, ch)
  }

}

func main() {
  ch := make(chan fakeResult)
  Crawl("http://golang.org/", 4, fetcher, ch)
  close(ch)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
    body string
    urls []string
}

func (f fakeFetcher) Fetch(url string, ch chan fakeResult) {
  if res, ok := f[url]; ok {
    ch <- *res
  }
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
