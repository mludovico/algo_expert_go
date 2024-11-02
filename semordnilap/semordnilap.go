package main

import (
	"fmt"
)

var testArray = []string{
	"abcde", "bcd", "dcb", "edcba", "aaa",
}

func main() {
	result := Semordnilap(testArray)
	fmt.Println(result)
}

func Semordnilap(words []string) [][]string {
	semordnilapMap := map[string]bool{}
	result := [][]string{}
	for _, str := range words {
		reversedStr := ""
		for i := len(str) - 1; i >= 0; i-- {
			reversedStr += string(str[i])
		}
		if !semordnilapMap[str] {
			semordnilapMap[reversedStr] = true
		} else {
			result = append(result, []string{str, reversedStr})
		}
	}
	return result
}
