package main

import (
  "fmt"
  "sync"
  "time"
)

type Counter struct{
  v map[string]int
  mux sync.Mutex
}

func (c Counter) Inc(key string) {
  c.mux.Lock()
  c.v[key] += 1
  c.mux.Unlock()
}

func (c Counter) Value(key string) int {
  defer c.mux.Unlock()

  c.mux.Lock()
  return c.v[key]
}

func main() {
  c := Counter{v: make(map[string]int)}
  for i:=0;i<100000;i++ {
    go c.Inc("hoge")
  }
  time.Sleep(1 * time.Millisecond)
  fmt.Println(c.Value("hoge"))
}
