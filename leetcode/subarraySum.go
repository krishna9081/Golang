func subarraySum(nums []int, k int) int {

	// A map containing the prefix Sums
	m := make(map[int]int)

	// A variable for caluculating Prefix Sums
	prefixSum := 0
	//Initialize the map with 0 sum of count 1
	m[0] = 1

	//counter for calulating the number of times the subset sum equals K
	count := 0

	//Iterate over the range of nums
	for _, val := range nums {
		//Caluculate prefix sum at each point
		prefixSum += val
		// At any instace the counter can be increased if the difference between prefix Sum and k exists in the map of Sums
		count += m[prefixSum-k]

		//Increase the counter for each sum in general
		m[prefixSum]++

	}

	return count
}