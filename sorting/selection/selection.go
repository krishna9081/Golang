package main

import "fmt"

func main() {

	a := []int{12, 1, 5, 6, 82, 34, 56}
	SelectionSort(a)
	fmt.Println(a)
}

func SelectionSort(a []int) {

	size := len(a)
	var i, min, minidx int

	for i = 0; i < size; i++ {
		min = a[i]
		for j := i + 1; j < size; j++ {
			if a[j] < min {
				min = a[j]
				minidx = j
			}

		}
		a[i], a[minidx] = a[minidx], a[i]

	}

}
