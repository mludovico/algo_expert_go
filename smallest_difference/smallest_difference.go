package main

import (
	"fmt"
	"math"
	"sort"
)

var testArray1 = []int{
	10, 0, 20, 25, 2000,
}
var testArray2 = []int{
	1005, 1006, 1014, 1032, 1031,
}

func main() {
	result := SmallestDifference(testArray1, testArray2)
	fmt.Println(result)
}

func SmallestDifference(array1, array2 []int) []int {
	sort.Ints(array1)
	sort.Ints(array2)
	pointer1, pointer2 := 0, 0
	difference := math.Inf(1)
	candidates := []int{}
	for pointer1 < len(array1) && pointer2 < len(array2) {
		currentDifference := math.Inf(1)
		if array1[pointer1] < array2[pointer2] {
			currentDifference = float64(array2[pointer2] - array1[pointer1])
		} else if array2[pointer2] < array1[pointer1] {
			currentDifference = float64(array1[pointer1] - array2[pointer2])
		} else {
			return []int{array1[pointer1], array2[pointer2]}
		}
		if currentDifference < difference {
			difference = currentDifference
			candidates = []int{array1[pointer1], array2[pointer2]}
		}
		if array1[pointer1] < array2[pointer2] {
			pointer1++
		} else if array2[pointer2] < array1[pointer1] {
			pointer2++
		}
	}
	return candidates
}
