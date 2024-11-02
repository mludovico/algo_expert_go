package main

import (
	"fmt"
)

var testArray = []int{
	5, 4, 3, 2, 1, 2, 10, 12, -3, 5, 6, 7, 10,
}

func main() {
	result := LongestPeak(testArray)
	fmt.Println(result)
}

func LongestPeak(array []int) int {
	var longestPeakLength int
	for i := 1; i < len(array)-1; {
		if array[i-1] < array[i] && array[i] > array[i+1] {
			startIdx := i - 2
			for startIdx >= 0 && array[startIdx] < array[startIdx+1] {
				startIdx--
			}
			endIdx := i + 2
			for endIdx < len(array) && array[endIdx] < array[endIdx-1] {
				endIdx++
			}
			length := endIdx - startIdx - 1
			if length > longestPeakLength {
				longestPeakLength = length
			}
			i = endIdx
		} else {
			i++
		}
	}
	return longestPeakLength
}
