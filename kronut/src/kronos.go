package kronut

import (
	"bytes"
	"github.com/mattrtaylor/go-rtmidi"
)

const MIDI_BUFSIZ = 1024
const SYSEX_BUF_EVENTS = 1024
const SYSEX_START_TIMEOUT_SECS = 5
const SYSEX_READ_TIMEOUT_SECS = 60
const TIMEOUT_ERROR_REPLY = 100

func sysexHeader() [4]uint8 {
	return [...]uint8{0xf0, 0x42, 0x30, 0x68}
}

// var ERROR_REPLY_MESSAGES = [...]string{
// 	"no error",
// 	"parameter type specified is incorrect for current mode",
// 	"unknown param message type, unknown parameter id or index",
// 	"short or otherwise mangled message",
// 	"target object not found",
// 	"insufficient resources to complete request",
// 	"parameter value is out of range",
// 	"(internal error code)",
// 	"other error: program bank is wrong type for received program dump (Func 73, 75); invalid data in Preset Pattern Dump (Func 7B).",
// 	"target object is protected",
// 	"memory overflow",
// 	"(unknown error code)",
// 	// Not Kronos errors, but kronut errors
// 	"timeout",
// }

type Kronos struct {
	input   rtmidi.MIDIIn
	output  rtmidi.MIDIOut
	channel uint8
	sysex   bytes.Buffer
}
