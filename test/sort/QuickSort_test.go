package sort_test

import (
	"fmt"
	"testing"

	_sort "github.com/frzw/algorithm-go/base/sort"
)

func TestQuickSortAsc(t *testing.T) {
	arr := []int{3, 2, 1}
	arr = _sort.QuickSortAsc(arr)
	fmt.Println(arr)
}

func TestQuickSortDesc(t *testing.T) {
	arr := []int{4, 32, 34, 2, 1, 3, 55, 88, 33, 12}
	arr = _sort.QuickSortDesc(arr)
	fmt.Println(arr)
}
