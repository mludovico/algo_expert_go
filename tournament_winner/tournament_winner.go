package main

import "fmt"

type testCase struct {
	competitions [][]string
	results      []int
}

var testCases testCase = testCase{
	competitions: [][]string{
		{"HTML", "C#"},
		{"C#", "Python"},
		{"Python", "HTML"},
	},
	results: []int{0, 0, 1},
}

func main() {
	winner := TournamentWinner(testCases.competitions, testCases.results)
	fmt.Println(winner)
}

func TournamentWinner(competitions [][]string, results []int) string {
	// Write your code here.
	scoreboard := map[string]int{}
	winner := ""
	maxValue := 0
	for i, competition := range competitions {
		winnerIdx := 1 - results[i]
		scoreboard[competition[winnerIdx]] += 3
		if scoreboard[competition[winnerIdx]] >= maxValue {
			maxValue = scoreboard[competition[winnerIdx]]
			winner = competition[winnerIdx]
		}
	}
	return winner
}
