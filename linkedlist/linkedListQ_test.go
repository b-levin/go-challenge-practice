package linkedlist

import (
	"fmt"
	"testing"
)

func TestFindMiddle(t *testing.T) {
	var tests = []struct {
		list linkedList
		want int
	}{
		{*listFromArr([]int{1, 2, 3, 4, 5}), 3},
		{*listFromArr([]int{2, 4, 6, 7, 5, 1}), 7},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf(tt.list.toString())
		t.Run(testName, func(t *testing.T) {
			ans := findMiddle(tt.list)
			if ans != tt.want {
				t.Errorf("Got: %d, want: %d", ans, tt.want)
			}
		})
	}
}
