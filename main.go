package main

import (
	"fmt"

	base "github.com/frzw/algorithm-go/base/sort"
)

func main() {
	arr := []int{4, 32, 34, 2, 1, 3, 55, 88, 33, 12}
	res := base.MergeSortDesc(arr)
	fmt.Println(res)
}
