package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()
	ch2 := make(chan int, 10)
	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	for {
		out1, ok1 := <- ch1
		out2, ok2 := <- ch2
		if ok1 != ok2 {
			return false
		}
		if !ok1{
			break
		}
		
		if out1 != out2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int, 10)
	myTree := tree.New(1)
	//fmt.Printf("my tree: %v\n", myTree)	
	go func() {
		Walk(myTree, ch)
		close(ch)
	}()
	//for i := range ch {
	//	fmt.Println(i)
	//}
	
	t1 := tree.New(1)
	t2 := tree.New(1)
	fmt.Println(Same(t1, t2))
	
}
