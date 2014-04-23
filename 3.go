package main

import (
  "fmt"
  "net"
  "os"
  "time"
  )

func main() {
  fmt.Println("Welceomt to the Playground")

  fmt.Println("The time is", time.Now())

  fmt.Println("And if you try to open a file:")
  fmt.Println(os.Open("filename"))

  fmt.Println("Or access the network:")
  fmt.Println(net.Dial("tcp", "google.com:80"))
}
