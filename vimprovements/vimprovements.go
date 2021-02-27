// Package vimprovements implements the actual Vimprover checks.
package vimprovements

import (
  "oreflow.com/vimprover/keyboard"
)

// ViolationType provides an unique identifier for each Vimprovement type.
type ViolationType string

const (
    LeftShiftViolation ViolationType = "LeftShiftViolation "
    RightShiftViolation ViolationType = "RightShiftViolation "
    RepeatedKeypressViolation ViolationType = "RepeatedKeypressViolation "
)

// Violation is used to represent an area to vimprove in.
type Violation struct {
  Type ViolationType
  Message string
}

// Vimprovement defines the interface of a vimprovement implementation.
type Vimprovement func (*keyboard.State) (*Violation, error)

// EnabledVimprovements defines the set of vimprovements to monitor for
var EnabledVimprovements = []Vimprovement{
  leftShift,
  rightShift,
  repeatedKeypress,
}
