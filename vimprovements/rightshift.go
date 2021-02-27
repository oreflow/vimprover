package vimprovements

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"oreflow.com/vimprover/keyboard"
)

var rsViolation = Violation{
  Type: RightShiftViolation,
  Message: "Right shift was used with button on right side of the keyboard",
}

// rightShift vimprovement detects if the right shift key is used in combination
// with keys on the right half of the keyboard.
func rightShift(s *keyboard.State) (*Violation, error) {
  if len(s.LastCharacters) < 1 {
    return nil, status.Error(
      codes.FailedPrecondition,
      "rightshift vimprovement check ran before character input was received.",
    )
  }
  if s.RightShiftDown && rightShiftViolatingKeys[s.LastCharacters[0]] {
    return &rsViolation, nil
  }
  return nil, nil
}



// rightShiftViolatingKeys defines a set of keys that should not be used together with the right shift key.
var rightShiftViolatingKeys = map[uint16]bool{
  keyboard.KEY_7: true,
  keyboard.KEY_8: true,
  keyboard.KEY_9: true,
  keyboard.KEY_0: true,
  keyboard.KEY_MINUS: true,
  keyboard.KEY_EQUAL: true,
  keyboard.KEY_U: true,
  keyboard.KEY_I: true,
  keyboard.KEY_O: true,
  keyboard.KEY_P: true,
  keyboard.KEY_LEFTBRACE: true,
  keyboard.KEY_RIGHTBRACE: true,
  keyboard.KEY_H: true,
  keyboard.KEY_J: true,
  keyboard.KEY_K: true,
  keyboard.KEY_L: true,
  keyboard.KEY_SEMICOLON: true,
  keyboard.KEY_APOSTROPHE: true,
  keyboard.KEY_N: true,
  keyboard.KEY_M: true,
  keyboard.KEY_COMMA: true,
  keyboard.KEY_DOT: true,
  keyboard.KEY_SLASH: true,
}
