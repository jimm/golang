package main

import (
	"fmt"
	"github.com/rakyll/portmidi"
)

func main() {
	portmidi.Initialize()
	num_devices := portmidi.CountDevices()
	fmt.Println("num devices:", num_devices)
	portmidi.Terminate()
}
