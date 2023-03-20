package main

func isUnique_1(s string) bool {
	chars := make(map[rune]struct{})
	for _, c := range s {
		if _, ok := chars[c]; ok {
			return false
		}
		chars[c] = struct{}{}
	}
	return true
}

// No additional Data Structure
func isUnique_2(s string) bool {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}
