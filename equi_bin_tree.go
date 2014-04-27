package main

import (
  "code.google.com/p/go-tour/tree"
  "fmt"
  )

func walkImpl(t *tree.Tree, ch chan int) {
  if t.Left != nil {
    walkImpl(t.Left, ch)
  }
  ch <- t.Value
  if t.Right != nil {
    walkImpl(t.Right, ch)
  }
}

func Walk(t *tree.Tree, ch chan int) {
  walkImpl(t, ch)
  close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
  t1ch, t2ch := make(chan int), make(chan int)

  go Walk(t1, t1ch)
  go Walk(t2, t2ch)

  for {
    v1, ok1 := <-t1ch
    v2, ok2 := <-t2ch

    if v1 != v2 || ok1 != ok2 {
      return false
    }

    if !ok1 {
      break
    }
  }

  return true
}

func main() {
  v := Same(tree.New(1), tree.New(1))
  fmt.Println(v)
}
