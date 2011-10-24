// Prime contains functions for generating and testing for primes.
// Based on clojure.contrib.lazy-seqs/prime.

package prime

import (
	"math"
)

var wheel [48]int

func init() {
	wheel = [48]int{
		2, 4, 2, 4, 6, 2, 6, 4, 2, 4, 6, 6, 2,  6, 4,  2,
		6, 4, 6, 8, 4, 2, 4, 2, 4, 8, 6, 4, 6,  2, 4,  6,
		2, 6, 6, 4, 2, 4, 6, 2, 6, 4, 2, 4, 2, 10, 2, 10,
	}
}

func primesFrom(n int, wheelIndex int, ch chan int) {
	ch <- 2; ch <- 3; ch <- 5; ch <- 7
	for {
		if IsPrime(n) {
			ch <- n
		}
		n += wheel[wheelIndex]
		wheelIndex = (wheelIndex + 1) % 48
	}
}

func Primes() <-chan int {
	ch := make(chan int)
    go primesFrom(11, 0, ch)
	return ch
}

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n & 1 == 0 {
		return false
	}
	s := int(math.Sqrt(float64(n)))
	for i := 3; i <= s; i += 2 {
		if n % i == 0 {
			return false
		}
	}
	return true
}
