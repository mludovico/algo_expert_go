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
var node4 *BinaryTree = &BinaryTree{Left: node8, Right: node9, Value: 4}
var node5 *BinaryTree = &BinaryTree{Left: nil, Right: nil, Value: 5}
var node2 *BinaryTree = &BinaryTree{Left: node4, Right: node5, Value: 2}
var node3 *BinaryTree = &BinaryTree{Left: node6, Right: node7, Value: 3}
var node1 *BinaryTree = &BinaryTree{Left: node2, Right: node3, Value: 1}

func main() {
	result := NodeDepths(node1)
	fmt.Println(result)
}

var depthSum int
var depth int

func NodeDepths(root *BinaryTree) int {
	depthSum = 0
	depth = 0
	return getDepths(root)
}

func getDepths(tree *BinaryTree) int {
	if tree.Left == nil && tree.Right == nil {
		depthSum += depth
		return depthSum
	}
	if tree.Left != nil {
		depth++
		getDepths(tree.Left)
		depth--
	}
	if tree.Right != nil {
		depth++
		getDepths(tree.Right)
		depth--
	}
	depthSum += depth
	return depthSum
}
