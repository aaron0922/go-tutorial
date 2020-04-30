package main

import (
    "fmt"
	"./append"
	newutil "../internal/append"
)

func main() {
  fmt.Println(append.Add("Hello world"))
  fmt.Println(newutil.Add("Hello world"))
}