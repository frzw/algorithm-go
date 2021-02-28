package sort_test

import (
	"fmt"
	"testing"

	_sort "github.com/frzw/algorithm-go/base/sort"
)

func TestHeapSortAsc(t *testing.T) {
	arr := []int{1, 3, 2, 1, 0}
	arr = _sort.HeapSortAsc(arr)
	fmt.Println("res", arr)
}

func TestHeapSortDesc(t *testing.T) {
	arr := []int{1, 3, 2, 1, 0}
	arr = _sort.HeapSortDesc(arr)
	fmt.Println("res", arr)
}
