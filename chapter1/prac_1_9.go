package ch1

import "strings"

func stringRotation(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}

	return strings.Contains(a+a, b)
}
