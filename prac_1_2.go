package main

func isPermutation(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}

	charsCount := make(map[rune]int)
	for _, c := range a {
		charsCount[c]++
	}
	for _, c := range b {
		charsCount[c]--
	}

	for _, count := range charsCount {
		if count != 0 {
			return false
		}
	}
	return true
}
