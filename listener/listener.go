package listener

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"syscall"
	"unsafe"

	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)

// EventType represents the types of keyboard events we may receive.
type EventType int32

// Relevant event types for Vimprover, required to match Value number in inputEvent struct.
const (
  KeyUp EventType = 0
  KeyDown EventType = 1
  KeyHold EventType = 2
)


type KeyEvent struct {
  KeyCode uint16
  EventType EventType
}

// RunListener establishes a keyboard listener that calls provided callback
// for each key event.
//
// The function attempts to detect which device to listen to, and may require
// user input if multiple potential devices were detected.
func RunListener(callback func(KeyEvent)) error {
  if !isRoot() {
    return status.Error(
      codes.PermissionDenied,
      "since vimprover runs a keylistener it requires root privileges.")
  }
  device, err := selectDevice()
  if err != nil {
    return err
  }
  listenToDevice(device, callback)
  return nil
}

// Device is used to represent an input device.
type Device struct {
  path string
  name string
}

// selectDevice returns the input device to listen for keystrokes from.
func selectDevice() (*Device, error) {
  devices := avaliableDevices()
  if len(devices) == 0 {
    return nil, status.Error(codes.NotFound, "Unable to find input device to listen to")
  }
  if len(devices) == 1 {
    fmt.Printf("Using Device: %s", devices[0].name)
    return &devices[0], nil
  }
  return selectDeviceFromMultiple(devices), nil
}

// availableDevices returns all input devices avaliable to listen to.
func avaliableDevices() []Device {
  inputRootDir := "/sys/class/input/"
  files, err := ioutil.ReadDir(inputRootDir)
  if err != nil {
    log.Fatal(err)
  }

  devices := make([]Device, 0)
  for _, inputDir := range files {
    namePath := fmt.Sprintf("%s%s/device/name", inputRootDir, inputDir.Name())
    // Ignore errors.
    buff, _ := ioutil.ReadFile(namePath)
    deviceName := string(buff)
    if !strings.Contains(strings.ToLower(deviceName), "key") {
      continue
    }
    devices = append(devices, Device{
      path: fmt.Sprintf("/dev/input/%s", inputDir.Name()),
      name: deviceName,
    })
  }
  return devices
}

// selectDeviceFromMultiple shows options for the user to select a device and returns selected device.
func selectDeviceFromMultiple(devices []Device) *Device {
  for {
    fmt.Printf("Select a device\n")
    for i, device := range devices {
      fmt.Printf("[%d]: %s", i, device.name)
    }
    reader := bufio.NewReader(os.Stdin)
    line, _, _ := reader.ReadLine()
    selection, _ := strconv.Atoi(string(line))

    if selection < 0 || selection >= len(devices) {
      continue
    }

    fmt.Printf("Selected: [%d] %s", selection, devices[selection].name)
    return &devices[selection]
  }
}

// IsRoot checks if the process is run with root permission
func isRoot() bool {
	return syscall.Getuid() == 0
}

// listenToDevice starts a listener for input events on the provided device.
func listenToDevice(device *Device, callback func(KeyEvent)) {
  file, err := os.Open(device.path)
  if err != nil {
    log.Fatal(err)
  }
  for {
    ie, err := readEvent(file)
    if err != nil {
      log.Fatal(err)
      break
    }
    if ie != nil {
      callback(KeyEvent{
        EventType: EventType(ie.Value),
        KeyCode: ie.Code,
      })
    }
  }
}

// inputEvent represents a keyboard input event.
//
// The inputEvent struct has to match the Kernel input event:
// https://www.kernel.org/doc/Documentation/input/input.txt
type inputEvent struct {
	Time  syscall.Timeval
	Type  uint16
  // The event code includes what key was pressed, see: include/uapi/linux/input-event-codes.h
	Code  uint16
  // Value represents the value the event carries. e.g. KEY_DOWN.
	Value int32
}

// Reads the last input event from file.
func readEvent(fd *os.File) (*inputEvent, error) {
	buffer := make([]byte, int(unsafe.Sizeof(inputEvent{})))
	_, err := fd.Read(buffer)
	if err != nil {
		return nil, err
	}
  event := &inputEvent{}
	binary.Read(bytes.NewBuffer(buffer), binary.LittleEndian, event)
  // Ignores other types of system event (not keyboard).
  if event.Type != 1 {
    return nil, nil
  }
  return event, nil
}
