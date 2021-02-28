package sort_test

import (
	"fmt"
	"testing"

	_sort "github.com/frzw/algorithm-go/base/sort"
)

func TestSelectSortAsc(t *testing.T) {
	arr := []int{4, 32, 34, 2, 1, 3, 55, 88, 33, 12}
	arr = _sort.SelectSortAsc(arr)
	fmt.Println(arr)
}

func TestSelectSortDesc(t *testing.T) {
	arr := []int{4, 32, 34, 2, 1, 3, 55, 88, 33, 12}
	arr = _sort.SelectSortDesc(arr)
	fmt.Println(arr)
}
