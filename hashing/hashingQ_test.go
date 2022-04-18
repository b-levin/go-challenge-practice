package hashing

import (
	"fmt"
	"reflect"
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

func TestCommonEle(t *testing.T) {
	var tests = []struct {
		arr1 []int
		arr2 []int
		want []int
	}{
		{[]int{3, 4, 2, 2, 4}, []int{3, 2, 2, 7}, []int{2, 2, 3}},
		{[]int{5, 4, 3, 6, 7, 1, 2}, []int{5, 3, 9, 10, 4, 7, 8}, []int{3, 4, 5, 7}},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d, %d", tt.arr1, tt.arr2)
		t.Run(testName, func(t *testing.T) {
			ans := CommonEle(tt.arr1, tt.arr2)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("Got: %d, want: %d", ans, tt.want)
			}
		})
	}
}
