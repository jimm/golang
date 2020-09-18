package main

import (
	"fmt"
	"github.com/rakyll/portmidi"
)

func main() {
	portmidi.Initialize()

	num_devices := portmidi.CountDevices()
	devices := make([]*portmidi.DeviceInfo, num_devices)
	for i := 0; i < num_devices; i++ {
		devices[i] = portmidi.Info(portmidi.DeviceID(i))
	}

	fmt.Printf("Inputs:\n")
	for i := 0; i < num_devices; i++ {
		if devices[i].IsInputAvailable {
			fmt.Printf("  %s\n", devices[i].Name)
		}
	}

	fmt.Printf("Outputs:\n")
	for i := 0; i < num_devices; i++ {
		if devices[i].IsOutputAvailable {
			fmt.Printf("  %s\n", devices[i].Name)
		}
	}

	portmidi.Terminate()
}
