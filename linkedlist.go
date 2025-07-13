package main

import (
	"fmt"
)

// List represents a singly-linked list that holds
// values of any type.
type List[T comparable] struct {
	next *List[T]
	val  T
}

func (l *List[T]) show() {
	for {
		if l.next == nil {
			fmt.Printf("%v\n", l.val)
			return
		}
		fmt.Printf("%v -> ", l.val)
		l = l.next
	}
}

func (l *List[T]) add(x T) {
	newList := List[T]{val: x, next: nil}
	for {
		if l.next == nil {
			l.next = &newList
			return
		}
		l = l.next
	}
}

func (l *List[T]) contains(x T) bool{
	for l != nil {
		if l.val == x{
			return true
		}
		l = l.next
	}
	return false
}

func main() {
	myLink := List[int]{val: 1, next: nil}
	myLink.add(2)
	myLink.add(3)
	myLink.add(4)
	myLink.show()
	fmt.Println(myLink.contains(4))
	fmt.Println(myLink.contains(20))

}
