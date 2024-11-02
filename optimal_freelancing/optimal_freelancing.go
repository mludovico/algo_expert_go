package main

import "fmt"

var jobs = []map[string]int{
	{
		"deadline": 8,
		"payment":  3,
	},
	{
		"deadline": 1,
		"payment":  1,
	},
	{
		"deadline": 1,
		"payment":  2,
	},
}

func main() {
	result := OptimalFrelancing(jobs)
	fmt.Println(result)
}

func OptimalFrelancing(jobs []map[string]int) int {
	SortByParameter(jobs, "payment", true)
	var profit int
	var timeline = make([]bool, 7)
	for _, job := range jobs {
		deadlineIndex := job["deadline"] - 1
		for i := deadlineIndex; i >= 0; i-- {
			if len(timeline) > i && !timeline[i] {
				timeline[i] = true
				profit += job["payment"]
				break
			}
		}
	}
	return profit
}

func SortByParameter(jobs []map[string]int, parameter string, desc bool) {
	for i := 0; i < len(jobs); i++ {
		for j := 0; j < len(jobs)-i-1; j++ {
			if desc {
				if jobs[j][parameter] < jobs[j+1][parameter] {
					jobs[j], jobs[j+1] = jobs[j+1], jobs[j]
				}
			} else {
				if jobs[j][parameter] > jobs[j+1][parameter] {
					jobs[j], jobs[j+1] = jobs[j+1], jobs[j]
				}
			}
		}
	}
}
