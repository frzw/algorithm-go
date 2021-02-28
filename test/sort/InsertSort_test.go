package sort_test

import (
	"fmt"
	"testing"

	_sort "github.com/frzw/algorithm-go/base/sort"
)

func TestInsertSortAsc(t *testing.T) {
	arr := []int{4, 32, 34, 2, 1, 3, 55, 88, 33, 12}
	arr = _sort.InsertSortAsc(arr)
	fmt.Println(arr)
}

func TestInsertSortDesc(t *testing.T) {
	arr := []int{4, 32, 34, 2, 1, 3, 55, 88, 33, 12}
	arr = _sort.InsertSortDesc(arr)
	fmt.Println(arr)
}
