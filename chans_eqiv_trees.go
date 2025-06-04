// https://go.dev/tour/concurrency/8

package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	var walk func(t *tree.Tree)
	walk = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walk(t.Left)
		ch <- t.Value
		walk(t.Right)
	}
	walk(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if ok1 && ok2 {
			if v1 != v2 {
				return false // different value
			}
		} else if !ok1 && !ok2 {
			return true // all same, same size trees
		} else {
			return false // different size trees
		}
	}
}

func main() {
	// Test Walk. Should print 0-9
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for range 10 {
		fmt.Println(<-ch)
	}

	// Test Same. Should print true then false
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
