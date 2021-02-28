package base

/* 插入排序*/

// Asc:降序
func Asc(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		j := i
		for j > 0 && arr[j-1] > temp {
			arr[j] = arr[j-1]
			j--
		}
		if j != i {
			arr[j] = temp
		}
	}
	return arr
}

// Desc：升序
func InsertSortDesc(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		j := i

		for j > 0 && arr[j-1] < temp {
			arr[j] = arr[j-1]
			j--
		}
		if j != i {
			arr[j] = temp
		}
	}
	return arr
}
