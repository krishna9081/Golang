func findBuildings(heights []int) []int {

	res := []int{}

	for i, ele := range heights {
		for len(res) != 0 && heights[res[len(res)-1]] <= ele {
			res = res[:len(res)-1]
		}
		res = append(res, i)
	}

	return res
}