package linkedlist

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Linkedlist struct {
	head   *Node
	length int
}

func (l *Linkedlist) InsertAtBeggning(val int) {
	n := Node{}
	n.value = val
	n.next = l.head
	l.head = &n
	l.length++
}

func (l *Linkedlist) Print() {
	if l.length == 0 {
		fmt.Println("No nodes in list")
	}
	ptr := l.head
	for i := 0; i < l.length; i++ {
		fmt.Println("Node: ", ptr)
		ptr = ptr.next
	}
}
