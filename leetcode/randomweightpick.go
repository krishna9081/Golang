/* Random weight pick
Prefix sum
Binary Search
Cumulutaive frequency distribution
*/

import "math/rand"

type Solution struct {
	PrefixSums []int
}

// Prefix Sum O(n)
func Constructor(w []int) Solution {

	prefixSum := make([]int, len(w))
	totalSum := 0

	for idx := range w {

		totalSum += w[idx]
		prefixSum[idx] = totalSum

	}

	return Solution{PrefixSums: prefixSum}

}

func (this *Solution) PickIndex() int {

	target := rand.Intn(this.PrefixSums[len(this.PrefixSums)-1]) //O(n)

	// implement binary search O(logn)
	low, high := 0, len(this.PrefixSums)
	for low < high {
		median := int(uint(high+low) >> 1) // avoid overflow when computing median
		if this.PrefixSums[median] <= target {
			low = median + 1
		} else {
			high = median
		}
	}

	return low
}

