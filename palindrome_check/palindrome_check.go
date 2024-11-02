package main

import "fmt"

var testString = "abcdcbaa"

func main() {
	isPalindrome := PalindromeCheck(testString)
	fmt.Println(isPalindrome)
}

func PalindromeCheck(s string) bool {
	var reversedString string
	for i := len(s) - 1; i >= 0; i-- {
		reversedString = fmt.Sprintf("%s%c", reversedString, s[i])
	}
	return reversedString == s
}
