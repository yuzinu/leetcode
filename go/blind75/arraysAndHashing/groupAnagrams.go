package arraysAndHashing

import (
	"sort"
	"strings"
)

// 49. Group Anagrams (Medium)
func GroupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, str := range strs {
		key := sortString(str)
		if _, ok := groups[key]; !ok {
			groups[key] = []string{}
		}
		groups[key] = append(groups[key], str)
	}

	res := make([][]string, len(groups))
	i := 0
	for _, anagrams := range groups {
		res[i] = anagrams
		i++
	}

	return res
}

func sortString(str string) string {
	strArr := strings.Split(str, "")
	sort.Strings(strArr)
	return strings.Join(strArr, "")
}
