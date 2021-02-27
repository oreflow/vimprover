// Package main is the main executable of vimprover.
package main

import (
  "fmt"
  "oreflow.com/vimprover/listener"
  "oreflow.com/vimprover/keyboard"
)

func main() {
  fmt.Printf("Starting vimprover.");
  state := keyboard.NewKeyboardState()
  err := listener.RunListener(func(ke listener.KeyEvent) {
    fmt.Printf("Got Key Event %+v\n", ke)
    state = state.AddEvent(ke)
    fmt.Printf("Current State %+v\n", state)
  })
  if err != nil {
    fmt.Printf("Vimprover terminated with error %+v.", err);
  } else {
    fmt.Printf("Vimprover terminated.");
  }
}
