package main

import (
	"fmt"

	"github.com/yuzinu/leetcode/go/blind75/arraysAndHashing"
)

func main() {
	// 217. Contains Duplicate (Easy)
	// var containsDuplicateTest1 = arraysAndHashing.ContainsDuplicate([]int{1, 2, 3, 1})
	// var containsDuplicateTest2 = arraysAndHashing.ContainsDuplicate([]int{1, 2, 3, 4})
	// var containsDuplicateTest3 = arraysAndHashing.ContainsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})
	// fmt.Println(containsDuplicateTest1)
	// fmt.Println(containsDuplicateTest2)
	// fmt.Println(containsDuplicateTest3)

	// 242. Valid Anagram (Easy)
	// var isAnagramTest1 = arraysAndHashing.IsAnagram("anagram", "nagaram")
	// var isAnagramTest2 = arraysAndHashing.IsAnagram("rat", "car")
	// fmt.Println(isAnagramTest1)
	// fmt.Println(isAnagramTest2)

	// 1. Two Sum (Easy)
	var twoSumTest1 = arraysAndHashing.TwoSum([]int{2, 7, 11, 15}, 9)
	var twoSumTest2 = arraysAndHashing.TwoSum([]int{3, 2, 4}, 6)
	var twoSumTest3 = arraysAndHashing.TwoSum([]int{3, 3}, 6)
	fmt.Println(twoSumTest1)
	fmt.Println(twoSumTest2)
	fmt.Println(twoSumTest3)
}
