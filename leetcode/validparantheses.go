package main

import "fmt"

func minRemoveToMakeValid(s string) string {

	stack := []int{}

	for i := 0; i < len(s); i++ {

		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' && len(stack) != 0 {
			stack = stack[:len(stack)-1]
		} else if s[i] == ')' {
			s = s[:i] + s[i+1:]
			i--
		}
	}

	for len(stack) != 0 {
		lastIdx := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		s = s[:lastIdx] + s[lastIdx+1:]
	}

	return s
}

func main() {

	fmt.Println(minRemoveToMakeValid("lee(t(c)o)de)"))
}
