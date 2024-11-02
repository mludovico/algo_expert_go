package main

import "fmt"

// This is the struct of the input root. Do not edit it.
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

var node6 *BinaryTree = &BinaryTree{Left: nil, Right: nil, Value: 6}
var node7 *BinaryTree = &BinaryTree{Left: nil, Right: nil, Value: 7}
var node8 *BinaryTree = &BinaryTree{Left: nil, Right: nil, Value: 8}
var node9 *BinaryTree = &BinaryTree{Left: nil, Right: nil, Value: 9}
var node10 *BinaryTree = &BinaryTree{Left: nil, Right: nil, Value: 10}
var node4 *BinaryTree = &BinaryTree{Left: node8, Right: node9, Value: 4}
var node2 *BinaryTree = &BinaryTree{Left: node4, Right: node5, Value: 2}
var node1 *BinaryTree = &BinaryTree{Left: node2, Right: node3, Value: 1}
var node3 *BinaryTree = &BinaryTree{Left: node6, Right: node7, Value: 3}
var node5 *BinaryTree = &BinaryTree{Left: node10, Right: nil, Value: 5}

func main() {
	result := BranchSums(node1)
	fmt.Println(result)
}

func BranchSums(root *BinaryTree) []int {
	// Write your code here.
	results = []int{}
	return sumSubTree(root, 0)
}

var results []int

func sumSubTree(tree *BinaryTree, sum int) []int {
	sum += tree.Value
	if tree.Left == nil && tree.Right == nil {
		results = append(results, sum)
		return results
	}
	if tree.Left != nil {
		sumSubTree(tree.Left, sum)
	}
	if tree.Right != nil {
		sumSubTree(tree.Right, sum)
	}
	return results
}
