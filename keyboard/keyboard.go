package keyboard

import (
	"oreflow.com/vimprover/listener"
)

// maxHistoryLength defines how many characters to keep track of in history.
//
// Since vimprover will monitor all keypresses this should be kept to a small
// number leading to never holding a full password in the input history.
const maxHistoryLength = 6

// Representation of Keyboard state to detect Vimprovements.
type State struct {
  RightShiftDown bool
  LeftShiftDown bool
  LastCharacters []uint16
}

// Creates a new State.
//
// Note: Shift key status in newly created State will not be accurate as
// Vimprover is not reading state.
func NewKeyboardState() *State {
  return &State{
    RightShiftDown : false,
    LeftShiftDown : false,
    LastCharacters : make([]uint16, maxHistoryLength),
  }
}


// Adds a new KeyEvent to the state, updating affected values.
func (s *State) AddEvent(ke listener.KeyEvent) (*State) {
  if ke.EventType == listener.KeyDown || ke.EventType == listener.KeyHold {
    // Destructuring Last characters to prevent longer keyboard sequences in memory.
    s.LastCharacters = append([]uint16{ke.KeyCode}, s.LastCharacters...)[0:maxHistoryLength]
  }
  if ke.KeyCode == KEY_RIGHTSHIFT && ke.EventType == listener.KeyDown {
    s.RightShiftDown = true
  }
  if ke.KeyCode == KEY_RIGHTSHIFT && ke.EventType == listener.KeyUp {
    s.RightShiftDown = false
  }
  if ke.KeyCode == KEY_LEFTSHIFT && ke.EventType == listener.KeyDown {
    s.LeftShiftDown = true
  }
  if ke.KeyCode == KEY_LEFTSHIFT && ke.EventType == listener.KeyUp {
    s.LeftShiftDown = false
  }
  return s
}

const(
  // KeyCodes
  // From: https://github.com/torvalds/linux/blob/master/include/uapi/linux/input-event-codes.h
  KEY_RESERVED =		0
  KEY_ESC =			1
  KEY_1 =			2
  KEY_2 =			3
  KEY_3 =			4
  KEY_4 =			5
  KEY_5 =			6
  KEY_6 =			7
  KEY_7 =			8
  KEY_8 =			9
  KEY_9 =			10
  KEY_0 =			11
  KEY_MINUS =		12
  KEY_EQUAL =		13
  KEY_BACKSPACE =		14
  KEY_TAB =			15
  KEY_Q =			16
  KEY_W =			17
  KEY_E =			18
  KEY_R =			19
  KEY_T =			20
  KEY_Y =			21
  KEY_U =			22
  KEY_I =			23
  KEY_O =			24
  KEY_P =			25
  KEY_LEFTBRACE =		26
  KEY_RIGHTBRACE =		27
  KEY_ENTER =		28
  KEY_LEFTCTRL =		29
  KEY_A =			30
  KEY_S =			31
  KEY_D =			32
  KEY_F =			33
  KEY_G =			34
  KEY_H =			35
  KEY_J =			36
  KEY_K =			37
  KEY_L =			38
  KEY_SEMICOLON =		39
  KEY_APOSTROPHE =		40
  KEY_GRAVE =		41
  KEY_LEFTSHIFT =		42
  KEY_BACKSLASH =		43
  KEY_Z =			44
  KEY_X =			45
  KEY_C =			46
  KEY_V =			47
  KEY_B =			48
  KEY_N =			49
  KEY_M =			50
  KEY_COMMA =		51
  KEY_DOT =			52
  KEY_SLASH =		53
  KEY_RIGHTSHIFT =		54
  KEY_KPASTERISK =		55
  KEY_LEFTALT =		56
  KEY_SPACE =		57
  KEY_CAPSLOCK =		58
  KEY_F1 =			59
  KEY_F2 =			60
  KEY_F3 =			61
  KEY_F4 =			62
  KEY_F5 =			63
  KEY_F6 =			64
  KEY_F7 =			65
  KEY_F8 =			66
  KEY_F9 =			67
  KEY_F10 =			68
  KEY_NUMLOCK =		69
  KEY_SCROLLLOCK =		70
  KEY_KP7 =			71
  KEY_KP8 =			72
  KEY_KP9 =			73
  KEY_KPMINUS =		74
  KEY_KP4 =			75
  KEY_KP5 =			76
  KEY_KP6 =			77
  KEY_KPPLUS =		78
  KEY_KP1 =			79
  KEY_KP2 =			80
  KEY_KP3 =			81
  KEY_KP0 =			82
  KEY_KPDOT =		83
  KEY_ZENKAKUHANKAKU =	85
  KEY_102ND =		86
  KEY_F11 =			87
  KEY_F12 =			88
  KEY_RO =			89
  KEY_KATAKANA =		90
  KEY_HIRAGANA =		91
  KEY_HENKAN =		92
  KEY_KATAKANAHIRAGANA =	93
  KEY_MUHENKAN =		94
  KEY_KPJPCOMMA =		95
  KEY_KPENTER =		96
  KEY_RIGHTCTRL =		97
  KEY_KPSLASH =		98
  KEY_SYSRQ =		99
  KEY_RIGHTALT =		100
  KEY_LINEFEED =		101
  KEY_HOME =		102
  KEY_UP =			103
  KEY_PAGEUP =		104
  KEY_LEFT =		105
  KEY_RIGHT =		106
  KEY_END =			107
  KEY_DOWN =		108
  KEY_PAGEDOWN =		109
  KEY_INSERT =		110
  KEY_DELETE =		111
  KEY_MACRO =		112
  KEY_MUTE =		113
  KEY_VOLUMEDOWN =		114
  KEY_VOLUMEUP =		115
  KEY_POWER =		116	/* SC System Power Down */
  KEY_KPEQUAL =		117
  KEY_KPPLUSMINUS =		118
  KEY_PAUSE =		119
  KEY_SCALE =		120	/* AL Compiz Scale (Expose) */
  KEY_KPCOMMA =		121
  KEY_HANGEUL =		122
  KEY_HANGUEL =		KEY_HANGEUL
  KEY_HANJA =		123
  KEY_YEN =			124
  KEY_LEFTMETA =		125
  KEY_RIGHTMETA =		126
  KEY_COMPOSE =		127
)