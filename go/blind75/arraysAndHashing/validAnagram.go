package arraysAndHashing

// 217. Contains Duplicate (Easy)
func IsAnagram(s string, t string) bool {
	s_map := map[rune]int{}
	t_map := map[rune]int{}

	for _, char := range s {
		if _, ok := s_map[char]; !ok {
			s_map[char] = 1
		} else {
			s_map[char] += 1
		}
	}

	for _, char := range t {
		if _, ok := t_map[char]; !ok {
			t_map[char] = 1
		} else {
			t_map[char] += 1
		}
	}

	for _, char := range s {
		if s_map[char] != t_map[char] {
			return false
		}
	}

	return true
}
