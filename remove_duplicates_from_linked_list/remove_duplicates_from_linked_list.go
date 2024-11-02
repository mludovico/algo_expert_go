package main

import (
	"fmt"
)

type LinkedList struct {
	Value int
	Next  *LinkedList
}

var node52 = LinkedList{
	6,
	nil,
}
var node5 = LinkedList{
	6,
	&node52,
}

var node4 = LinkedList{
	5,
	&node5,
}

var node33 = LinkedList{
	4,
	&node4,
}

var node32 = LinkedList{
	4,
	&node33,
}

var node3 = LinkedList{
	4,
	&node32,
}

var node2 = LinkedList{
	3,
	&node3,
}

var node13 = LinkedList{
	1,
	&node2,
}

var node12 = LinkedList{
	1,
	&node13,
}

var node1 = LinkedList{
	1,
	&node12,
}

func main() {
	cleanSlice := RemoveDuplicatesFromLinkedList(&node1)
	for node := cleanSlice; node != nil; node = node.Next {
		fmt.Printf("%v", node.Value)
		if node.Next != nil {
			fmt.Printf(" -> ")
		}
	}
}

func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {
	for listNode := linkedList; listNode.Next != nil; listNode = listNode.Next {
		for listNode.Value == listNode.Next.Value {
			listNode.Next = listNode.Next.Next
		}
	}
	return linkedList
}
