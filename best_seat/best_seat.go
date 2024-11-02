package main

import (
	"fmt"
)

var testArray = []int{
	1,
}

func main() {
	result := BestSeat(testArray)
	fmt.Println(result)
}

func BestSeat(seats []int) int {
	var emptySpaces [][]int = make([][]int, len(seats))
	spacesIdx := 0
	for i := 1; i < len(seats)-1; i++ {
		if seats[i] == 0 {
			emptySpaces[spacesIdx] = append(emptySpaces[spacesIdx], i)

		} else {
			spacesIdx++
		}
	}
	biggestSpace := BiggestSlice(emptySpaces)
	if len(biggestSpace) == 0 {
		return -1
	}
	idx := (len(biggestSpace) - 1) / 2
	return biggestSpace[idx]
}

func BiggestSlice(array [][]int) []int {
	biggestSlice := []int{}
	for _, slc := range array {
		if len(slc) > len(biggestSlice) {
			biggestSlice = slc
		}
	}
	return biggestSlice
}
