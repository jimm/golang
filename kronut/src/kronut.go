package main

import (
	"errors"
	"fmt"
	"github.com/rakyll/portmidi"
	"strings"
)

func find_kronos_input_num() (int, error) {
	for i := 0; i < portmidi.CountDevices(); i++ {
		info := portmidi.Info(portmidi.DeviceID(i))
		if info.IsInputAvailable && strings.EqualFold(info.Name, "kronos keyboard") {
			return i, nil
		}
	}
	return -1, errors.New("Kronos input not found")
}

func find_kronos_output_num() (int, error) {
	for i := 0; i < portmidi.CountDevices(); i++ {
		info := portmidi.Info(portmidi.DeviceID(i))
		if info.IsOutputAvailable && strings.EqualFold(info.Name, "kronos sound") {
			return i, nil
		}
	}
	return -1, errors.New("Kronos output not found")
}

func main() {
	defer portmidi.Terminate()

	portmidi.Initialize()

	input_num, err := find_kronos_input_num()
	if err != nil {
		fmt.Printf("error: %s\n", err)
	} else {
		fmt.Printf("Kronos input number: %d\n", input_num)
	}
}
