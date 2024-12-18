package main

import (
	"fmt"
	"slices"
	"sort"
)

var testSubject = []int{1, 4, 3}

func main() {
	result := missingNumbers4(testSubject)
	fmt.Println(result)
}

func missingNumbers1(nums []int) []int {
	missingNums := []int{}
	for i := 1; i <= len(nums)+2; i++ {
		if !slices.Contains(nums, i) {
			missingNums = append(missingNums, i)
		}
	}
	return missingNums
}

func missingNumbers2(nums []int) []int {
	existingSet := map[int]bool{}
	for _, val := range nums {
		existingSet[val] = true
	}

	var missingSet []int
	for i := 1; i <= len(nums)+2; i++ {
		if !existingSet[i] {
			missingSet = append(missingSet, i)
		}
	}
	return missingSet
}

func missingNumbers3(nums []int) []int {
	var sum int
	for i := 1; i <= len(nums)+2; i++ {
		sum += i
	}
	for _, val := range nums {
		sum -= val
	}
	average := sum / 2

	var found1stHalf, found2ndHalf int
	for _, val := range nums {
		if val <= average {
			found1stHalf += val
		} else {
			found2ndHalf += val
		}
	}

	var expected1stHalf, expected2ndHalf int
	for i := 1; i <= average; i++ {
		expected1stHalf += i
	}
	for i := average + 1; i <= len(nums)+2; i++ {
		expected2ndHalf += i
	}
	return []int{expected1stHalf - found1stHalf, expected2ndHalf - found2ndHalf}
}

func missingNumbers4(nums []int) []int {
	var xor = 0
	for i := 0; i <= len(nums)+2; i++ {
		xor ^= i

		if i < len(nums) {
			xor ^= nums[i]
		}
	}

	var solution = []int{0, 0}
	setBit := xor & -xor
	for i := 0; i <= len(nums)+2; i++ {
		if i&setBit == 0 {
			solution[0] ^= i
		} else {
			solution[1] ^= i
		}

		if i < len(nums) {
			if nums[i]&setBit == 0 {
				solution[0] ^= nums[i]
			} else {
				solution[1] ^= nums[i]
			}
		}
	}

	sort.Ints(solution)
	return solution
}
