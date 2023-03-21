package main

import (
	"strconv"
	"strings"
)

func stringCompression(s string) string {
	if len(s) == 0 {
		return ""
	}

	var sb strings.Builder
	var currentCount int
	for i := range s {
		if currentCount == 0 {
			sb.WriteByte(s[i])
			currentCount++
			continue
		}

		if s[i] == s[i-1] {
			currentCount++
		} else {
			sb.WriteString(strconv.Itoa(currentCount))
			currentCount = 1
			sb.WriteByte(s[i])
		}

		if i == len(s)-1 {
			sb.WriteString(strconv.Itoa(currentCount))
		}
	}

	if len(s) > sb.Len() {
		return sb.String()
	}

	return s
}
