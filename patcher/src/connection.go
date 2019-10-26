package patcher

type Program struct {
	BankMSB byte
	BankLSB byte
	Prog byte
}

type Zone struct {
	Low byte
	High byte
}

type Connection struct {
	Input *Input
	Output *Output
	InputChan byte
	OutputChan byte
	Prog Program
	Zone Zone
	Xpose byte
	CcMaps [128]Controller
}

func NewConnection(input *Input, inputChan byte, output *Output, outputChan byte) Connection {
	c := Connection()
	c.Input = input
	c.InputChan = inputChan
	c.Output = output
	c.OutputChan = outputChan
	c.Zone = Zone(0, 0xff)

	for i := 0; i < 128; i++ {
		c.CcMaps[i] = NewController(i)
	}

	return c
}

func (c Connection) Start(start_messages [][3]byte) {
}

func (c Connection) Stop(stop_messages [][3]byte) {
}

func (c Connection) MidiIn(msg [3]byte) {
}

func (c Connection) acceptFromInput(msg [3]byte) bool {
	return true
}

func (c Connection) insideZone(msg [3]byte) bool {
	return true
}

func (c Connection) midiOut(msg [3]byte) {
}
