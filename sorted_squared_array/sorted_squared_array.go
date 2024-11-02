package main

import (
	"fmt"
	"sort"
)

var testCases [][]int = [][]int{
	{1, 2, 3, 5, 6, 8, 9},
	{1},
	{1, 2},
	{1, 2, 3, 4, 5},
	{0},
	{10},
	{-1},
	{-2, -1},
	{-5, -4, -3, -2, -1},
	{-10},
	{-10, -5, 0, 5, 10},
	{-7, -3, 1, 9, 22, 30},
	{-50, -13, -2, -1, 0, 0, 1, 1, 2, 3, 19, 20},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{-1, -1, 2, 3, 3, 3, 4},
	{-3, -2, -1},
	{-3, -2, -1},
}

func main() {
	for _, test := range testCases {
		result := SortedSquaredArray1(test)
		fmt.Println(result)
	}
}

func SortedSquaredArray1(array []int) []int {
	squaredArray := make([]int, len(array))
	for i, num := range array {
		squaredArray[i] = num * num
	}
	sort.Ints(squaredArray)
	return squaredArray
}

func SortedSquaredArray2(array []int) []int {
	squaredArray := make([]int, len(array))
	left := 0
	right := len(array) - 1
	for i := len(array) - 1; i >= 0; i-- {
		leftValue := array[left]
		rightValue := array[right]
		if abs(leftValue) > abs(rightValue) {
			squaredArray[i] = leftValue * leftValue
			left++
		} else {
			squaredArray[i] = rightValue * rightValue
			right--
		}
	}
	return squaredArray
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
