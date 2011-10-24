package main

import ( "fmt"; "rand"; "strconv" )

const (
	kNumLines = 10000
	kMaxNum = 1000000
	)

func gen_line() string {
	op := ""
	switch rand.Intn(3) {
	case 0:
		op = "+"
	case 1:
		op = "-"
	case 2:
		op = "*"
	}
	n1 := strconv.Itoa(rand.Intn(kMaxNum))
	n2 := strconv.Itoa(rand.Intn(kMaxNum))
	if op == "-" && n1 < n2 {
		n1, n2 = n2, n1
	}
	prefix := ""
	s1 := " "
	s2 := " "
	suffix := ""
	switch rand.Intn(200) {
	case 0:
		op = ""
	case 1:
		n1 = ""
	case 2:
		n2 = ""
	case 3:
		s1 = ""
	case 4:
		s2 = ""
	case 5:
		prefix = "   "
	case 6:
		suffix = "   "
	case 7:
		prefix = ""
		n1 = ""
		s1 = ""
		op = ""
		s2 = ""
		n2 = ""
		suffix = ""
	case 8:
		n1 = "foo"
	case 9:
		n2 = "bar"
	}
	return prefix + n1 + s1 + op + s2 + n2 + suffix
}

func main() {
	for i := 0; i < kNumLines; i++ { fmt.Println(gen_line()) }
}
