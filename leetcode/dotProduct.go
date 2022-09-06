type SparseVector struct {
	nonZeros map[int]int
}

func Constructor(nums []int) SparseVector {

	m := map[int]int{}

	for i, v := range nums {
		if v != 0 {
			m[i] = v
		}

	}

	return SparseVector{m}

}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {

	res := 0

	for i, first := range this.nonZeros {

		if second, ok := vec.nonZeros[i]; ok {
			res += first * second

		}

	}

	return res
}

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */