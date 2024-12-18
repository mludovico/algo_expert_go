package main

import "fmt"

var testArray = []int{
	1, 2, 2, 7, 2, 2,
}

func main() {
	result := MajorityElement2(testArray)
	fmt.Println(result)
}

func MajorityElement1(array []int) int {
	counts := map[int]int{}
	for _, val := range array {
		counts[val]++
	}
	for i, count := range counts {
		if count > len(array)/2 {
			return i
		}
	}
	return -1
}

func MajorityElement2(array []int) int {
	var candidate, count int
	for _, val := range array {
		if count == 0 {
			candidate = val
		}
		if val == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
