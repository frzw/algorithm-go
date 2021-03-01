package array

import _sort "github.com/frzw/algorithm-go/base/sort"

func CombineTwoOrderedArrays(arr1 []int, arr2 []int) []int {
	lenArr1 := len(arr1)
	lenArr2 := len(arr2)
	lenArr := lenArr1 + lenArr2
	arr := make([]int, lenArr)
	for k, _ := range arr {
		if lenArr1 > k {
			arr[k] = arr1[k]
		} else {
			arr[k] = arr2[lenArr-k-1]
		}

	}
	arr = _sort.QuickSortAsc(arr)
	return arr
}
