package main

import "fmt"

type testCase struct {
	array    []int
	sequence []int
	expected bool
}

var testCases []testCase = []testCase{
	{
		array:    []int{5, 1, 22, 25, 6, -1, 8, 10},
		sequence: []int{1, 6, -1, 10},
	},
	{
		array:    []int{5, 1, 22, 25, 6, -1, 8, 10},
		sequence: []int{5, 1, 22, 25, 6, -1, 8, 10},
	},
	{
		array:    []int{5, 1, 22, 25, 6, -1, 8, 10},
		sequence: []int{5, 1, 22, 6, -1, 8, 10},
	},
	{
		array:    []int{5, 1, 22, 25, 6, -1, 8, 10},
		sequence: []int{22, 25, 6},
	},
	{
		array:    []int{5, 1, 22, 25, 6, -1, 8, 10},
		sequence: []int{1, 6, 10},
	},
	{
		array:    []int{5, 1, 22, 25, 6, -1, 8, 10},
		sequence: []int{5, 1, 22, 10},
	},
	{
		array:    []int{5, 1, 22, 25, 6, -1, 8, 10},
		sequence: []int{5, -1, 8, 10},
	},
	{
		array:    []int{5, 1, 22, 25, 6, -1, 8, 10},
		sequence: []int{25},
	},
	{
		array:    []int{1, 1, 1, 1, 1},
		sequence: []int{1, 1, 1},
	},
	{
		array:    []int{5, 1, 22, 25, 6, -1, 8, 10},
		sequence: []int{5, 1, 22, 25, 6, -1, 8, 10, 12},
	},
	{
		array:    []int{5, 1, 22, 25, 6, -1, 8, 10},
		sequence: []int{4, 5, 1, 22, 25, 6, -1, 8, 10},
	},
	{
		array:    []int{5, 1, 22, 25, 6, -1, 8, 10},
		sequence: []int{5, 1, 22, 23, 6, -1, 8, 10},
	},
	{
		array:    []int{5, 1, 22, 25, 6, -1, 8, 10},
		sequence: []int{5, 1, 22, 22, 25, 6, -1, 8, 10},
	},
	{
		array:    []int{5, 1, 22, 25, 6, -1, 8, 10},
		sequence: []int{5, 1, 22, 22, 6, -1, 8, 10},
	},
	{
		array:    []int{5, 1, 22, 25, 6, -1, 8, 10},
		sequence: []int{1, 6, -1, -1},
	},
	{
		array:    []int{5, 1, 22, 25, 6, -1, 8, 10},
		sequence: []int{1, 6, -1, -1, 10},
	},
	{
		array:    []int{5, 1, 22, 25, 6, -1, 8, 10},
		sequence: []int{1, 6, -1, -2},
	},
	{
		array:    []int{5, 1, 22, 25, 6, -1, 8, 10},
		sequence: []int{26},
	},
	{
		array:    []int{5, 1, 22, 25, 6, -1, 8, 10},
		sequence: []int{5, 1, 25, 22, 6, -1, 8, 10},
	},
	{
		array:    []int{5, 1, 22, 25, 6, -1, 8, 10},
		sequence: []int{5, 26, 22, 8},
	},
	{
		array:    []int{1, 1, 6, 1},
		sequence: []int{1, 1, 1, 6},
	},
	{
		array:    []int{5, 1, 22, 25, 6, -1, 8, 10},
		sequence: []int{1, 6, -1, 10, 11, 11, 11, 11},
	},
	{
		array:    []int{5, 1, 22, 25, 6, -1, 8, 10},
		sequence: []int{5, 1, 22, 25, 6, -1, 8, 10, 10},
	},
	{
		array:    []int{5, 1, 22, 25, 6, -1, 8, 10},
		sequence: []int{1, 6, -1, 5},
	},
}

func main() {
	for _, test := range testCases {
		result := IsValidSubsequence2(test.array, test.sequence)
		fmt.Println(result)
	}
}

func IsValidSubsequence1(array []int, sequence []int) bool {
	// Write your code here.
	if len(sequence) == 1 && Contains(array, sequence[0]) {
		return true
	}
	currentItem := sequence[0]
	arrIdx := Index(array, currentItem)
	if arrIdx < 0 {
		return false
	} else {
		return IsValidSubsequence1(array[arrIdx+1:], sequence[1:])
	}
}

func IsValidSubsequence2(array []int, sequence []int) bool {
	var seqIdx int
	for _, arrItem := range array {
		if arrItem == sequence[seqIdx] {
			seqIdx++
		}
		if seqIdx == len(sequence) {
			return true
		}
	}
	return false
}

func Contains(array []int, item int) bool {
	index := Index(array, item)
	return index >= 0
}

func Index(array []int, item int) int {
	for idx, arrayItem := range array {
		if arrayItem == item {
			return idx
		}
	}
	return -1
}
