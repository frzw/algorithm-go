package base

/* 快速排序 */
func QuickSort(arr []int) []int {
	return _quickSort(arr, 0, len(arr)-1)
}

func _quickSort(arr []int, left, right int) []int {
	if left < right {
		partitionIndex := partition(arr, left, right)
		_quickSort(arr, left, partitionIndex-1)
		_quickSort(arr, partitionIndex+1, right)
	}
	return arr
}

func partition(arr []int, left, right int) int {
	piovt := left
	index := piovt + 1

	for i := index; i <= right; i++ {
		if arr[i] < arr[piovt] {
			swap(arr, i, index)
			index++
		}
	}
	swap(arr, piovt, index-1)
	return index - 1
}
