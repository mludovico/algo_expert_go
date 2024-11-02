package main

import (
	"fmt"
)

var testArray = []string{
	"abcde", "aa", "foobar", "foobaz", "and this is a string", "aaaaaaaa", "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeea",
}

func main() {
	common := CommonCharacters(testArray)
	fmt.Println(common)
}

func CommonCharacters(strings []string) []string {
	charMap := map[string]int{}

	for _, str := range strings {
		inThisString := map[string]bool{}
		for _, char := range str {
			strChar := string(char)
			if !inThisString[strChar] {
				charMap[strChar] += 1
				inThisString[strChar] = true
			}
		}
	}

	presentChars := []string{}
	for key, val := range charMap {
		if val == len(strings) {
			presentChars = append(presentChars, key)
		}
	}
	return presentChars
}
