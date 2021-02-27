// Package main is the main executable of vimprover.
package main

import (
  "fmt"
  "time"
  "os/exec"

  "oreflow.com/vimprover/listener"
  "oreflow.com/vimprover/keyboard"
  "oreflow.com/vimprover/vimprovements"
)

// Number of times each Violation Type has been encountered.
var violationCounters = map[vimprovements.ViolationType]uint32{}

func main() {
  fmt.Printf("Starting vimprover.");
  state := keyboard.NewKeyboardState()
  err := listener.RunListener(func(ke listener.KeyEvent) {
    state = state.AddEvent(ke)
    if ke.EventType != listener.KeyDown {
      return
    }
    for _, v := range vimprovements.EnabledVimprovements {
      violation, err := v(state)
      if err != nil {
        fmt.Printf("Vimprover encountered error %+v.", err);
        continue
      }
      if violation != nil {
        fmt.Printf("Vimprovement: %+v\n", violation.Message);
        // Plays an annoying sound when violation happened.
        // https://en.wikipedia.org/wiki/Wilhelm_scream
        go exec.Command("aplay", "assets/wilhelm_scream.wav").Run()
        violationCounters[violation.Type]++
      }
    }
    throttledPrintViolationStats()
  })
  if err != nil {
    fmt.Printf("Vimprover terminated with error %+v.", err);
  } else {
    fmt.Printf("Vimprover terminated.");
  }
}

// minDurationBetweenPrints defines the duration between each time violations stats is printed.
var minDurationBetweenPrints = MustParseDuration("2m")

// lastPrint is the time of the last printed violation stats.
var lastPrint = time.Unix(0, 0)

// Prints the violation stats if not printed in the last minDurationBetweenPrints.
func throttledPrintViolationStats() {
  if time.Since(lastPrint) < minDurationBetweenPrints {
    return
  }
  lastPrint = time.Now()
  fmt.Println()
  fmt.Printf("Violations until %s\n", lastPrint.String())

  for key, value := range violationCounters {
    fmt.Printf("%s: %d\n", key, value)
  }
  fmt.Println()
}

// Helper to create time.Duration, that panics on error.
//
// Only to be used in tests or global instantiations.
func MustParseDuration(s string) time.Duration {
  d, err := time.ParseDuration(s)
  if err != nil {
    panic(err)
  }
  return d
}
