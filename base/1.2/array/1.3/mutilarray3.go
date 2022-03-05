package main

import "fmt"

func main() {
	// 二维数组 5行2列
	var a = [5][2]int{{0, 1}, {2, 3}, {4, 5}, {6, 7}, {8, 9}}
	var i, j int
	//输出数组元素
	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[% d][% d] = % d\n", i, j, a[i][j])
		}
	}
}
