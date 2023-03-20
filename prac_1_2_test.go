package main

import (
	"fmt"
	"testing"
)

func TestIsPermutation(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want bool
	}{
		{"abc", "cab", true},
		{"abc", "caba", false},
		{"abcdefga", "abcdefg", false},
		{"abcdefga", "aabcdefg", true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v %v", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := isPermutation(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
