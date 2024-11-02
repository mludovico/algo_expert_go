package main

import (
	"fmt"
	"math"
)

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

var node502 = &BST{Left: node204, Right: node55000, Value: 502}
var node55000 = &BST{Left: node1001, Right: nil, Value: 55000}
var node1001 = &BST{Left: nil, Right: node4500, Value: 1001}
var node4500 = &BST{Left: nil, Right: nil, Value: 4500}
var node204 = &BST{Left: node203, Right: node205, Value: 204}
var node205 = &BST{Left: nil, Right: node207, Value: 205}
var node207 = &BST{Left: node206, Right: node208, Value: 207}
var node208 = &BST{Left: nil, Right: nil, Value: 208}
var node206 = &BST{Left: nil, Right: nil, Value: 206}
var node203 = &BST{Left: nil, Right: nil, Value: 203}
var node5 = &BST{Left: node2, Right: node15, Value: 5}
var node15 = &BST{Left: node5_2, Right: node22, Value: 15}
var node22 = &BST{Left: nil, Right: node57, Value: 22}
var node57 = &BST{Left: nil, Right: node60, Value: 57}
var node60 = &BST{Left: nil, Right: nil, Value: 60}
var node5_2 = &BST{Left: nil, Right: nil, Value: 5}
var node2 = &BST{Left: node1, Right: node3, Value: 2}
var node3 = &BST{Left: nil, Right: nil, Value: 3}
var node1 = &BST{Left: nodeMinus51, Right: node1_2, Value: 1}
var node1_2 = &BST{Left: nil, Right: node1_3, Value: 1}
var node1_3 = &BST{Left: nil, Right: node1_4, Value: 1}
var node1_4 = &BST{Left: nil, Right: node1_5, Value: 1}
var node1_5 = &BST{Left: nil, Right: nil, Value: 1}
var nodeMinus51 = &BST{Left: nodeMinus403, Right: nil, Value: -51}
var nodeMinus403 = &BST{Left: nil, Right: nil, Value: -403}
var node100 = &BST{Left: node5, Right: node502, Value: 100}

var testCase, targetCase = node100, 208

func main() {
	closestResult := testCase.FindClosestValue2(targetCase)
	fmt.Println(closestResult)
}

var closestInt int

func (tree *BST) FindClosestValue1(target int) int {
	// Write your code here.
	if tree == nil {
		return closestInt
	}
	currentValue := tree.Value
	if math.Abs(float64(currentValue-target)) < math.Abs(float64(closestInt-target)) {
		closestInt = currentValue
	}
	left := tree.Left
	right := tree.Right
	if left == nil && target < currentValue {
		return closestInt
	}
	if right == nil && target > currentValue {
		return closestInt
	}
	if target < currentValue {
		return left.FindClosestValue1(target)
	} else {
		return right.FindClosestValue1(target)
	}
}

func (tree *BST) FindClosestValue2(target int) int {
	// Write your code here.
	currentNode := tree
	for currentNode != nil {
		currentValue := currentNode.Value
		if math.Abs(float64(currentValue-target)) < math.Abs(float64(closestInt-target)) {
			closestInt = currentValue
		}
		left := currentNode.Left
		right := currentNode.Right
		if left == nil && target < currentValue {
			return closestInt
		}
		if right == nil && target > currentValue {
			return closestInt
		}
		if target < currentValue {
			currentNode = left
		} else {
			currentNode = right
		}
	}
	return closestInt
}
