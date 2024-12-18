package main

import (
	"fmt"
)

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

var testSubject = []int{
	1, 2, 5, 7, 10, 13, 14, 15, 22,
}

func main() {
	result := MinHeightBst(testSubject)
	fmt.Println(result)
}

func MinHeightBst(array []int) *BST {
	return &BST{
		Value: len(array) / 2,
	}
}
