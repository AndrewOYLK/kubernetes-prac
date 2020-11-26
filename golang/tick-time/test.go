package main

import (
  "fmt"
  "time"
)

func main() {
  tick := time.Tick(2 * time.Second)
  for {
    fmt.Println(<-tick)
  }
}
