package main

import (
	"fmt"
)

var testSubject = []int{-5, -5, 2, 3, -2}

func main() {
	result := zeroSumArray2(testSubject)
	fmt.Println(result)
}

func zeroSumArray1(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		if sum == 0 {
			return true
		}
		for j := i + 1; j < len(nums); j++ {
			if sum += nums[j]; sum == 0 {
				return true
			}
		}
	}
	return false
}

func zeroSumArray2(nums []int) bool {
	sums := map[int]bool{}
	var runningSum int
	for i := 0; i < len(nums); i++ {
		runningSum += nums[i]
		if runningSum == 0 {
			return true
		}
		if !sums[runningSum] {
			sums[runningSum] = true
		} else {
			return true
		}
	}
	return false
}
