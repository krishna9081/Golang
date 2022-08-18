package main

import "fmt"

func minRemoveToMakeValid(s string) string {
	//Declaring a Slice
	stack := []int{}
	//Iterating over slice
	for i := 0; i < len(s); i++ {
		// String check ...Should be single quotes ..Keeping the index of opening braces in the stack
		if s[i] == '(' {
			//appened method is different from Python
			stack = append(stack, i)
		} else if s[i] == ')' && len(stack) != 0 { //unwinding the stack as long as you see matching open braces until the length of the stack
			//Bottom line ...never add index of closing braces to the stack
			stack = stack[:len(stack)-1]
		} else if s[i] == ')' { // If the length of stack is already empty ...if you still find a closing brace from the stack
			// Just make the element disappear from the stack Bottom line ...never add index of closing braces to the stack
			s = s[:i] + s[i+1:]
			i--
		}
	}

	// You may still find opening brace indices left out in the stack ... capture those one by one , remove those from the actual string
	//Resultant string is the one which has matching parentheses with minimum number of removals
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
