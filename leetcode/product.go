package main

import "fmt"

type SparseVector struct {
	nonZeros [][]int
}

func Constructor(nums []int) SparseVector {
	m := [][]int{}
	for i, val := range nums {
		if val != 0 {
			m = append(m, []int{i, val})
		}
	}

	return SparseVector{m}

}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {

	a := this.nonZeros
	b := vec.nonZeros

	if len(a) >= len(b) {
		return binarySearch(a, b)
	} else {
		return binarySearch(b, a)
	}

	//	fmt.Println(a)
	//	fmt.Println(b)

}

func binarySearch(a, b [][]int) int {
	res := 0
	for i := range a {

		low := 0
		high := len(b) - 1

		for low <= high {
			//median = (high + low) / 2 //but it can overflow
			median := int(uint(high+low) >> 1) // avoid overflow when computing median
			if b[median][0] == a[i][0] {
				res += a[i][1] * b[median][1]
				break
			}
			if b[median][0] < a[i][0] {
				low = median + 1
			} else {
				high = median - 1
			}

		}

	}
	return res
}

func main() {
	nums1 := []int{1, 0, 0, 2, 3}
	nums2 := []int{0, 3, 0, 4, 0}
	v1 := Constructor(nums1)
	v2 := Constructor(nums2)
	ans := v1.dotProduct(v2)

	fmt.Println(ans)

}
