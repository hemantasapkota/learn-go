package main

import (
  "fmt"
  "runtime"
  )

func main() {
  fmt.Print("Go rus on ")
  switch os := runtime.GOOS; os {
    case "darwin":
      fmt.Println("OX X")
    case "linux":
      fmt.Println("Linux")
    default:
      fmt.Printf("%s.", os)
  }
}
