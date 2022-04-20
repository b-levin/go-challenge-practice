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

func TestFirstKTimes(t *testing.T) {
	var tests = []struct {
		arr  []int
		k    int
		want int
	}{
		{[]int{1, 7, 4, 3, 4, 8, 7}, 2, 4},
		{[]int{3, 4, 6, 5, 6, 11, 9, 11, 11, 9, 6}, 3, 6},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d, %d", tt.arr, tt.k)
		t.Run(testName, func(t *testing.T) {
			ans := FirstKTimes(tt.arr, tt.k)
			if ans != tt.want {
				t.Errorf("Got: %d, want: %d", ans, tt.want)
			}
		})
	}
}
