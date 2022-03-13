package main

import "fmt"

func main() {
	// 创建1个数组
	array := []int{1, 2, 3, 4, 5}
	// 迭代每1个元素并打印
	for k, v := range array {
		fmt.Printf("index:% d value: % d element-address: % x \n", k, v, &array[k])
	}
	fmt.Println("\n 使用匿名变量(下划线)来忽略索引值:")
	for _, v := range array {
		fmt.Printf("value: %d \n", v)
	}
	fmt.Println("\n 使用for 循环对切片进行迭代")
	for i := 0; i < len(array); i++ {
		fmt.Printf(" value: %d \n", array[i])
	}
}
