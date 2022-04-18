package hashing

import (
	"fmt"
	"testing"
)

func TestArrEquals(t *testing.T) {
	var tests = []struct {
		arr1 []int
		arr2 []int
		want int
	}{
		{[]int{1, 2, 5, 4, 0}, []int{2, 4, 5, 0, 1}, 1},
		{[]int{1, 2, 5}, []int{2, 4, 15}, 0},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d, %d", tt.arr1, tt.arr2)
		t.Run(testName, func(t *testing.T) {
			ans := ArrEqual(tt.arr1, tt.arr2)
			if ans != tt.want {
				t.Errorf("Got: %d, want: %d", ans, tt.want)
			}
		})
	}
}
