package sort

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
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
