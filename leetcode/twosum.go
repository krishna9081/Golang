package main

func twoSum(nums []int, target int) []int {

	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {

		t := target - nums[i]

		if j, isPresent := m[t]; isPresent {
			return []int{i, j}

		}

		//storing the Indices in the map
		m[nums[i]] = i

	}
	return []int{}

}
