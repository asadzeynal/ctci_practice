package ch1

import (
	"fmt"
	"testing"
)

func TestStringRotation(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want bool
	}{
		{"abc", "cab", true},
		{"abc", "caba", false},
		{"abcdefga", "aabcdefg", true},
		{"waterbottle", "erbottlewat", true},
		{"abcd", "dcab", false},
		{"abcd", "cdab", true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v %v", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := stringRotation(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
