package main

import (
	"fmt"
)

var testArray = []int{
	2, 1, 5, 3, 3, 2, 4,
}

func main() {
	result := FirstDuplicateValue2(testArray)
	fmt.Println(result)
}

func FirstDuplicateValue1(array []int) int {
	verificationMap := map[int]bool{}
	for _, val := range array {
		if verificationMap[val] {
			return val
		} else {
			verificationMap[val] = true
		}
	}
	return -1
}

func FirstDuplicateValue2(array []int) int {
	for _, val := range array {
		idx := Abs(val)
		if array[idx-1] < 0 {
			return int(idx)
		}
		array[idx-1] *= -1
	}
	return -1
}

func Abs(val int) int {
	if val < 0 {
		val *= -1
	}
	return val
}
