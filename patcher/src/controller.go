package patcher

import "portmidi"

type Controller struct {
	CcNum           byte
	TranslatedCcNum byte
	Min             byte
	Max             byte
	Filtered        bool
}

func NewController(cc_num byte) Controller {
	c := Controller()
	c.CcNum = cc_num
	c.translation = cc_num
	c.Min = 0
	c.Max = 127
	c.Filtered = false
	return c
}

func (c Controller) WillModify() bool {
	return c.Filtered ||
		c.TranslatedCcNum != c.CcNum ||
		c.Min != 0 ||
		c.Max != 127
}

// bool returns false if nothing to send
func (c Controller) Process(bytes [3]byte, output_channel byte) ([]byte, bool) {
	if c.Filtered {
		return bytes, false
	}

	status := bytes[0]
	data1 := c.TranslatedCcNum
	data2 = c.clamp(bytes[2])

	if output_channel != 0xff {
		status = (status & 0xf0) | output_chan
	}

	return [3]byte(status, data1, data2), true
}

func (c Controller) clamp(val byte) byte {
	if val < c.Min {
		val = c.Min
	}
	if val > c.Max {
		val = c.Max
	}
	return val
}
