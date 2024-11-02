package main

import (
	"fmt"
)

var testArray = []int{
	12, 3, 1, 2, -6, 5, -8, 6,
}

func main() {
	result := ThreeNumberSum(testArray, 0)
	fmt.Println(result)
}

func ThreeNumberSum(array []int, target int) [][]int {
	if len(array) < 3 {
		return [][]int{}
	}

	sortedArray := Sort(array)
	result := [][]int{}

	for i, val := range sortedArray {
		leftPointer := i + 1
		rightPointer := len(sortedArray) - 1
		for leftPointer < len(sortedArray)-1 && rightPointer > leftPointer {
			sum := val + sortedArray[leftPointer] + sortedArray[rightPointer]
			if sum == target {
				result = append(result, []int{val, sortedArray[leftPointer], sortedArray[rightPointer]})
				leftPointer++
				rightPointer--
			} else if sum < target {
				leftPointer++
				continue
			} else {
				rightPointer--
			}
		}
	}
	return result
}

func Sort(s []int) []int {
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	return s
}
