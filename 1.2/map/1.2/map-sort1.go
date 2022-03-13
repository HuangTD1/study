package main

import (
	"fmt"
	"sort"
)

func main() {
	/* 这个例子因为key 和value 都是int 类型所以b 切片只对value
	排序了 key 的值和切边的索引一样
	*/
	var arr map[int]int
	arr = make(map[int]int, 5)
	arr[0] = 88
	arr[1] = 66
	arr[2] = 99
	// 定义1个切片
	var b []int
	fmt.Println("排序前的值")
	// map 是无序的
	for k, v := range arr {
		fmt.Println(k, v)
		b = append(b, v)
	}
	fmt.Println(b)
	// 调用sort 排序方法
	sort.Ints(b)
	fmt.Println("排序后的值如下")
	for k, v := range b {
		fmt.Println(k, v)
	}
	fmt.Println(arr)
}
