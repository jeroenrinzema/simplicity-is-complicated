package main

import (
  "fmt"
  "time"
)

func sleep(animal string) {
  counter := 0

  for {
    counter++
    fmt.Println(counter, animal)
    time.Sleep(time.Second)
  }
}

func main() {
  sleep("sheep")
  sleep("chicken")
}