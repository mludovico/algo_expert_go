package main

import (
	"fmt"
	"time"
)

var testArray = []int{
	5, 1, 4, 2,
}

func main() {
	startTime := time.Now().UnixMicro()
	result := ArrayOfProductsMultipleLoops(testArray)
	endTime := time.Now().UnixMicro()
	fmt.Printf("Total: %vms\n", endTime-startTime)
	fmt.Println(result)
}

func ArrayOfProductsMultipleLoops(array []int) []int {
	products := []int{}
	for i := 0; i < len(array); i++ {
		product := 1
		for j := i - 1; j >= 0; j-- {
			product *= array[j]
		}
		for j := i + 1; j < len(array); j++ {
			product *= array[j]
		}
		products = append(products, product)
	}

	return products

}

func ArrayOfProducts2Loops(array []int) []int {
	rightProducts := make([]int, len(array))
	leftProducts := make([]int, len(array))
	for i := 0; i < len(array); i++ {
		leftProducts[i] = 1
		if i != 0 {
			leftProducts[i] = leftProducts[i-1] * array[i-1]
		}
		rightIndex := len(array) - i - 1
		rightProducts[rightIndex] = 1
		if rightIndex != len(array)-1 {
			rightProducts[rightIndex] = rightProducts[rightIndex+1] * array[rightIndex+1]
		}
	}
	for i := 0; i < len(array); i++ {
		array[i] = leftProducts[i] * rightProducts[i]
	}
	return array
}
