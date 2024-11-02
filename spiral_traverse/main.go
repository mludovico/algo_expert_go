package main

import (
	"fmt"
	"spiral_traverse/spiral_traverse"
)

var testArray = [][]int{
	{1, 2, 3},
	{8, 9, 4},
	{7, 6, 5},
}

func main() {
	result := spiral_traverse.SpiralTraverse(testArray)
	fmt.Println(result)
}
