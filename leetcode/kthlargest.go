import "math/rand"

func findKthLargest(nums []int, k int) int {

	targetPos := len(nums) - k
	left := 0
	right := len(nums) - 1

	for {
		pivotIdx := partition(nums, left, right)

		if pivotIdx == targetPos {
			return nums[targetPos]
		} else if pivotIdx < targetPos {
			left = pivotIdx + 1

		} else if pivotIdx > targetPos {
			right = pivotIdx - 1
		}
	}

	return -1
}

func partition(nums []int, left, right int) int {

	pivotIdx := left + rand.Intn(right-left+1)
	pivot := nums[pivotIdx]
	nums[pivotIdx], nums[right] = nums[right], nums[pivotIdx]

	for i := left; i < right; i++ {
		if nums[i] < pivot {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}

	nums[left], nums[right] = nums[right], nums[left]
	return left

}