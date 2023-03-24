package ch2

import (
	"fmt"
	"testing"

	"github.com/asadzeynal/ctci_practice/util"
)

func TestSumIsLinkedListPalindrome(t *testing.T) {
	var tests = []struct {
		n    []int
		want bool
	}{
		{[]int{}, true},
		{[]int{1, 2}, false},
		{[]int{1}, true},
		{[]int{1, 1}, true},
		{[]int{7, 1, 6}, false},
		{[]int{7, 1, 7}, true},
		{[]int{7, 7, 7, 6, 6, 6}, false},
		{[]int{7, 7, 7, 7, 7}, true},
		{[]int{7, 7, 7, 7}, true},
		{[]int{7, 1, 7, 1, 7}, true},
		{[]int{1, 2, 3, 4, 5, 4, 3, 2, 1, 0}, false},
		{[]int{1, 2, 3, 4, 5, 4, 3, 2, 1}, true},
		{[]int{1, 2, 3, 4, 4, 3, 2, 1}, true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.n)
		t.Run(testname, func(t *testing.T) {
			node := util.LinkedListFromSlice(tt.n)

			ans := isLinkedListPalindrome(node)

			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.n)
		t.Run(testname, func(t *testing.T) {
			node := util.LinkedListFromSlice(tt.n)

			ans := isLinkedListPalindromeV2(node)

			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
