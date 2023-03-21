package main

import (
	"fmt"
	"testing"
)

func TestIsPalindromePermutation(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"tacotac", true},
		{"abc", false},
		{"aaabbbccc", false},
		{"aabbcceef", true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.s)
		t.Run(testname, func(t *testing.T) {
			ans := isPalindromePermutation(tt.s)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
