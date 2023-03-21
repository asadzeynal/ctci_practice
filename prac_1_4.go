package main

func isPalindromePermutation(s string) bool {
	charsCount := make(map[rune]int)
	for _, c := range s {
		charsCount[c]++
	}

	var oddCount int
	for _, v := range charsCount {
		if v%2 == 1 {
			oddCount++
		}
	}

	return oddCount <= 1
}
