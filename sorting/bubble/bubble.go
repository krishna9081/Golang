package main

import "fmt"

func main() {

	a := []int{12, 1, 5, 6, 82, 34, 56}
	BubbleSort(a)
	fmt.Println(a)
}

func BubbleSort(a []int) {

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1; j++ {
			if a[j+1] < a[j] {
				a[j], a[j+1] = a[j+1], a[j]
			}

		}
	}
}
