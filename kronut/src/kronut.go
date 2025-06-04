package main

import (
	"errors"
	"fmt"
	"github.com/mattrtaylor/go-rtmidi"
	"os"
	"strings"
)

func find_kronos_input_num(input rtmidi.MIDIIn) (int, error) {
	num_ports, err := input.PortCount()
	if err != nil {
		return -1, err
	}

	for i := 0; i < num_ports; i++ {
		name, err := input.PortName(i)
		if err != nil {
			return -1, err
		}
		if strings.EqualFold(name, "kronos keyboard") {
			return i, nil
		}
	}
	return -1, errors.New("Kronos input not found")
}

func find_kronos_output_num(output rtmidi.MIDIOut) (int, error) {
	num_ports, err := output.PortCount()
	if err != nil {
		return -1, err
	}

	for i := 0; i < num_ports; i++ {
		name, err := output.PortName(i)
		if err != nil {
			return -1, err
		}
		if strings.EqualFold(name, "kronos sound") {
			return i, nil
		}
	}
	return -1, errors.New("Kronos output not found")
}

func main() {
	input, err := rtmidi.NewMIDIInDefault()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(0)
	}

	output, err := rtmidi.NewMIDIOutDefault()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(0)
	}

	input_num, err := find_kronos_input_num(input)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	} else {
		fmt.Printf("Kronos input number: %d\n", input_num)
	}

	output_num, err := find_kronos_output_num(output)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	} else {
		fmt.Printf("Kronos output number: %d\n", output_num)
	}
}
