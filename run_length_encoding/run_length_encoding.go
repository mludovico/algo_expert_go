package main

import (
	"fmt"
)

func main() {
	rlEncoded := RunLengthEncoding("........______=========AAAA   AAABBBB   BBB")
	fmt.Println(rlEncoded)
}

func RunLengthEncoding(str string) string {
	var currentChar byte = str[0]
	var currentCount int = 0
	var encSlice []byte
	for i := 0; i < len(str); i++ {
		if str[i] == byte(currentChar) {
			currentCount++
			if currentCount == 9 {
				encSlice = append(encSlice, byte(currentCount), currentChar)
				currentCount = 0
			}
		} else {
			if currentCount > 0 {
				encSlice = append(encSlice, byte(currentCount), currentChar)
			}
			currentChar = str[i]
			currentCount = 1
		}
	}
	encSlice = append(encSlice, byte(currentCount), currentChar)
	return encodeBlock(encSlice)
}

func encodeBlock(encSlice []byte) string {
	var encStr string
	fmt.Println(encSlice)
	for _, val := range encSlice {
		if val > 9 {
			encStr = fmt.Sprintf("%s%v", encStr, string(val))
		} else {
			encStr = fmt.Sprintf("%s%v", encStr, val)
		}
	}
	return encStr
}
