package sort

import (
	"fmt"
	"testing"

	_sort "github.com/frzw/algorithm-go/base/sort"
)

func TestBubbleSortAsc(t *testing.T) {
	arr := []int{4, 32, 34, 2, 1, 3, 55, 88, 33, 12}
	arr = _sort.BubbleSortAsc(arr)
	fmt.Println(arr, 10/2, len(arr))
}

func TestBubbleSortDesc(t *testing.T) {
	arr := []int{4, 32, 34, 2, 1, 3, 55, 88, 33, 12}
	arr = _sort.BubbleSortDesc(arr)
	fmt.Println(arr)
}
