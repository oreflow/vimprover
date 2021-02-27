package vimprovements

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"oreflow.com/vimprover/keyboard"
)

var lsViolation = Violation{
  Type: LeftShiftViolation,
  message: "Left shift was used with button on left side of the keyboard",
}

// leftShift vimprovement detects if the left shift key is used in combination
// with keys on the left half of the keyboard.
func leftShift(s *keyboard.State) (*Violation, error) {
  if len(s.LastCharacters) < 1 {
    return nil, status.Error(
      codes.FailedPrecondition,
      "leftshift vimprovement check ran before character input was received.",
    )
  }
  if s.LeftShiftDown && leftShiftViolatingKeys[s.LastCharacters[0]] {
    return &lsViolation, nil
  }
  return nil, nil
}


// leftShiftViolatingKeys defines a set of keys that should not be used together with the left shift key.
var leftShiftViolatingKeys = map[uint16]bool{
  keyboard.KEY_GRAVE: true,
  keyboard.KEY_1: true,
  keyboard.KEY_2: true,
  keyboard.KEY_3: true,
  keyboard.KEY_4: true,
  keyboard.KEY_5: true,
  keyboard.KEY_Q: true,
  keyboard.KEY_W: true,
  keyboard.KEY_E: true,
  keyboard.KEY_R: true,
  keyboard.KEY_T: true,
  keyboard.KEY_A: true,
  keyboard.KEY_S: true,
  keyboard.KEY_D: true,
  keyboard.KEY_F: true,
  keyboard.KEY_G: true,
  keyboard.KEY_102ND: true,
  keyboard.KEY_Z: true,
  keyboard.KEY_X: true,
  keyboard.KEY_C: true,
  keyboard.KEY_V: true,
}
