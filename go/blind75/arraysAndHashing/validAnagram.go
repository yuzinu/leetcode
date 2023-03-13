package arraysAndHashing

// 217. Contains Duplicate (Easy)

// hashmap solution
// func IsAnagram(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}

// 	s_map := map[rune]int{}
// 	t_map := map[rune]int{}

// 	for _, char := range s {
// 		if _, ok := s_map[char]; !ok {
// 			s_map[char] = 1
// 		} else {
// 			s_map[char] += 1
// 		}
// 	}

// 	for _, char := range t {
// 		if _, ok := t_map[char]; !ok {
// 			t_map[char] = 1
// 		} else {
// 			t_map[char] += 1
// 		}
// 	}

// 	for _, char := range s {
// 		if s_map[char] != t_map[char] {
// 			return false
// 		}
// 	}

// 	return true
// }

// array solution
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var freq [26]int

	for idx := 0; idx < len(s); idx++ {
		freq[s[idx]-'a']++
		freq[t[idx]-'a']--
	}

	for idx := 0; idx < len(freq); idx++ {
		if freq[idx] != 0 {
			return false
		}
	}

	return true
}
