package main

import (
	"fmt"
	"testing"
)

func TestStringCompression(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"aabcccccaaa", "a2b1c5a3"},
		{"aabccccca", "a2b1c5a1"},
		{"abcde", "abcde"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.s)
		t.Run(testname, func(t *testing.T) {
			ans := stringCompression(tt.s)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
