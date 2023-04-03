package arraysAndHashing

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 217. Contains Duplicate (Easy)
type containsDuplicate struct {
	arg1     []int
	expected bool
}

var containsDuplicates = []containsDuplicate{
	{[]int{1, 2, 3, 1}, true},
	{[]int{1, 2, 3, 4}, false},
	{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
}

func TestContainsDuplicate(t *testing.T) {
	for i, test := range containsDuplicates {
		if output := ContainsDuplicate(test.arg1); output != test.expected {
			t.Errorf("Output %t not equal to expected %t for test %d", output, test.expected, i+1)
		}
	}
}

// 242. Valid Anagram (Easy)
type isAnagram struct {
	arg1     string
	arg2     string
	expected bool
}

var isAnagrams = []isAnagram{
	{"anagram", "nagaram", true},
	{"rat", "car", false},
}

func TestIsAnagram(t *testing.T) {
	for i, test := range isAnagrams {
		if output := IsAnagram(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %t not equal to expected %t for test %d", output, test.expected, i+1)
		}
	}
}

// 1. Two Sum (Easy)
type twoSum struct {
	arg1     []int
	arg2     int
	expected []int
}

var twoSums = []twoSum{
	{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	{[]int{3, 2, 4}, 6, []int{1, 2}},
	{[]int{3, 3}, 6, []int{0, 1}},
}

func TestTwoSum(t *testing.T) {
	for i, test := range twoSums {
		if output := TwoSum(test.arg1, test.arg2); !reflect.DeepEqual(output, test.expected) {
			t.Errorf("Output %q not equal to expected %q for test %d", output, test.expected, i+1)
		}
	}
}

// 49. Group Anagrams (Medium)
type groupAnagram struct {
	arg1     []string
	expected [][]string
}

var groupAnagrams = []groupAnagram{
	{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
	{[]string{""}, [][]string{{""}}},
	{[]string{"a"}, [][]string{{"a"}}},
}

func TestGroupAnagrams(t *testing.T) {
	for i, test := range groupAnagrams {
		t.Run(fmt.Sprintf("Group anagrams test %d", i+1), func(t *testing.T) {
			output := GroupAnagrams(test.arg1)

			// Sorting got inners
			for _, inner := range output {
				sort.Slice(inner, func(i, j int) bool {
					return inner[i] < inner[j]
				})
			}

			// Sorting want inners
			for _, inner := range test.expected {
				sort.Slice(inner, func(i, j int) bool {
					return inner[i] < inner[j]
				})
			}

			// Match
			if ok := assert.ElementsMatch(t, output, test.expected); !ok {
				t.Errorf("Output %q not equal to expected %q for test %d", output, test.expected, i+1)
			}
		})
	}
}

// 347. Top K Frequent Elements (Medium)
type topKFrequent struct {
	arg1     []int
	arg2     int
	expected []int
}

var topKFrequents = []topKFrequent{
	{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
	{[]int{1}, 1, []int{1}},
	{[]int{200, 200, 0, 1, 1, 3, 3, 3, 200}, 3, []int{200, 3, 1}},
}

func TestTopKFrequent(t *testing.T) {
	for i, test := range topKFrequents {
		output := TopKFrequent(test.arg1, test.arg2)
		if ok := assert.ElementsMatch(t, output, test.expected); !ok {
			t.Errorf("Output %d not equal to expected %d for test %d", output, test.expected, i+1)
		}
	}
}

// # 238. Product of Array Except Self (Medium)
type productExceptSelf struct {
	arg1     []int
	expected []int
}

var productExceptSelfs = []productExceptSelf{
	{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
	{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
}

func TestProductExceptSelf(t *testing.T) {
	for i, test := range productExceptSelfs {
		if output := ProductExceptSelf(test.arg1); !reflect.DeepEqual(output, test.expected) {
			t.Errorf("Output %q not equal to expected %q for test %d", output, test.expected, i+1)
		}
	}
}
