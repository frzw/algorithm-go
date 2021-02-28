package sort

/* 堆排序 */

// Asc：升序
func HeapSortAsc(arr []int) []int {
	arrLen := len(arr)
	sortBool := true
	buildHeap(arr, arrLen, sortBool)

	for i := arrLen - 1; i >= 0; i-- {
		swap(arr, i, 0)
		arrLen--
		refactorHeap(arr, 0, arrLen, sortBool)
	}
	return arr
}

// Desc：降序
func HeapSortDesc(arr []int) []int {
	arrLen := len(arr)
	sortBool := false
	buildHeap(arr, arrLen, sortBool)

	for i := arrLen - 1; i >= 0; i-- {
		swap(arr, i, 0)
		arrLen--
		refactorHeap(arr, 0, arrLen, sortBool)
	}
	return arr
}

func buildHeap(arr []int, arrLen int, sortBool bool) {
	for i := arrLen / 2; i >= 0; i-- {
		refactorHeap(arr, i, arrLen, sortBool)
	}
}

func refactorHeap(arr []int, i, arrLen int, sortBool bool) {
	left := 2*i + 1
	right := 2*i + 2
	mid := i
	if sortBool {
		if left < arrLen && arr[left] > arr[mid] {
			mid = left
		}
		if right < arrLen && arr[right] > arr[mid] {
			mid = right
		}
	}

	if !sortBool {
		if left < arrLen && arr[left] < arr[mid] {
			mid = left
		}
		if right < arrLen && arr[right] < arr[mid] {
			mid = right
		}
	}

	if mid != i {
		swap(arr, mid, i)
		refactorHeap(arr, mid, arrLen, sortBool)
	}
}
