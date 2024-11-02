package main

import (
	"fmt"
	"sort"
)

var queries = []int{3, 2, 1, 2, 6}

func main() {
	result := MinimumWaitingTime(queries)
	fmt.Println(result)
}

func MinimumWaitingTime(queries []int) int {
	var s int
	Sort(&queries)
	stop := len(queries) - 1
	for i := 0; i < stop; i++ {
		s += queries[i] * (stop - i)
	}
	return s
}

func Sort(s *[]int) {
	sort.Ints()
	for i := 1; i < len(*s); i++ {
		for j := i; j > 0; j-- {
			if (*s)[j] < (*s)[j-1] {
				(*s)[j], (*s)[j-1] = (*s)[j-1], (*s)[j]
			}
		}
	}
}
