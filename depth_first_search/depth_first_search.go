package main

import "fmt"

// Do not edit the class below except
// for the depthFirstSearch method.
// Feel free to add new properties
// and methods to the class.
type Node struct {
	Name     string
	Children []*Node
}

var nodeI = &Node{Name: "I"}
var nodeJ = &Node{Name: "J"}
var nodeK = &Node{Name: "K"}
var nodeE = &Node{Name: "E"}
var nodeF = &Node{Name: "F", Children: []*Node{nodeI, nodeJ}}
var nodeG = &Node{Name: "G", Children: []*Node{nodeK}}
var nodeH = &Node{Name: "H"}
var nodeB = &Node{Name: "B", Children: []*Node{nodeE, nodeF}}
var nodeC = &Node{Name: "C"}
var nodeD = &Node{Name: "D", Children: []*Node{nodeG, nodeH}}
var nodeA = &Node{Name: "A", Children: []*Node{nodeB, nodeC, nodeD}}

func main() {
	var array []string = []string{}
	result := nodeA.DepthFirstSearch(array)
	fmt.Println(result)
}

func (n *Node) DepthFirstSearch(array []string) []string {
	// Write your code here.
	return searchAndFillReference(n, &array)
}

func searchAndFillReference(node *Node, array *[]string) []string {
	// Write your code here.
	*array = append(*array, node.Name)
	for _, node := range node.Children {
		searchAndFillReference(node, array)
	}
	return *array

}
