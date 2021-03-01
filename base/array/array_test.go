package array_test

import (
	"fmt"
	"testing"

	_array "github.com/frzw/algorithm-go/base/array"
)

func TestCombineTwoOrderedArrays(t *testing.T) {
	arr1 := []int{1, 2, 3}
	arr2 := []int{1, 2, 3, 3}
	arr := _array.CombineTwoOrderedArrays(arr1, arr2)
	fmt.Println(arr)
}
