package sort

/* 计数排序 */

// Asc：升序
func CountSortAsc(arr []int) []int {
	maxValue := getMaxValue(arr)
	minValue := getMinValue(arr)
	countArr := make([]int, maxValue-minValue+1)
	arrIndex := 0
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
	maxValue := getMaxValue(arr)
	countArr := make([]int, maxValue+1)
	arrIndex := 0
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

func getMaxValue(arr []int) int {
	maxValue := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxValue {
			maxValue = arr[i]
		}
	}
	return maxValue
}

func getMinValue(arr []int) int {
	minValue := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < minValue {
			minValue = arr[i]
		}
	}
	return minValue
}
