package main

import (
	"fmt"
	"sort"
)

var testSubject = []int{
	2, 5, -4, -7, 12, 100, -25,
}

func main() {
	result := sweetSavory(testSubject, 7)
	fmt.Println(result)
}

func sweetSavory(dishes []int, target int) []int {
	sort.Ints(dishes)
	var left, runningBalance int
	right := len(dishes) - 1
	result := []int{0, 0}
	for left < right && dishes[left] < 0 && dishes[right] > 0 {
		balance := dishes[left] + dishes[right]
		if balance < target {
			if balance > runningBalance {
				runningBalance = balance
				result = []int{dishes[left], dishes[right]}
			}
			left++
		} else if balance > target {
			right--
		} else {
			return []int{left, right}
		}
	}
	return result
}
