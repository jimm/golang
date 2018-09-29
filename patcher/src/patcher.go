package main

import (
	"flag"
	"fmt"
	"github.com/rakyll/portmidi"
	"os"
)

func cleanup() {
	portmidi.Terminate()
}

func initialize() {
	portmidi.Initialize()
}

func usage(progname string) {
	fmt.Printf("%s: TODO: usage (try to use flag module)\n", progname)
}

func load(fname string, testing bool) {
}

func run_web() {
}

func run() {
}

func list_devices(title string, infos map[int]*portmidi.DeviceInfo, num_devices int) {
 	fmt.Println(title)
	for i := 0; i < num_devices; i++ {
		info, ok := infos[i]
		if !ok {
			continue
		}

		name := info.Name
		q := ""
		if name[0] == ' ' || name[len(name)-1] == ' ' {
			q = "\""
		}
		opened := ""
		if info.IsOpened {
			opened = " (open)"
		}
		fmt.Printf("  %2d: %s%s%s%s\n", i, q, name, q, opened)
	}
}

func list_all_devices() {
	inputs := make(map[int]*portmidi.DeviceInfo)
	outputs := make(map[int]*portmidi.DeviceInfo)

	num_devices := portmidi.CountDevices()
	for i := 0; i < num_devices; i++ {
		info := portmidi.Info(portmidi.DeviceID(i))
		if info.IsInputAvailable {
			inputs[i] = info
		}
		if info.IsOutputAvailable {
			outputs[i] = info
		}
	}

	list_devices("Inputs", inputs, num_devices)
	list_devices("Outputs", outputs, num_devices)
}

func main() {
	list_ports_ptr := flag.Bool("list-ports", false, "List all attached MIDI ports")
	testing_ptr := flag.Bool("no-midi", false, "No MIDI (ignores bad/unknown MIDI ports")
	web_ptr := flag.Bool("web", false, "Use web interface on port 8080")
	flag.Parse()

	if *list_ports_ptr {
		list_all_devices()
		os.Exit(0)
	}

	if len(flag.Args()) == 0 {
		usage(os.Args[0])
	}

	initialize()
	load(os.Args[0], *testing_ptr)
	if *web_ptr {
		run_web()
	} else {
		run()
	}

	cleanup()
}
