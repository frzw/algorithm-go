package sort_test

import (
	"fmt"
	"testing"

	_sort "github.com/frzw/algorithm-go/base/sort"
)

func TestShellSortAsc(t *testing.T) {
	arr := []int{4, 32, 34, 2, 1, 3, 55, 88, 33, 12}
	arr = _sort.ShellSortAsc(arr)
	fmt.Println(arr)
}

func TestShellSortDesc(t *testing.T) {
	arr := []int{4, 32, 34, 2, 1, 3, 55, 88, 33, 12}
	arr = _sort.ShellSortDesc(arr)
	fmt.Println(arr)
}
