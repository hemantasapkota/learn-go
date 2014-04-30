package main

import (
  "fmt"
  "time"
  "math/rand"
  )

//Returns recieve-only channel
func boring(msg string) <-chan string {
  c := make(chan string)
  go func() {
    for i := 0; ; i++ {
      c <- fmt.Sprintf("%s %d", msg, i)
      time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
    }
  }()
  return c
}

func main() {
  ann := boring("Ann!")
  joe := boring("Joe!")
  for i := 0; i < 5; i++ {
    fmt.Println(<-ann)
    fmt.Println(<-joe)
  }
  fmt.Println("You're boring. I'm leaving")
}
