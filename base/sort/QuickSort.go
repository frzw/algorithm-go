package sort

import "fmt"

/* 快速排序 */

// Asc：升序
func QuickSortAsc(arr []int) []int {
	return _quickSort(arr, 0, len(arr)-1, true)
}

// Desc：降序
func QuickSortDesc(arr []int) []int {
	return _quickSort(arr, 0, len(arr)-1, false)
}

func _quickSort(arr []int, left, right int, sortBool bool) []int {
	if left < right {
		partitionIndex := partition(arr, left, right, sortBool)
		_quickSort(arr, left, partitionIndex-1, sortBool)
		_quickSort(arr, partitionIndex+1, right, sortBool)
	}
	return arr
}

func partition(arr []int, left, right int, sortBool bool) int {
	piovt := left
	index := piovt + 1

	for i := index; i <= right; i++ {
		if sortBool && arr[i] < arr[piovt] {
			swap(arr, i, index)
			index++
			fmt.Println(index, "-", arr)
		}
		if !sortBool && arr[i] > arr[piovt] {
			swap(arr, i, index)
			index++
		}
	}
	swap(arr, piovt, index-1)
	return index - 1
}
