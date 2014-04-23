package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
  p := make([][]uint8, dy)
  for i := 0; i < len(p); i++ {
    p[i] = make([]uint8, dx)
    for j := 0; j < len(p[i]); j++ {
      p[i][j] = uint8(1)
    }
  }
  return p
}

func main() {
  pic.Show(Pic)
}
