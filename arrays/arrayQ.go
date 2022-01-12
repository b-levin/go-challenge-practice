package arrays

import "sort"

//https://practice.geeksforgeeks.org/problems/subarray-with-given-sum-1587115621/1
func subArrSum(arr []int, sum int) (int, int) {
	for i := range arr {
		if j := subSum(arr[i:], sum); j > -1 {
			return i + 1, i + j + 1
		}
	}
	return -1, -1
}

func subSum(arr []int, sum int) int {
	tmpSum := 0
	for i, v := range arr {
		tmpSum += v
		if tmpSum == sum {
			return i
		} else if tmpSum > sum {
			break
		}
	}
	return -1
}

//https://practice.geeksforgeeks.org/problems/count-the-triplets4615/1
func countTriplets(arr []int) int {
	sort.Ints(arr)
	total := 0
	for i := range arr {
		if i+2 >= len(arr) {
			return total
		} else if arr[i]+arr[i+1] == arr[i+2] {
			total++
		}
	}
	return 0
}
