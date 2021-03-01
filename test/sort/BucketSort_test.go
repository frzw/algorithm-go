package sort_test

import (
	"fmt"
	"testing"

	_sort "github.com/frzw/algorithm-go/base/sort"
)

func TestBucketSortAsc(t *testing.T) {
	arr := []int{4, 32, 34, 2, 1, 3, 55, 88, 33, 12}
	arr = _sort.BucketSortAsc(arr)
	fmt.Println(arr)
}
