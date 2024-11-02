package main

import "fmt"

var testArray = []int{
	8, 5, 2, 9, 5, 6, 3,
}

func main() {
	sortedArray := SelectionSort(testArray)
	fmt.Println(sortedArray)
}

func SelectionSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		smallestIndex := i
		for j := smallestIndex; j < len(array); j++ {
			if array[j] < array[smallestIndex] {
				smallestIndex = j
			}
		}
		if array[smallestIndex] < array[i] {
			array[i], array[smallestIndex] = array[smallestIndex], array[i]
		}
	}
	return array
}
