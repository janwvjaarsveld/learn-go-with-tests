package linkedList

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	head *Node
	size int
}

func (l *LinkedList) Add(d int) {
	list := &Node{value: d, next: nil}
	l.size += 1
	if l.head == nil {
		l.head = list
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = list
	}
}

func (l *LinkedList) Show() {
	curr := l.head
	for curr != nil {
		fmt.Printf("-> %v ", curr.value)
		curr = curr.next
	}
	fmt.Println()
}

func (l *LinkedList) Remove(idx int) (int, error) {
	if idx > l.size {
		return 0, errors.New("index out of bounds")
	} else {
		curr := l.head
		var prev *Node
		currIdx := 0
		for currIdx != idx {
			prev = curr
			curr = curr.next
			currIdx += 1
		}
		prev.next = curr.next
		return curr.value, nil
	}
}
