package base

/* 归并排序 */

// Asc：升序
func MergeSortAsc(arr []int) []int {
	sortBool := true
	length := len(arr)
	if length < 2 {
		return arr
	}
	middle := length / 2
	left := arr[:middle]
	right := arr[middle:]
	return merge(MergeSortAsc(left), MergeSortAsc(right), sortBool)
}

// Desc：降序
func MergeSortDesc(arr []int) []int {
	sortBool := false
	length := len(arr)
	if length < 2 {
		return arr
	}
	middle := length / 2
	left := arr[:middle]
	right := arr[middle:]
	return merge(MergeSortDesc(left), MergeSortDesc(right), sortBool)
}

// 归并
func merge(left, right []int, sortBool bool) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] >= right[0] {
			if sortBool {
				result = append(result, right[0])
				right = right[1:]
			} else {
				result = append(result, left[0])
				left = left[1:]
			}
		} else {
			if sortBool {
				result = append(result, left[0])
				left = left[1:]
			} else {
				result = append(result, right[0])
				right = right[1:]
			}
		}
	}

	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}

	return result
}
