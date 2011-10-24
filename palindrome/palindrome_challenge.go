package main

import ( "fmt"; "os"; "bufio"; "regexp"; "strconv" )

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

func reverse(s string) string {
    // Get Unicode code points.
    n := 0
    rune := make([]int, len(s))
    for _, r := range s {
        rune[n] = r
        n++
    }
    rune = rune[0:n]
    // Reverse
    for i := 0; i < n/2; i++ {
        rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
    }
    // Convert back to UTF-8.
    return string(rune)
}

func is_palindrome(i int64) bool {
	s := strconv.Itoa64(i)
	return s == reverse(s)
}

func parse_matches(matches []string) (result int64) {
	if len(matches) < 4 {
		result = 0
		return
	}
	d1, err1 := strconv.Atoi(matches[1])
	d2, err2 := strconv.Atoi(matches[3])
	if err1 != nil || err2 != nil {
		result = 0
		return
	}
	switch matches[2] {
	case "+":
		result = int64(d1) + int64(d2)
	case "-":
		result = int64(d1) - int64(d2)
	case "*":
		result = int64(d1) * int64(d2)
	}
	return result
}


func main() {
	f, err := openExistingFile(os.Args[1])
	defer f.Close()
	if err != nil {
		fmt.Println("error opening file", os.Args[1], ":", err)
		os.Exit(1)
	}

	r, err := regexp.Compile("([0-9]+)[ \t]+([\\-\\+\\*])[ \t]+([0-9]+)")
	if err != nil {
		fmt.Println("error compiling regex", err)
		os.Exit(1)
	}

	var sum int64 = 0
	for line := range eachLine(f) {
		matches := r.FindStringSubmatch(line)
		val := parse_matches(matches)
		if is_palindrome(val) {
			sum += val
		}
	}
	fmt.Println("sum =", sum)
}
