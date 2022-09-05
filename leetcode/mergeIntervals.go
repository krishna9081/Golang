func merge(intervals [][]int) [][]int {

	//Creating a resulting array of size 0

	res := make([][]int, 0)

	//Sorting the given input array ..this will sort the input into Sorted output

	sort.Slice(intervals, func(x, y int) bool { return intervals[x][0] < intervals[y][0] })

	//Initialize the resultant list with first interval of intervals

	res = append(res, intervals[0])

	// Iterate over intervals

	for i := 1; i < len(intervals); i++ {
		//compare the last element of last index of the resultant array with first element first Index of the Intervals array
		//if it is greater , merge the interval
		if res[len(res)-1][1] >= intervals[i][0] {
			res[len(res)-1][1] = max(intervals[i][1], res[len(res)-1][1]) //missed the max part , have to be careful
		} else { //else ...append the new element of intervals to the result.
			res = append(res, intervals[i])
		}

	}

	return res

}

func max(x, y int) int {

	if x > y {
		return x
	} else {
		return y
	}

}