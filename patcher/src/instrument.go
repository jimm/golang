package patcher

import "github.com/rakyll/portmidi"

const MidiBufsize = 128

type Instrument struct {
	Sym           string
	Name          string
	PortNum       portmidi.DeviceID
	Stream        *portmidi.Stream
	IoMessages    [MidiBufsize][3]byte // testing only
	NumIoMessages int                  // ditto
}

func NewInstrument(sym_str string, name string, portmidi_port_num portmidi.DeviceID) Instrument {
	return Instrument(sym_str, name, portmidi_port_num, nil, nil, 0)
}

func (inst Instrument) RealPort() bool {
	return inst.PortNum != -1
}

// testing only
func (inst Instrument) Clear() {
	inst.NumIoMessages = 0
}
