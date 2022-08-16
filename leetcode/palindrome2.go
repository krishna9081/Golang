package main

import (
	"fmt"
)

func validPalindrome(s string) bool {

	return checkPalin(s, 0, len(s)-1, 0)

}

func checkPalin(s string, i int, j int, count int) bool {

	for i < j {
		if s[i] != s[j] {
			if count >= 1 {
				return false
			} else {
				return checkPalin(s, i+1, j, count+1) || checkPalin(s, i, j-1, count+1)
			}
		}
		i++
		j--

	}
	return true
}

func main() {
	fmt.Println(validPalindrome("aba"))
}
