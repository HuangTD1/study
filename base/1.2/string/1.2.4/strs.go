package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// 转换为大写
	fmt.Println(strings.ToUpper("hello world"))
	// 转换小写
	fmt.Println(strings.ToLower("HELLO WORLD"))
	/*
		使用regexp 中的MustCompile 和replaceAllString 来处理实现
		区分大小写替换字符串
	*/
	test := "I,Love,GoLang"
	str := test
	keywordSlice := strings.Split(test, ",")
	for _, v := range keywordSlice {
		reg := regexp.MustCompile("(?i)" + v)
		// 大写变小写
		str = reg.ReplaceAllString(str, strings.ToLower(v))
		fmt.Println(str)
		fmt.Println("=================")
		// 小写变大写
		str = reg.ReplaceAllString(str, strings.ToUpper(v))
		fmt.Println(str)
	}
	// 去除首尾空格
	strs := " Go world "
	str1 := strings.TrimSpace(strs)
	str2 := strings.Trim(strs, " ")
	fmt.Println(strs)
	fmt.Println(str1)
	fmt.Println(str2)
}
