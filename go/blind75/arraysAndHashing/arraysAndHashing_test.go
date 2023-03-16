package arraysAndHashing

import "testing"

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
	for _, test := range containsDuplicates {
		if output := ContainsDuplicate(test.arg1); output != test.expected {
			t.Errorf("Output %t not equal to expected %t", output, test.expected)
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
	for _, test := range isAnagrams {
		if output := IsAnagram(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %t not equal to expected %t", output, test.expected)
		}
	}
}
