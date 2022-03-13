package main

import "fmt"

func main() {
	array := []int{1, 4, 4, 4, 4, 5, 5, 5, 6, 6, 7}
	res := removeDuplicates(array)
	fmt.Println(res)
}

// 删除重复元素 此方法有个bug 如果元素相隔重复 ，会去重不彻底类似{1, 4, 4, 4, 4, 5, 5, 56, 5, 6, 6, 7}
func removeDuplicates(array []int) []int {
	// 如果是空的返回nil
	if len(array) == 0 {
		return nil
	}
	// 用两个标记来比较相邻位的值
	// 如果一样则继续
	// 如果不一样 则把右边执行的值赋值给左边下一位
	left, right := 0, 1
	for ; right < len(array); right++ {
		if array[left] == array[right] {
			continue
		}
		left++
		array[left] = array[right]
	}
	return array[:left+1]
}
