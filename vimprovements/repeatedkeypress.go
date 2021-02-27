package vimprovements

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"oreflow.com/vimprover/keyboard"
)

var rkViolation = Violation{
  Type: RepeatedKeypressViolation,
  message: "Repetition of keys H, J, K, L was detected",
}

// Limit for how many repeated presses are considered a violation.
const repeatedKeypressLimit = 5

// repeatedKeypress detects if any of the keys H,J,K,L was pressed too many
// times consecutively.
func repeatedKeypress(s *keyboard.State) (*Violation, error) {
  if len(s.LastCharacters) < 1 {
    return nil, status.Error(
      codes.FailedPrecondition,
      "repeatedKeypress vimprovement check ran before character input was received.",
    )
  }
  lastKey := s.LastCharacters[0]
  if !keysToCheck[lastKey] {
    return nil, nil
  }
  repetitions := 0
  for _, k := range s.LastCharacters {
    if k != lastKey {
      return nil, nil
    }
    repetitions++
    if repetitions >= repeatedKeypressLimit {
      return &rkViolation, nil
    }
  }
  return nil, nil
}

var keysToCheck = map[uint16]bool{
  keyboard.KEY_H: true,
  keyboard.KEY_J: true,
  keyboard.KEY_K: true,
  keyboard.KEY_L: true,
}
