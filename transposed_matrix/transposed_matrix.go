package main

import "fmt"

func main() {
	tm := TransposedMatrix([][]int{{1, 2}})
	fmt.Println(tm)
}

func TransposedMatrix(matrix [][]int) [][]int {
	// Write your code here.
	var transposed [][]int = make([][]int, len(matrix[0]))
	for i := range transposed {
		transposed[i] = make([]int, len(matrix))
	}
	for i, row := range matrix {
		for j, val := range row {
			transposed[j][i] = val
		}
	}
	return transposed
}
