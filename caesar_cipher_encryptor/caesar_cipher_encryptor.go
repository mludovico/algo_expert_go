package main

import (
	"fmt"
	"strings"
)

func main() {
	encrypted := CaesarCipherEncriptor("xyz", 2)
	fmt.Println(encrypted)
}

func CaesarCipherEncriptor(str string, key int) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	var encryptedStr string
	for _, char := range str {
		fmt.Println(char)
		index := strings.IndexRune(alphabet, char)
		cipherIndex := (index + key) % len(alphabet)
		encryptedStr = fmt.Sprintf("%s%c", encryptedStr, alphabet[cipherIndex])
	}
	return encryptedStr
}
