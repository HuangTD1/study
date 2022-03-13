package main

import (
	"fmt"
	"sort"
)

func main() {
	// create map
	m := make(map[int]string)
	m[0] = "I"
	m[2] = "Go"
	m[1] = "Love"
	// 将键值按排序顺序存储切片中
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	// To perform theopertion you want
	for _, k := range keys {
		fmt.Println(k, m[k])
	}
}
