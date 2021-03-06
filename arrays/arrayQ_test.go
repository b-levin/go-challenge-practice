package arrays

import (
	"fmt"
	"reflect"
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

func TestMissingNums(t *testing.T) {
	var tests = []struct {
		arr  []int
		size int
		want int
	}{
		{[]int{1, 2, 3, 5}, 5, 4},
		{[]int{6, 1, 2, 8, 3, 4, 7, 10, 5}, 10, 9},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d,%d", tt.arr, tt.size)
		t.Run(testName, func(t *testing.T) {
			ans := missingNums(tt.arr, tt.size)
			if ans != tt.want {
				t.Errorf("Got: %d, want: %d", ans, tt.want)
			}
		})
	}
}

func TestMergeArr(t *testing.T) {
	var tests = []struct {
		arr1 []int
		arr2 []int
		want []int
	}{
		{[]int{1, 3, 5, 7}, []int{0, 2, 6, 8, 9}, []int{0, 1, 2, 3, 5, 6, 7, 8, 9}},
		{[]int{5, 10}, []int{12, 18, 20}, []int{5, 10, 12, 18, 20}},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d,%d", tt.arr1, tt.arr2)
		t.Run(testName, func(t *testing.T) {
			ans := mergeArr(tt.arr1, tt.arr2)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("Got: %d, want: %d", ans, tt.want)
			}
		})
	}
}

func TestReArrangeArrAlt(t *testing.T) {
	var tests = []struct {
		arr  []int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, []int{6, 1, 5, 2, 4, 3}},
		{[]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110}, []int{110, 10, 100, 20, 90, 30, 80, 40, 70, 50, 60}},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d", tt.arr)
		t.Run(testName, func(t *testing.T) {
			ans := reArrangeArrAlt(tt.arr)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("Got: %d, want: %d", ans, tt.want)
			}
		})
	}
}

func TestCountInversions(t *testing.T) {
	var tests = []struct {
		arr  []int
		want int
	}{
		{[]int{2, 4, 1, 3, 5}, 3},
		{[]int{2, 3, 4, 5, 6}, 0},
		{[]int{10, 10, 10}, 0},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d", tt.arr)
		t.Run(testName, func(t *testing.T) {
			ans := countInversions(tt.arr)
			if ans != tt.want {
				t.Errorf("Got: %d, want: %d", ans, tt.want)
			}
		})
	}
}

func TestSort012(t *testing.T) {
	var tests = []struct {
		arr  []int
		want []int
	}{
		{[]int{0, 2, 1, 2, 0}, []int{0, 0, 1, 2, 2}},
		{[]int{0, 1, 0}, []int{0, 0, 1}},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d", tt.arr)
		t.Run(testName, func(t *testing.T) {
			ans := sort012(tt.arr)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("Got: %d, want: %d", ans, tt.want)
			}
		})
	}
}

func TestEquilibriumPoint(t *testing.T) {
	var tests = []struct {
		arr  []int
		want int
	}{
		{[]int{1, 3, 5, 2, 2}, 3},
		{[]int{1}, 1},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d", tt.arr)
		t.Run(testName, func(t *testing.T) {
			ans := equilibriumPoint(tt.arr)
			if ans != tt.want {
				t.Errorf("Got: %d, want: %d", ans, tt.want)
			}
		})
	}
}

func TestLeaders(t *testing.T) {
	var tests = []struct {
		arr  []int
		want []int
	}{
		{[]int{16, 17, 4, 3, 5, 2}, []int{17, 5, 2}},
		{[]int{1, 2, 3, 4, 0}, []int{4, 0}},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d", tt.arr)
		t.Run(testName, func(t *testing.T) {
			ans := leaders(tt.arr)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("Got: %d, want: %d", ans, tt.want)
			}
		})
	}
}

func TestMinPlatform(t *testing.T) {
	var tests = []struct {
		arr  []int
		dep  []int
		want int
	}{
		{[]int{900, 940, 950, 1100, 1500, 1800}, []int{910, 1200, 1120, 1130, 1900, 2000}, 3},
		{[]int{900, 1100, 1235}, []int{1000, 1200, 1240}, 1},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d,%d", tt.arr, tt.dep)
		t.Run(testName, func(t *testing.T) {
			ans := minPlatforms(tt.arr, tt.dep)
			if ans != tt.want {
				t.Errorf("Got: %d, want: %d", ans, tt.want)
			}
		})
	}
}

func TestReverseGroups(t *testing.T) {
	var tests = []struct {
		arr  []int
		flip int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, []int{3, 2, 1, 5, 4}},
		{[]int{5, 6, 8, 9}, 3, []int{8, 6, 5, 9}},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d,%d", tt.arr, tt.flip)
		t.Run(testName, func(t *testing.T) {
			ans := reverseGroups(tt.arr, tt.flip)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("Got: %d, want: %d", ans, tt.want)
			}
		})
	}
}

func TestKSmallest(t *testing.T) {
	var tests = []struct {
		arr  []int
		pos  int
		want int
	}{
		{[]int{7, 10, 4, 3, 20, 15}, 3, 7},
		{[]int{7, 10, 4, 20, 15}, 4, 15},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d,%d", tt.arr, tt.pos)
		t.Run(testName, func(t *testing.T) {
			ans := kSmallest(tt.arr, tt.pos)
			if ans != tt.want {
				t.Errorf("Got: %d, want: %d", ans, tt.want)
			}
		})
	}
}

func TestTrapRainWater(t *testing.T) {
	var tests = []struct {
		arr  []int
		want int
	}{
		{[]int{3, 0, 0, 2, 0, 4}, 10},
		{[]int{7, 4, 0, 9}, 10},
		{[]int{6, 6, 9}, 0},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d", tt.arr)
		t.Run(testName, func(t *testing.T) {
			ans := trapRainWater(tt.arr)
			if ans != tt.want {
				t.Errorf("Got: %d, want: %d", ans, tt.want)
			}
		})
	}
}

func TestPyTriplet(t *testing.T) {
	var tests = []struct {
		arr  []int
		want bool
	}{
		{[]int{3, 2, 4, 6, 5}, true},
		{[]int{3, 5, 8}, false},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d", tt.arr)
		t.Run(testName, func(t *testing.T) {
			ans := pyTriplet(tt.arr)
			if ans != tt.want {
				t.Errorf("Got: %t, want: %t", ans, tt.want)
			}
		})
	}
}

func TestChocolateDist(t *testing.T) {
	var tests = []struct {
		arr  []int
		m    int
		want int
	}{
		{[]int{3, 4, 1, 9, 56, 7, 9, 12}, 5, 6},
		{[]int{7, 3, 2, 4, 9, 12, 56}, 3, 2},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d", tt.arr)
		t.Run(testName, func(t *testing.T) {
			ans := chocolateDist(tt.arr, tt.m)
			if ans != tt.want {
				t.Errorf("Got: %d, want: %d", ans, tt.want)
			}
		})
	}
}
