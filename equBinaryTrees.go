package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var WalkHelper func(t *tree.Tree, ch chan int)
	WalkHelper = func(t *tree.Tree, ch chan int) {
		if t.Left != nil {
			WalkHelper(t.Left, ch)
		}
		ch <- t.Value
		if t.Right != nil {
			WalkHelper(t.Right, ch)
		}
	}
	WalkHelper(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		i1, ok1 := <-ch1
		i2, ok2 := <-ch2
		if i1 != i2 || ok1 != ok2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}

func main() {
	result1 := Same(tree.New(1), tree.New(1))
	result2 := Same(tree.New(1), tree.New(2))
	fmt.Println("result1 should be true: ", result1)
	fmt.Println("result2 should be false: ", result2)
}
