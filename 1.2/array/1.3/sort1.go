package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{4, 3, 2, 1, 5, 6}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println("排序后:", a)
}
