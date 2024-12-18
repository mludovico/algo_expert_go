package main

import (
	"fmt"
)

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

type TestCase struct {
	Bst   *BST
	Array []int
}

var testSubject = TestCase{
	&BST{
		Value: 10,
		Left: &BST{
			Value: 5,
			Left: &BST{
				Value: 2,
				Left: &BST{
					Value: 1,
				},
			},
			Right: &BST{
				Value: 5,
			},
		},
		Right: &BST{
			Value: 15,
			Right: &BST{
				Value: 22,
			},
		},
	},
	[]int{},
}

func main() {
	resultIn := testSubject.Bst.InOrderTraversal(testSubject.Array)
	testSubject.Array = []int{}
	resultPre := testSubject.Bst.PreOrderTraversal(testSubject.Array)
	testSubject.Array = []int{}
	resultPos := testSubject.Bst.PostOrderTraversal(testSubject.Array)
	fmt.Printf("In order: %v\nPre order: %v\nPost order: %v\n", resultIn, resultPre, resultPos)
}

func (tree *BST) InOrderTraversal(array []int) []int {
	if tree.Left != nil {
		array = tree.Left.InOrderTraversal(array)
	}
	array = append(array, tree.Value)
	if tree.Right != nil {
		array = tree.Right.InOrderTraversal(array)
	}
	return array
}

func (tree *BST) PreOrderTraversal(array []int) []int {
	array = append(array, tree.Value)
	if tree.Left != nil {
		array = tree.Left.PreOrderTraversal(array)
	}
	if tree.Right != nil {
		array = tree.Right.PreOrderTraversal(array)
	}
	return array
}

func (tree *BST) PostOrderTraversal(array []int) []int {
	if tree.Left != nil {
		array = tree.Left.PostOrderTraversal(array)
	}
	if tree.Right != nil {
		array = tree.Right.PostOrderTraversal(array)
	}
	array = append(array, tree.Value)
	return array
}
