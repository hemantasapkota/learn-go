package main

import "fmt"

func main() {
  c := make(chan int, 100)
  for i := 0; i < 1000; i++ {
    c <- i
    fmt.Println(<-c)
    fmt.Println(<-c)
  }
}
