package main

import (
	"fmt"

	base "github.com/frzw/algorithm-go/base"
)

func main() {
	arr := []int{4, 2}
	res := base.ShellSortAsc(arr)
	fmt.Println(res)
}
