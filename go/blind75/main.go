package main

import (
	"fmt"

	"github.com/yuzinu/leetcode/go/blind75/arraysAndHashing"
)

func main() {
	// Contains Duplicate
	var test1 = arraysAndHashing.ContainsDuplicate([]int{1, 2, 3, 1})
	var test2 = arraysAndHashing.ContainsDuplicate([]int{1, 2, 3, 4})
	var test3 = arraysAndHashing.ContainsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})
	fmt.Println(test1)
	fmt.Println(test2)
	fmt.Println(test3)
}
