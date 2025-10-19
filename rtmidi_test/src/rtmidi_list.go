package main

import (
	"fmt"
	"github.com/mattrtaylor/go-rtmidi"
	"os"
)

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

	fmt.Printf("Inputs:\n")
	num_ports, err := input.PortCount()
	if err != nil {
		fmt.Println("%s\n", err)
	}
	for i := 0; i < num_ports; i++ {
		name, err := input.PortName(i)
		if err == nil {
			fmt.Printf("  %2d: %s\n", i, name)
		} else {
			fmt.Printf("%s\n", err)
		}
	}

	fmt.Printf("\n")
	fmt.Printf("Outputs:\n")
	num_ports, err = output.PortCount()
	if err != nil {
		fmt.Println("%s\n", err)
	}
	for i := 0; i < num_ports; i++ {
		name, err := output.PortName(i)
		if err == nil {
			fmt.Printf("  %2d: %s\n", i, name)
		} else {
			fmt.Printf("%s\n", err)
		}
	}
}
