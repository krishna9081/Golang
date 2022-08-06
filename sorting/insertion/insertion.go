//My First program without referring code elsewhere ...Thanks Mr.S

package main

import "fmt"

func insertionSort(arr []int) []int {
	var i, j int

	if len(arr) < 2 {
		return arr
	}

	for i = 1; i < len(arr); i++ {
		j = i
		for j > 0 && arr[j] < arr[j-1] {

			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--

		}

	}

	return arr
}

func main() {

	a := []int{12, 1, 5, 6, 82, 34, 56}
	a = insertionSort(a)
	fmt.Println(a)
}
