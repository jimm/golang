package main

import (
	"bytes"
	"github.com/mattrtaylor/go-rtmidi"
)

SYSEX_HEADER []uint8 := [0xf0, 0x42, 0x30, 0x68]
MIDI_BUFSIZ := 1024
SYSEX_BUF_EVENTS := 1024
SYSEX_START_TIMEOUT_SECS := 5
SYSEX_READ_TIMEOUT_SECS := 60
TIMEOUT_ERROR_REPLY := 100

ERROR_REPLY_MESSAGES := [
  "no error",
  "parameter type specified is incorrect for current mode",
  "unknown param message type, unknown parameter id or index",
  "short or otherwise mangled message",
  "target object not found",
  "insufficient resources to complete request",
  "parameter value is out of range",
  "(internal error code)",
  "other error: program bank is wrong type for received program dump (Func 73, 75); invalid data in Preset Pattern Dump (Func 7B).",
  "target object is protected",
  "memory overflow",
  "(unknown error code)",
  // The following errors are kronut errors, not Kronos errors
  "timeout"
]

func isRealtime(b: uint8) bool { return b >= 0xf8 }

type Kronos struct {
	input  rtmidi.MIDIIn
	output rtmidi.MIDIOut
	uint8  channel
	sysex  bytes.Buffer
}

type SysexState int
const (
	Waiting SysexState = 1,
	Receiving
	Received
	Error
)

