package sort

func HeapSort(arr []int) []int {
	arrLen := len(arr)
	// 构建大顶堆
	buildMaxHeap(arr, arrLen)

	// 交换顶尾元素、剩余元素构建大顶堆
	for i := arrLen - 1; i >= 0; i-- {
		swap(arr, 0, i)
		arrLen -= 1
		heapify(arr, i, arrLen)
	}
	return arr
}

func buildMaxHeap(arr []int, arrLen int) {
	for i := arrLen / 2; i >= 0; i-- {
		heapify(arr, i, arrLen)
	}
}

func heapify(arr []int, i, arrLen int) {
	left := 2*i + 1                                // 左孩子
	right := 2*i + 2                               // 右孩子
	largest := i                                   // 父节点
	if left < arrLen && arr[left] > arr[largest] { // 左孩子>父节点
		largest = left
	}
	if right < arrLen && arr[right] > arr[largest] { // 右孩子>父节点
		largest = right
	}
	if largest != i {
		swap(arr, i, largest)
		heapify(arr, largest, arrLen)
	}
}
