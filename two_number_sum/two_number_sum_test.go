package twonumbersum_test

import (
	"testing"
	twonumbersum "two_number_sum"
)

type twoNumberSumPair struct {
	inputArray  []int
	inputTarget int
	expectedSum []int
}

var tests []twoNumberSumPair = []twoNumberSumPair{
	{[]int{4, 6}, 10, []int{4, 6}},
	{[]int{4, 6, 1}, 5, []int{4, 1}},
	{[]int{4, 6, 1, -3}, 3, []int{6, -3}},
	{[]int{3, 5, -4, 8, 11, 1, -1, 6}, 10, []int{11, -1}},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 17, []int{8, 9}},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 15}, 18, []int{3, 15}},
	{[]int{-7, -5, -3, -1, 0, 1, 3, 5, 7}, -5, []int{0, -5}},
	{[]int{-21, 301, 12, 4, 65, 56, 210, 356, 9, -47}, 163, []int{-47, 210}},
	{[]int{3, 5, -4, 8, 11, 1, -1, 6}, 15, []int{}},
}

func ExampleTwoNumberSum1(t *testing.T) {
	for _, test := range tests {
		result := twonumbersum.TwoNumberSum1(test.inputArray, test.inputTarget)
		if result[0] != test.expectedSum[0] || result[1] != test.expectedSum[1] {
			t.Errorf("Using input %v, got: %v, want: %v.", test.inputArray, result, test.expectedSum)
		}
	}

}

func ExampleTwoNumberSum2(t *testing.T) {
	for _, test := range tests {
		result := twonumbersum.TwoNumberSum2(test.inputArray, test.inputTarget)
		if result[0] != test.expectedSum[0] || result[1] != test.expectedSum[1] {
			t.Errorf("Using input %v, got: %v, want: %v.", test.inputArray, result, test.expectedSum)
		}
	}
}
