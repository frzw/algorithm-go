package sort

/* 选择排序 */

// ASC：升序
func SelectSortAsc(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		swap(arr, min, i)
	}
	return arr
}

// DESC：降序
func SelectSortDesc(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		max := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > arr[max] {
				max = j
			}
		}
		swap(arr, max, i)
	}
	return arr
}
