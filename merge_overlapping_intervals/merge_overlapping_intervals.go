package main

import (
	"fmt"
	"math"
)

var testArray = [][]int{
	{1, 2},
	{3, 5},
	{4, 7},
	{6, 8},
	{9, 10},
}

func main() {
	result := MergeOverlappingIntervals(testArray)
	fmt.Println(result)
}

func MergeOverlappingIntervals(intervals [][]int) [][]int {

	for i := 0; i < len(intervals)-1; {
		pair1 := intervals[i]
		pair2 := intervals[i+1]
		if pair1[1] >= pair2[0] {
			newPair := []int{pair1[0], Max(pair1[1], pair2[1])}
			newEnding := append([][]int{newPair}, intervals[i+2:]...)
			intervals = append(intervals[:i], newEnding...)
		} else {
			i++
		}
	}
	return intervals
}

func Sort(array [][]int) [][]int {
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i][0] > array[j][0] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	return array
}

func Max(vals ...int) int {
	max := int(math.Inf(-1))
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}
