package main

import (
	"fmt"
)

var testArray = []int{
	2, 1, 2, 2, 2, 3, 4, 2,
}

func main() {
	result := MoveElementToEnd(testArray, 2)
	fmt.Println(result)
}

func MoveElementToEnd(array []int, toMove int) []int {
	p1 := 0
	p2 := len(array) - 1
	for p1 < p2 {
		if array[p1] == toMove && array[p2] != toMove {
			array[p1], array[p2] = array[p2], array[p1]
			p1++
			p2--
			continue
		}
		if array[p1] != toMove {
			p1++
			continue
		}
		if array[p2] == toMove {
			p2--
		}
	}
	return array
}

// |p1|p2|            |  |
// |= |= |p2--        |  |
// |!=|!=|p1++        |  |
// |= |!=|mv,p1++,p2--|  |
// |!=|==|p1++,p2--   |  |
