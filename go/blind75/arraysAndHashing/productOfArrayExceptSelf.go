package arraysAndHashing

// # 238. Product of Array Except Self (Medium)
func ProductExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	prefix := 1
	for i, num := range nums {
		res[i] = prefix
		prefix *= num
	}

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= postfix
		postfix *= nums[i]
	}
	return res
}
