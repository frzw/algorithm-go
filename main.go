package main

import (
	"fmt"

	base "github.com/frzw/algorithm-go/base"
)

func main() {
	arr := []int{4, 32, 34, 2, 1, 3, 55, 88, 33, 12}
	res := base.BubbleSortDesc(arr)
	fmt.Println(res)
}
