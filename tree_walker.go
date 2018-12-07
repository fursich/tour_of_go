package main

import (
  "fmt"
  "golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
  if t.Left != nil {
    Walk(t.Left, ch)
  }
  if t.Right != nil {
    Walk(t.Right, ch)
  }
  ch <- t.Value
  return
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
  ch1 := make(chan int, 10)
  ch2 := make(chan int, 10)
  Walk(t1, ch1)
  Walk(t2, ch2)

  var found bool

  for i := 0; i < 10; i++ {
    v1 := <-ch1
    found = false
    for j := 0; j < 10; j++ {
      v2 := <- ch2
      if v1 == v2 {
        found = true
        break
      }
      ch2 <- v2
    }
    if found == false {
      return false
    }
  }
  return true
}

func main() {
//  ch := make(chan int)
//  go Walk(tree.New(1), ch)
//  for i := 0; i < 10; i++ {
//    fmt.Println(<-ch)
//  }
  fmt.Printf("Tree(1) == Tree(1): %v\n", Same(tree.New(1), tree.New(1)))
  fmt.Printf("Tree(1) == Tree(2): %v\n", Same(tree.New(1), tree.New(2)))
}
WhiWhi
