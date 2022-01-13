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

//https://practice.geeksforgeeks.org/problems/missing-number-in-array1416/1
func missingNums(arr []int, size int) int {
	sort.Ints(arr)
	for i, v := range arr {
		if arr[i+1]-v > 1 {
			return v + 1
		}
	}
	return -1
}

//https://practice.geeksforgeeks.org/problems/merge-two-sorted-arrays-1587115620/1
func mergeArr(arr1 []int, arr2 []int) []int {
	out := []int{}
	pos1 := 0
	pos2 := 0

	for pos1 < len(arr1) && pos2 < len(arr2) {
		if arr1[pos1] < arr2[pos2] {
			out = append(out, arr1[pos1])
			pos1++
		} else {
			out = append(out, arr2[pos2])
			pos2++
		}
	}
	for ; pos1 < len(arr1); pos1++ {
		out = append(out, arr1[pos1])
	}
	for ; pos2 < len(arr2); pos2++ {
		out = append(out, arr2[pos2])
	}
	return out
}

//https://practice.geeksforgeeks.org/problems/-rearrange-array-alternately-1587115620/1
func reArrangeArrAlt(arr []int) []int {
	out := []int{}

	i := 0
	j := len(arr) - 1

	for i+1 != j && i != j {
		out = append(out, arr[j])
		out = append(out, arr[i])
		i++
		j--
	}

	if i == j {
		out = append(out, arr[i])
	}

	if i+1 == j {
		out = append(out, arr[j])
		out = append(out, arr[i])
	}

	return out
}
