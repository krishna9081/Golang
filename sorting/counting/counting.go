//No errors in First Execution ...Thanks Mrs.M

package main

import "fmt"

func main() {

	a := []int{12, 72, 63, 6, 82, 34, 56}
	r := countingSort(a, 100)
	fmt.Println(r)
}

func countingSort(arr []int, k int) []int {

	c := make([]int, k)
	// Caluculate the Frequency of each element that occurs in the Original array and Fill the Counter Array
	for i := 0; i < len(arr); i++ {
		c[arr[i]] += 1
	}
	/*
		Caluculate Prefix sums of each element in counter array [Gives you the target position value
		Hint: Sum won't increase if there are Zeros in the counter array
		Also helps you with the duplicates
	*/
	for i := 1; i < len(c); i++ {
		c[i] += c[i-1]
	}
	//Declare the Resultant array
	res := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		// Find the original element to be placed
		e := arr[i]
		//Find target postion
		t := c[e] - 1

		//Place element at the target postion in result array
		res[t] = e
		//Decrease the frequency on the counter array
		c[e] = c[e] - 1
	}

	return res
}
