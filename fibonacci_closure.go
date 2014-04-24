package main

import "fmt"

func fibonacci() func(int) int {
  fiba := make([]int, 10)
  return func(i int) int {
    if i == 0 {
      fiba[0] = 0
    } else
    if i == 1 {
      fiba[1] = 1
    } else {
      fiba[i] = fiba[i - 1] + fiba[i - 2]
    }
    return fiba[i]
  }
}

func main() {
  f := fibonacci()
  for i := 0; i < 10; i++ {
    fmt.Println(f(i))
  }
}
