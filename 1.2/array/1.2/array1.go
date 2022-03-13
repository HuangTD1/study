package main

import "fmt"

// 检查字符串是否在切片中

func Exist(target string, array []string) bool {
	for _, element := range array {
		if target == element {
			return true
		}
	}
	return false
}

func main() {
	nameList := []string{"zhangsan", "list"}
	str1 := "zhangsan"
	str2 := "wangwu"
	result := Exist(str1, nameList)
	fmt.Println("zhangsan name 是否在 nameList 中", result)
	result = Exist(str2, nameList)
	fmt.Println("wangwu 是否在 nameList 中", result)
}
