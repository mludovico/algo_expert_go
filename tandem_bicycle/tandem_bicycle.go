package main

import (
	"sort"
)

var redShirtSpeeds = []int{5, 5, 3, 9, 2}
var blueShirtSpeeds = []int{3, 6, 7, 2, 1}
var fastest = true

func main() {
	result := TandemBicycle(redShirtSpeeds, blueShirtSpeeds, fastest)
	println(result)
}

func TandemBicycle(redShirtSpeeds []int, blueShirtSpeeds []int, fastest bool) int {
	sort.Ints(redShirtSpeeds)
	if fastest {
		sort.Sort(sort.Reverse(sort.IntSlice(blueShirtSpeeds)))
	} else {
		sort.Ints(blueShirtSpeeds)
	}
	totalSpeed := 0
	for i, s := range redShirtSpeeds {
		if s > blueShirtSpeeds[i] {
			totalSpeed += s
		} else {
			totalSpeed += blueShirtSpeeds[i]
		}
	}
	return totalSpeed
}
