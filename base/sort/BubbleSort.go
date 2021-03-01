package sort

/* 冒泡排序 */

// ASC:升序-优化外部、内部

func BubbleSortAsc(arr []int) []int {
	flag := 0
	k := len(arr) - 1
	for i := 0; i < len(arr); i++ {
		pos := 0
		for j := 0; j <= k; j++ {
			if j+1 <= k && arr[j] > arr[j+1] {
				swap(arr, j, j+1)
				flag = 1
				pos = j
			}
		}
		k = pos
		if flag == 0 {
			return arr
		}
	}
	return arr
}

// DESC:降序
func BubbleSortDesc(arr []int) []int {
	k := len(arr) - 1
	pos := 0
	for i := 0; i < len(arr); i++ {
		flag := 0
		for j := 0; j <= k; j++ {
			if j+1 <= k && arr[j] < arr[j+1] {
				swap(arr, j, j+1)
				flag = 1
				pos = j
			}
		}
		k = pos
		if flag == 0 {
			return arr
		}
	}
	return arr
}
