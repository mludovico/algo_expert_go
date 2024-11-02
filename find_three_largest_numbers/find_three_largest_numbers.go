package main

import (
	"fmt"
	"math"
)

var testArray = []int{
	-1, -2, -3, -7, -17, -27, -18, -541, -8, -7, 7,
}

func main() {
	largesThree := FindThreeLargestNumbers(testArray)
	fmt.Println(largesThree)
}

func FindThreeLargestNumbers(array []int) []int {
	var largest = math.MinInt
	var middle = largest
	var smallest = largest

	for _, val := range array {
		if val > largest {
			smallest = middle
			middle = largest
			largest = val
			continue
		}
		if val > middle {
			smallest = middle
			middle = val
			continue
		}
		if val > smallest {
			smallest = val
		}
	}

	return []int{
		smallest,
		middle,
		largest,
	}
}
