package main

import "fmt"

type LinkedList struct {
	Value int
	Next  *LinkedList
}

var node5 = LinkedList{
	5,
	nil,
}

var node3 = LinkedList{
	3,
	&node5,
}

var node7 = LinkedList{
	7,
	&node3,
}

var node2 = LinkedList{
	2,
	&node3,
}

func main() {
	middle := MiddleNode(&node2)
	fmt.Printf("%v", middle)
}

func MiddleNode(linkedList *LinkedList) *LinkedList {
	slow := linkedList
	fast := linkedList
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
