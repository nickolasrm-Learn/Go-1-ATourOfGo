package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func walkRecursive(t *tree.Tree, ch chan int) {
	if t != nil {
		walkRecursive(t.Left, ch)
		ch <- t.Value
		walkRecursive(t.Right, ch)
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkRecursive(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	var v1, v2 int
	run1, run2 := true, true
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for run1 && run2 {
		v1, run1 = <-ch1
		v2, run2 = <-ch2
		if v1 != v2 {
			return false
		}
	}
	if !run1 && !run2 {
		return true
	}
	return false
}

func main() {
	t1, t2, t3 := tree.New(1), tree.New(1), tree.New(2)
	fmt.Println(Same(t1, t2))
	fmt.Println(Same(t1, t3))
}
