package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {

	var start, end, maxLen int
	m := make(map[byte]int)

	//moving the end pointer
	for end < len(s) {

		// If a repeating character is found stop and update start pointer
		if _, ok := m[s[end]]; ok {
			start = max(start, m[s[end]]+1)
		}
		//if not found , update map
		m[s[end]] = end
		maxLen = max(maxLen, end-start+1)
		end++ // we're incrementing end pointer  at the end of loop
	}
	return maxLen
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {

	fmt.Println(lengthOfLongestSubstring("abcabcbb"))

}
