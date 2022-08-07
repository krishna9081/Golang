package main

import (
	"fmt"
	"math/rand"
)

func main() {

	a := []int{12, 1, 5, 6, 82, 34, 56}
	a = QuickSort(a)
	fmt.Println(a)
}

func QuickSort(arr []int) []int {

	if len(arr) < 2 {
		return arr

	}

	left, right := 0, len(arr)-1
	// Find pivot
	pivotidx := rand.Int() % len(arr)
	pivot := arr[pivotidx]
	//move element at pivot to the right of the array
	arr[pivotidx], arr[right] = arr[right], arr[pivotidx]

	// Scoop all smaller elements to the left of the pivot

	for i := range arr {
		if arr[i] < pivot {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	QuickSort(arr[:left])
	QuickSort(arr[left+1:])

	return arr

}
