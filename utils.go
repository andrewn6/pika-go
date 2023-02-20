package main

import (
  "time"
  "fmt"
)

func nowTimestamp() uint64 { 
  return uint64(time.Now().UnixNano() / 1000000)
}

func main() {
  timestamp := nowTimestamp()
  fmt.Println(timestamp)
}
