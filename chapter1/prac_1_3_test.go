package ch1

import (
	"fmt"
	"testing"
)

func TestUrlify(t *testing.T) {
	tests := []struct {
		s    string
		l    int
		want string
	}{
		{"Mr John Smith    ", 13, "Mr%20John%20Smith"},
		{"S a l a m A l e y k u m                      ", 23, "S%20a%20l%20a%20m%20A%20l%20e%20y%20k%20u%20m"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v %v", tt.s, tt.l)
		t.Run(testname, func(t *testing.T) {
			ans := urlify(tt.s, tt.l)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
