package main

import (
	"fmt"
)

var array = []int{
	1, 5, 23, 111,
}

var target = 35

func main() {
	index := BinarySearch2(array, target)
	fmt.Println(index)
}

func BinarySearch1(array []int, target int) int {
	return binarySearch(array, target, 0)
}

func binarySearch(array []int, target, initialIndex int) int {
	if len(array) < 2 && array[0] != target {
		return -1
	}
	middleIndex := len(array) / 2
	middleItem := array[middleIndex]
	if target > middleItem {
		return binarySearch(array[middleIndex:], target, initialIndex+middleIndex)
	}
	if target < middleItem {
		return binarySearch(array[:middleIndex], target, initialIndex)
	}
	return middleIndex + initialIndex

}

func BinarySearch2(array []int, target int) int {
	leftPointer := 0
	rightPointer := len(array) - 1
	middlePointer := len(array) / 2
	for array[middlePointer] != target && leftPointer < rightPointer {
		if target < array[middlePointer] {
			rightPointer = middlePointer - 1
			middlePointer = (leftPointer + rightPointer) / 2
		}
		if target > array[middlePointer] {
			leftPointer = middlePointer + 1
			middlePointer = (leftPointer + rightPointer) / 2
		}
	}
	if array[middlePointer] == target {
		return middlePointer
	} else {
		return -1
	}

}
