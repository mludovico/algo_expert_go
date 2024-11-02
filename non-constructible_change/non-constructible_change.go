package main

import (
	"fmt"
	"sort"
)

var testCases = [][]int{
	{5, 7, 1, 1, 2, 3, 22},
	{1, 1, 1, 1, 1},
	{1, 5, 1, 1, 1, 10, 15, 20, 100},
	{6, 4, 5, 1, 1, 8, 9},
	{},
	{87},
	{5, 6, 1, 1, 2, 3, 4, 9},
	{5, 6, 1, 1, 2, 3, 43},
	{1, 1},
	{2},
	{1},
	{109, 2000, 8765, 19, 18, 17, 16, 8, 1, 1, 2, 4},
	{1, 2, 3, 4, 5, 6, 7},
}

func main() {
	for _, testCase := range testCases {
		nonConstructibleChange := NonConstructibleChange(testCase)
		fmt.Println(nonConstructibleChange)
	}
}

func NonConstructibleChange(coins []int) int {
	// Write your code here.
	sort.Sort(sort.IntSlice(coins))
	constructedAmount := 0
	for _, coinValue := range coins {
		if coinValue > constructedAmount+1 {
			return constructedAmount + 1
		} else {
			constructedAmount += coinValue
		}
	}
	return constructedAmount + 1
}
