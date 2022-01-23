package arrays

import (
	"sort"
)

/**
TODO: tie together with cobra??
*/

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

//https://practice.geeksforgeeks.org/problems/inversion-of-array-1587115620/1
func countInversions(arr []int) int {
	out := 0
	outOfOrder := true
	for outOfOrder {
		swaps := 0
		for i := 0; i+1 < len(arr); i++ {
			if arr[i] > arr[i+1] {
				swaps++
				tmp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = tmp
			}
		}
		if swaps == 0 {
			outOfOrder = false
		} else {
			out += swaps
		}
	}
	return out
}

//https://practice.geeksforgeeks.org/problems/sort-an-array-of-0s-1s-and-2s4231/1
func sort012(arr []int) []int {
	out := []int{}
	vals := []int{0, 0, 0}
	for _, v := range arr {
		vals[v] = vals[v] + 1
	}
	for i, v := range vals {
		if v > 0 {
			for j := 0; j < v; j++ {
				out = append(out, i)
			}
		}
	}
	return out
}

//https://practice.geeksforgeeks.org/problems/equilibrium-point-1587115620/1
func equilibriumPoint(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	} else {
		lSum := 0
		rSum := sumArr(arr)
		for i := 0; i < len(arr); i++ {
			if lSum == rSum-arr[i] {
				return i + 1
			} else {
				lSum += arr[i]
				rSum -= arr[i]
			}
		}
	}
	return -1
}

func sumArr(arr []int) int {
	out := 0
	for _, v := range arr {
		out += v
	}
	return out
}

//https://practice.geeksforgeeks.org/problems/leaders-in-an-array-1587115620/1
func leaders(arr []int) []int {
	vals := []int{}
	highNum := arr[len(arr)-1]
	vals = append(vals, highNum)
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] > highNum {
			highNum = arr[i]
			vals = append(vals, highNum)
		}
	}
	out := []int{}
	for i := len(vals) - 1; i >= 0; i-- {
		out = append(out, vals[i])
	}
	return out
}

//https://practice.geeksforgeeks.org/problems/minimum-platforms-1587115620/1
func minPlatforms(arr []int, dep []int) int {
	maxPlats := 1
	platCnt := 1
	for i := 0; i < len(arr)-1; i++ {
		if dep[i] > arr[i+1] {
			platCnt++
		} else {
			if platCnt > maxPlats {
				maxPlats = platCnt
				platCnt = 1
			}
		}
	}
	return maxPlats
}

//https://practice.geeksforgeeks.org/problems/reverse-array-in-groups0255/1
func reverseGroups(arr []int, subCount int) []int {
	out := []int{}
	curFlipPos := subCount - 1
	for curFlipPos <= len(arr)-1 {
		for i := curFlipPos; i > curFlipPos-subCount; i-- {
			out = append(out, arr[i])
		}
		curFlipPos += subCount
	}
	//fill out remaining elements
	if len(out) != len(arr) {
		i := len(arr) - 1
		for len(out) != len(arr) {
			out = append(out, arr[i])
			i--
		}
	}
	return out
}

//https://practice.geeksforgeeks.org/problems/kth-smallest-element5635/1
func kSmallest(arr []int, pos int) int {
	sort.Ints(arr)
	return arr[pos-1]
}

//https://practice.geeksforgeeks.org/problems/trapping-rain-water-1587115621/1
func trapRainWater(arr []int) int {
	total := 0
	lMax := 0
	rMax := 0
	lPos := 0
	rPos := len(arr) - 1
	for lPos < rPos {
		if arr[lPos] < arr[rPos] {
			lMax = max(lMax, arr[lPos])
			total += lMax - arr[lPos]
			lPos++
		} else {
			rMax = max(rMax, arr[rPos])
			total += rMax - arr[rPos]
			rPos--
		}
	}
	return total
}

func max(i int, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

//https://practice.geeksforgeeks.org/problems/pythagorean-triplet3018/1
func pyTriplet(arr []int) bool {
	sort.Ints(arr)
	if len(arr) < 3 {
		return false
	}
	for i := 0; i < len(arr)-3; i++ {
		if isPythag(arr[i], arr[i+1], arr[i+2]) {
			return true
		}
	}
	return false
}

func isPythag(a int, b int, c int) bool {
	return (a*a + b*b) == c*c
}

//https://practice.geeksforgeeks.org/problems/chocolate-distribution-problem3825/1
func chocolateDist(arr []int, m int) int {
	sort.Ints(arr)
	out := arr[m-1] - arr[0]
	for i := 1; i < len(arr)-m; i++ {
		diff := arr[i+m-1] - arr[i]
		if out > diff {
			out = diff
		}
	}
	return out
}
