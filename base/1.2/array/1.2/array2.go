package main

import (
	"fmt"
	"reflect"
)

func main() {
	List := make([]int, 6)
	for i := 0; i < 6; i++ {
		List[i] = i + 2
	}
	index := arrayPoition(List, 4)
	fmt.Println(List)
	fmt.Println(index)
}

func arrayPoition(arr interface{}, d interface{}) int {
	array := reflect.ValueOf(arr)
	for i := 0; i < array.Len(); i++ { // 根据数组长度遍历索引
		v := array.Index(i)
		// fmt.Println(v.Interface())
		if v.Interface() == d { // 判断是否等于查找的值 等于返回索引
			return i
		}
	}
	return -1 // 未找到返回-1
}
