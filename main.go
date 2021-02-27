// Package main is the main executable of vimprover.
package main

import (
  "fmt"
  "oreflow.com/vimprover/listener"
)

func main() {
  fmt.Printf("Starting vimprover.");
  err := listener.RunListener(func(ke listener.KeyEvent) {
    fmt.Printf("Got Key Event %+v\n", ke)
  })
  if err != nil {
    fmt.Printf("Vimprover terminated with error %+v.", err);
  } else {
    fmt.Printf("Vimprover terminated.");
  }
}
