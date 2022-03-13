package main

import "fmt"

func main() {
	var array = [...]int{1, -2, 66, 33, 121, 4565, 1231, -12312, 999999999999999}
	maxValue := array[0]
	maxValueIndex := 0
	for i := 0; i < len(array); i++ {
		// 比较元素大小，如果发现更大的数进行交换
		if maxValue < array[i] {
			maxValue = array[i]
			maxValueIndex = i
		}
	}
	fmt.Printf("maxValue = % v maxValueIndex = % v\n", maxValue, maxValueIndex)
}
