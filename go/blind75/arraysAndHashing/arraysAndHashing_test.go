package arraysAndHashing

import (
	"reflect"
	"testing"
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
