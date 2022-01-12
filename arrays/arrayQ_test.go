package arrays

import (
	"fmt"
	"testing"
)

func TestSubArrSum1(t *testing.T) {
	var tests = []struct {
		sum          int
		arr          []int
		want1, want2 int
	}{
		{12, []int{1, 2, 3, 7, 5},
			2, 4},
		{15, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			1, 5},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d,%d", tt.sum, tt.arr)
		t.Run(testName, func(t *testing.T) {
			ans1, ans2 := subArrSum(tt.arr, tt.sum)
			if ans1 != tt.want1 || ans2 != tt.want2 {
				t.Errorf("Got: %d | %d, want: %d | %d", ans1, ans2, tt.want1, tt.want2)
			}
		})
	}
}

func TestCountTriplets(t *testing.T) {
	var tests = []struct {
		arr  []int
		want int
	}{
		{[]int{1, 5, 3, 2}, 2},
		{[]int{2, 3, 4}, 0},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d", tt.arr)
		t.Run(testName, func(t *testing.T) {
			ans := countTriplets(tt.arr)
			if ans != tt.want {
				t.Errorf("Got: %d, want: %d", ans, tt.want)
			}
		})
	}
}
