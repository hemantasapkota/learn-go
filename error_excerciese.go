package main

import (
  "fmt"
  "math"
  )

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
  return fmt.Sprintf("Cannot square root negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {

  if x < 0 {
    return 0, ErrNegativeSqrt(x)
  }

  z := float64(1)
  for i := 0; i < 10; i++ {
    z = z - (math.Pow(z, 2) - x) / (2 * z)
  }
  return z, nil
}

func main() {
  fmt.Println(Sqrt(2))
  fmt.Println(Sqrt(-2))
}
