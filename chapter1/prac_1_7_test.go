package ch1

import (
	"fmt"
	"testing"

	"github.com/asadzeynal/ctci_practice/util"
)

func TestRotateMatrix(t *testing.T) {
	tests := []struct {
		s    [][]int
		want [][]int
	}{
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}},
			[][]int{{13, 9, 5, 1}, {14, 10, 6, 2}, {15, 11, 7, 3}, {16, 12, 8, 4}},
		},
		{
			[][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}},
			[][]int{{21, 16, 11, 6, 1}, {22, 17, 12, 7, 2}, {23, 18, 13, 8, 3}, {24, 19, 14, 9, 4}, {25, 20, 15, 10, 5}},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.s)
		t.Run(testname, func(t *testing.T) {
			ans := rotateMatrix(tt.s)
			if !util.EqualMatrix(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestRotateMatrixFaster(t *testing.T) {
	tests := []struct {
		s    [][]int
		want [][]int
	}{
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}},
			[][]int{{13, 9, 5, 1}, {14, 10, 6, 2}, {15, 11, 7, 3}, {16, 12, 8, 4}},
		},
		{
			[][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}},
			[][]int{{21, 16, 11, 6, 1}, {22, 17, 12, 7, 2}, {23, 18, 13, 8, 3}, {24, 19, 14, 9, 4}, {25, 20, 15, 10, 5}},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.s)
		t.Run(testname, func(t *testing.T) {
			ans := rotateMatrixFaster(tt.s)
			if !util.EqualMatrix(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func BenchmarkRotateMatrix(b *testing.B) {
	testMatrix := [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}}
	for i := 0; i < b.N; i++ {
		rotateMatrix(testMatrix)
	}
}

func BenchmarkRotateMatrixFaster(b *testing.B) {
	testMatrix := [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}}
	for i := 0; i < b.N; i++ {
		rotateMatrixFaster(testMatrix)
	}
}
