package main

import (
	"fmt"
	"sort"
)

var redShirtHeights = []int{5, 8, 1, 3, 4}
var blueShirtHeights = []int{6, 9, 2, 4, 5}

func main() {
	result := ClassPhoto(redShirtHeights, blueShirtHeights)
	fmt.Println(result)
}

func ClassPhoto(redShirtHeights []int, blueShirtHeights []int) bool {
	sort.Ints(redShirtHeights)
	sort.Ints(blueShirtHeights)
	tallerTeam := getTallerTeamName(redShirtHeights[0], blueShirtHeights[0])
	if tallerTeam == "" {
		return false
	}
	for i, h := range redShirtHeights {
		currentTallerTeam := getTallerTeamName(h, blueShirtHeights[i])
		if currentTallerTeam != tallerTeam {
			return false
		}
	}
	return true
}

func getTallerTeamName(redShirtHeight int, blueShirtHeight int) string {
	if redShirtHeight < blueShirtHeight {
		return "BLUE"
	} else if redShirtHeight > blueShirtHeight {
		return "RED"
	}
	return ""
}
