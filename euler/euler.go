// Euler contains a few sample Project Euler (http://projecteuler.net/)
// solutions.

package main

import (
  "fmt"; "os"; "strconv"; "math"; "big"
  "prime"
)

// ================ channel filtering funcs ================

// ch2 := take(10, ch)
func take(n int, in <-chan int) <-chan int {
	out := make(chan int, n)
	go func() {
  		for i, num := 0, <-in; i < n; i, num = i+1, <-in { out <- num }
		close(out)
	}()
	return out
}

// ch2 := drop(5, ch)
func drop(n int, in <-chan int) <-chan int {
	// TODO returns orig channel; return new channel instead?
	for i := 0; i < n; i++ { <-in }
	return in
}

func filter(f func (int) bool, in <- chan int) <-chan int {
	ch := make(chan int)
	go func() {
		for n := range in {
			if f(n) { ch <- n }
		}
		close(ch)
	}()
	return ch
}

func max(in <-chan int) (m int) {
	m = math.MinInt32
	for n := range in {
		if n > m { m = n }
	}
	return
}

// ================ int utils ================

func sumOfDigitsD(n int64) int {
	return sumOfDigitsS(strconv.Itoa64(n))
}

func sumOfDigitsS(str string) (sum int) {
	sum = 0
	for _, c := range str { sum += c - '0' }
	return
}

func bigint(n int) *big.Int {
	return big.NewInt(int64(n))
}

// ================ number sequences ================

func triangleNumbers() <-chan int {
	ch := make(chan int)
	go func() { for i := 1; ; i++ { ch <- (i*i + i)/2 } }()
	return ch
}

func squareNumbers() <-chan int {
	ch := make(chan int)
	go func() { for i := 1; ; i++ { ch <- i*i } }()
	return ch
}

func pentagonalNumbers() <-chan int {
	ch := make(chan int)
	go func() { for i := 1; ; i++ { ch <- (3*i*i - i)/2 } }()
	return ch
}

func hexagonalNumbers() <-chan int {
	ch := make(chan int)
	go func() { for i := 1; ; i++ { ch <- 2*i*i - i } }()
	return ch
}

func heptagonalNumbers() <-chan int {
	ch := make(chan int)
	go func() { for i := 1; ; i++ { ch <- (5*i*i - 3*i)/2 } }()
	return ch
}

func octagonalNumbers() <-chan int {
	ch := make(chan int)
	go func() { for i := 1; ; i++ { ch <- 3*i*i - 2*i } }()
	return ch
}

// ================================

func factorial(n int) *big.Int {
	z := big.NewInt(0).MulRange(int64(2), int64(n))
	return z
}

// Return the number of combinatinos of n things taken k at a time.
func numCombinations(n int, k int) int64 {
	z := factorial(n)
	z.Div(z, factorial(k))
	return z.Div(z, factorial(n-k)).Int64()
}

func combinationsInt(a []int) <-chan [][]int {
	ch := make(chan [][]int)
	// for i, n := range a {
	// 	// TODO
	// }
    return ch
}

// ================================

func p56() int {
	ch := make(chan int)
	go func() {
		for a := 2; a < 100; a++ {
			for b := 2; b < 100; b++ {
				z := bigint(0).Exp(bigint(a), bigint(b), nil)
				ch <- sumOfDigitsS(z.String())
			}
		}
		close(ch)
	}()
	return max(ch)
}

// ================================

func main() {
	EULER_FUNCS := map[int] func () int { 56: p56 }

	if len(os.Args) > 1 {
		n, e := strconv.Atoi(os.Args[1])
		if e == nil {
			if EULER_FUNCS[n] == nil {
				fmt.Printf("not yet solved: p%d\n", n) 
			} else {
				fmt.Printf("p%d = %d\n", n, EULER_FUNCS[n]()) 
			}
		} else {
			fmt.Printf("error in command line arg: %s\n", e.String())
		}
	} else {
		fmt.Printf("usage: %s N\n", os.Args[0])
		fmt.Printf("  where N is a Project Euler problem number\n");
	}
}

func justFoolingAround() {
	if len(os.Args) > 1 {
		n, e := strconv.Atoi(os.Args[1])
		if e == nil {
			str := " not"
			if prime.IsPrime(n) { str = "" }
			fmt.Printf("%d is%s prime\n", n, str)
		} else {
			fmt.Printf("error: %s\n", e.String())
		}
	}

	fmt.Println("The next 10 octagonal numbers after the first 5. (take, drop)")
	ch := octagonalNumbers()
	for h := range take(10, drop(5, ch)) { fmt.Printf("%d ", h) }
	fmt.Printf("\n")

	fmt.Println("The first 200 primes. (take, prime.Primes)")
	for p := range take(200, prime.Primes()) { fmt.Printf("%d ", p) }
	fmt.Printf("\n")

	fmt.Println("The primes > 100 in the list of the first 200 primes. (filter, take, prime.Primes)")
	for p := range filter(func(n int) bool { return n > 100 }, take(200, prime.Primes())) {
		fmt.Printf("%d ", p) 
	}
	fmt.Printf("\n")

	fmt.Printf("max digit sum for a^b where a, b < 100: %d\n", p56())

	fmt.Printf("num combinations of 5 things taken 2 at a time is", numCombinations(5, 2))
}
