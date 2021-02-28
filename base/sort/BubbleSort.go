package sort

/* 冒泡排序 */

// ASC:升序
func BubbleSortAsc(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if j+1 < len(arr) && arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
		}
	}
	return arr
}

// DESC:降序
func BubbleSortDesc(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if j+1 < len(arr) && arr[j] < arr[j+1] {
				swap(arr, j, j+1)
			}
		}
	}
	return arr
}
