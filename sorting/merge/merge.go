package main

import "fmt"

func mergeSort(items []int) []int {

	if len(items) < 2 {
		return items
	}

	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])

	return merge(first, second)

}

func merge(a []int, b []int) []int {
	r := []int{}
	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			r = append(r, a[i])
			i++
		} else {
			r = append(r, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		r = append(r, a[i])
	}
	for ; j < len(b); j++ {
		r = append(r, b[j])
	}

	return r
}

func main() {
	unsorted := []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9}
	sorted := mergeSort(unsorted)
	fmt.Println(sorted)
	// sorted = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
}
