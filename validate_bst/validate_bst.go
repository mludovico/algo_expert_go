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

//{"id": "10", "left": "5", "right": "15", "value": 10},
//{"id": "15", "left": null, "right": "22", "value": 15},
//{"id": "22", "left": null, "right": null, "value": 22},
//{"id": "5", "left": "2", "right": "5-2", "value": 5},
//{"id": "5-2", "left": null, "right": null, "value": 5},
//{"id": "2", "left": "1", "right": null, "value": 2},
//{"id": "1", "left": "-5", "right": null, "value": 1},
//{"id": "-5", "left": "-15", "right": "-5-2", "value": -5},
//{"id": "-5-2", "left": null, "right": "-2", "value": -5},
//{"id": "-2", "left": null, "right": "-1", "value": -2},
//{"id": "-1", "left": null, "right": null, "value": -1},
//{"id": "-15", "left": "-22", "right": null, "value": -15},
//{"id": "-22", "left": null, "right": null, "value": -22}

//		               10
//					  /  \
//		             5    15
//		            / \     \
//		           2   5     22
//		          /
//		         1
//		        /
//		      -5
//		     /   \
//		  -15    -5
//		  /        \
//	 -22        -2
//	              \
//	              -1
var testSubject = BST{
	Value: 10,
	Left: &BST{
		Value: 5,
		Left: &BST{
			Value: 2,
			Left: &BST{
				Value: 1,
				Left: &BST{
					Value: -5,
					Left: &BST{
						Value: -15,
						Left: &BST{
							Value: -22,
						},
					},
					Right: &BST{
						Value: -5,
						Right: &BST{
							Value: -2,
							Right: &BST{
								Value: -1,
							},
						},
					},
				},
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
}

func main() {
	result := testSubject.validateBst()
	fmt.Println(result)
}

var negInfinity = int(math.Inf(-1))
var posInfinity = int(math.Inf(1))

func (tree *BST) validateBst() bool {
	return tree.validateWithLimits(negInfinity, posInfinity)
}

func (tree *BST) validateWithLimits(min, max int) bool {
	if tree.Value < min || tree.Value >= max {
		return false
	}
	left := true
	right := true
	if tree.Left != nil {
		left = tree.Left.validateWithLimits(min, tree.Value)
	}
	if tree.Right != nil {
		right = tree.Right.validateWithLimits(tree.Value, max)
	}
	return left && right
}
