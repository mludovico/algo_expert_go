package main

import "fmt"

func main() {
	for _, test := range tests {
		result := TwoNumberSum2(test.inputArray, test.inputTarget)
		fmt.Println(result)
	}
}

type twoNumberSumPair struct {
	inputArray  []int
	inputTarget int
	expectedSum []int
}

var tests []twoNumberSumPair = []twoNumberSumPair{
	{[]int{4, 6}, 10, []int{4, 6}},
	{[]int{4, 6, 1}, 5, []int{4, 1}},
	{[]int{4, 6, 1, -3}, 3, []int{6, -3}},
	{[]int{3, 5, -4, 8, 11, 1, -1, 6}, 10, []int{11, -1}},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 17, []int{8, 9}},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 15}, 18, []int{3, 15}},
	{[]int{-7, -5, -3, -1, 0, 1, 3, 5, 7}, -5, []int{0, -5}},
	{[]int{-21, 301, 12, 4, 65, 56, 210, 356, 9, -47}, 163, []int{-47, 210}},
	{[]int{3, 5, -4, 8, 11, 1, -1, 6}, 15, []int{}},
}

func TwoNumberSum1(array []int, target int) []int {
	nums := make(map[int]bool)
	for _, num := range array {
		if _, found := nums[target-num]; found {
			return []int{num, target - num}
		}
		nums[num] = true
	}
	return []int{}
}

func TwoNumberSum2(array []int, target int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i]+array[j] == target {
				return []int{array[i], array[j]}
			}
		}
	}
	return []int{}
}
