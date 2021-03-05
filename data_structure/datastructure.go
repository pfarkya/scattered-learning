package main

import (
	"fmt"

	linkedlist "github.com/pfarkya/scattered-learning/data_structure/linkedlist/go"
)

func main() {
	fmt.Println("Hello, playground")
	l := linkedlist.Linkedlist{}
	l.InsertAtBeggning(5)
	l.InsertAtBeggning(4)
	l.Print()
}
