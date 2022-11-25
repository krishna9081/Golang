package main

import "fmt"

func main() {

	a := []int{12, 1, 5, 6, 82, 34, 56}
	a = reverse(a, 3, 6)
	fmt.Println(a)
}

func reverse(a []int, l, r int) []int {

	for l < r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
	return a
}
