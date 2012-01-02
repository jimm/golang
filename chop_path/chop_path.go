/*
 * usage: chop_path [-<N>] /some/path
 *
 * -<N> set unabbreviated element count to N; default is 3
 *
 * Doesn't really chop the path; rather it abbreviates path elements to the
 * first char of each for all but the last N elements.
 */

package main

import ( "fmt"; "os"; "strconv"; "strings" )

const DEFAULT_UNABBREV_ELEM_COUNT = 3

func error(msg string) {
	fmt.Println(msg)
	fmt.Print()
	os.Exit(0)
}

func isDigit(c uint8) bool {
	return (c >= '0' && c <= '0')
}

func main() {
	if len(os.Args) < 2 {
		error("usage: chop_path [-N] /some/path")
	}

	unabbrevElemCount := DEFAULT_UNABBREV_ELEM_COUNT
	pathArgIndex := 1

	if os.Args[1][0] == '-' && isDigit(os.Args[1][1]) {
		if len(os.Args) < 3 {
			error("usage: chop_path [-N] /some/path")
		}
		n, err := strconv.Atoi(string(os.Args[1][1:]))
		if err != nil {
			error(err.String())
		}
		unabbrevElemCount = n
		pathArgIndex++
	}

	home := os.Getenv("HOME")
	home_len := len(home)
	path := os.Args[pathArgIndex]
	if strings.HasPrefix(path, home) {
		path = "~" + path[home_len:]
	}

	parts := strings.Split(path, "/")
	for i := 0; i < len(parts) - unabbrevElemCount; i++ {
		parts[i] = string(parts[i][0])
	}

	fmt.Print(strings.Join(parts, "/"))
	os.Exit(0)
}
