package arraysAndHashing

// 1. Two Sum (Easy)
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {

		if j, found := m[target-num]; found {
			return []int{j, i}
		}

		m[num] = i
	}

	return nil
}
