package arraysAndHashing

// 347. Top K Frequent Elements (Medium)
func TopKFrequent(nums []int, k int) (res []int) {
	countMap := make(map[int]int)
	countSlice := make([][]int, len(nums)+1)

	for _, num := range nums {
		if count, ok := countMap[num]; !ok {
			countMap[num] = 1
		} else {
			countMap[num] = count + 1
		}
	}

	for num, count := range countMap {
		countSlice[count] = append(countSlice[count], num)
	}

	for i := len(countSlice) - 1; i > 0; i-- {
		res = append(res, countSlice[i]...)
		if len(res) == k {
			return
		}
	}

	return
}
