package main

import (
  "strings"
  "code.google.com/p/go-tour/wc"
  )

func WordCount(s string) map[string]int {
  m := make(map[string]int)
  for _, key := range strings.Fields(s) {
    v, ok := m[key]
    if ok {
      m[key] = v + 1
    } else {
      m[key] = 1
    }
  }
  return m
}

func main() {
  wc.Test(WordCount)
}
