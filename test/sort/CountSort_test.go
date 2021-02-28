package sort_test

import (
	"fmt"
	"testing"

	"github.com/frzw/algorithm-go/base/sort"
)

func TestCountSortAsc(t *testing.T) {
	arr := []int{1, 3, 2, 1, 0}
	arr = sort.CountSortAsc(arr)
	fmt.Println("res", arr)
}

func TestCountSortDesc(t *testing.T) {
	arr := []int{1, 3, 2, 1, 0, 11, 33, 22}
	arr = sort.CountSortDesc(arr)
	fmt.Println("res", arr)
}
