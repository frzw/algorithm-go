package sort

/* 计数排序 */

// Asc：升序
func CountSortAsc(arr []int) []int {
	countArr, arrIndex := getCountArr(arr)
	for _, v := range arr {
		countArr[v] += 1
	}

	for i := 0; i < len(countArr); i++ {
		for countArr[i] > 0 {
			arr[arrIndex] = i
			arrIndex++
			countArr[i]--
		}
	}
	return arr
}

// Desc：降序
func CountSortDesc(arr []int) []int {
	countArr, arrIndex := getCountArr(arr)
	for _, v := range arr {
		countArr[v] += 1
	}

	for i := len(countArr) - 1; i >= 0; i-- {
		for countArr[i] > 0 {
			arr[arrIndex] = i
			arrIndex++
			countArr[i]--
		}
	}
	return arr
}

func getCountArr(arr []int) ([]int, int) {
	countIndex := 0
	maxValue := getMaxValue(arr)
	minValue := getMinValue(arr)
	countArr := make([]int, maxValue-minValue+1)
	return countArr, countIndex
}
