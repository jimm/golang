package main

type MIDIData struct {
	MIDIBytes     []byte // TODO use bytes.Buffer
	InternalBytes []byte
	MIDILen       int // TODO use bytes.Buffer
	InternalLen   int
}

func (md MIDIData) midi_changed() {
	md.midi_to_internal()

}

func (md MIDIData) internal_changed() {
	md.internal_to_midi()
}

func (md MIDIData) midi_to_internal() {
	mlen := md.MIDILen
	midi := md.MIDIBytes
	internal := md.InternalBytes

	for mlen > 0 {
		chunk_len := 8
		if mlen < 8 {
			chunk_len = mlen
		}
		for i = 0; i < chunk_len-1; i++ {
			internal[i] = midi[i+1]
			if (midi[0] & (1 << i)) != 0 {
				internal[i] += 0x80
			}
		}
		midi += chunk_len
		mlen -= chunk_len
		internal += chunk_len - 1
	}
}

func (md MIDIData) internal_to_midi() {
	ilen := md.InternalLen
	internal := md.InternalBytes
	midi := md.MIDIBytes

	for ilen > 0 {
		chunk_len := ilen
		if chunk_len < 7 {
			chunk_len = 7
		}
		midi[0] = 0
		for i = 0; i < chunk_len; i++ {
			if (internal[i] & 0x80) != 0 {
				midi[0] += (1 << i)
			}
			midi[i+1] = internal[i] & 0x7f
		}
		internal += chunk_len
		ilen -= chunk_len
		midi += chunk_len + 1
	}
}

func (md MIDIData) midi_to_internal_len() {
	md.InternalLen = (md.MIDILen / 8) * 7
	if (md.MIDILen % 8) != 0 {
		md.InternalLen += md.MIDILen%8 - 1
	}
}

func (md MIDIData) internal_to_MIDILen() {
	md.MIDILen = md.InternalLen + (md.InternalLen+6)/7
}
