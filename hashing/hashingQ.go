package hashing

import "sort"

//https://practice.geeksforgeeks.org/problems/check-if-two-arrays-are-equal-or-not3847/1
func ArrEqual(arr1 []int, arr2 []int) int {
	arr1Map := make(map[int]bool)
	for _, val := range arr1 {
		arr1Map[val] = true
	}
	for _, val := range arr2 {
		if _, prs := arr1Map[val]; !prs {
			return 0
		}
	}
	return 1
}

//https://practice.geeksforgeeks.org/problems/common-elements5420/1
func CommonEle(arr1 []int, arr2 []int) []int {
	out := []int{}
	//sort arrs
	sort.Ints(arr1)
	sort.Ints(arr2)
	//assuming that array size is equal
	arr1Map := make(map[int]bool)
	for _, val := range arr1 {
		arr1Map[val] = true
	}
	for _, val := range arr2 {
		if _, prs := arr1Map[val]; prs {
			out = append(out, val)
		}
	}
	return out
}

//https://practice.geeksforgeeks.org/problems/first-element-to-occur-k-times5150/1
func FirstKTimes(arr []int, k int) int {
	//sort arr
	sort.Ints(arr)
	countMap := make(map[int]int)
	for _, key := range arr {
		if v, prs := countMap[key]; prs {
			if v++; v == k {
				return key
			} else {
				countMap[key] = v
			}
		} else {
			countMap[key] = 1
		}
	}
	// no elements occuring k times
	return -1
}
