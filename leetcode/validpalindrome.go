package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)

	}
	ps := reg.ReplaceAllString(s, "")
	ps = strings.ToLower(ps)
	j := len(ps) - 1
	for i := 0; i < len(ps); i++ {
		if ps[i] != ps[j] {
			return false
		}
		j--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}
