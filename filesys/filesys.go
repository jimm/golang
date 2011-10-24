package main

import ( "fmt"; "os"; "bufio" )

func openExistingFile(filename string) (f *os.File, err os.Error) {
	f, err = os.Open(filename)
	if f == nil {
		err = os.ENOENT
	}
	return
}

// Returns a channel that outputs lines in filename.
func eachLine(f *os.File) <-chan string {
	out := make(chan string)
	go func() {
		in := bufio.NewReader(f)
		line_so_far := ""
		for {
			bytes, isPrefix, err := in.ReadLine()
			if err != nil {
				break
			} else if isPrefix { // incomplete line
				line_so_far += string(bytes)
			} else {			// complete line or remainder of line
				out <- line_so_far + string(bytes)
				line_so_far = ""
			}
		}
		close(out)
    }()
	return out
}

func main() {
	f, err := openExistingFile(os.Args[1])
	defer f.Close()
	if err != nil {
		fmt.Println("error opening file", os.Args[1], ":", err)
		os.Exit(1)
	}

	for line := range eachLine(f) { fmt.Println(line) }
}
