package hashing

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
