package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("字符串拼接")
	str := []string{"1", "2", "test", "Love", "Go"}
	res := strings.Join(str, "-") // 定义拼接符号
	fmt.Println(res)

	// 字符串切割
	fmt.Println("字符串切割")
	s := "I_Love_Go_web"
	rs := strings.Split(s, "_")
	for i := range rs {
		fmt.Println(rs[i])
	}
	rs1 := strings.SplitN(s, "_", 2)
	for i := range rs1 {
		fmt.Println(rs1[i])
	}
	rs2 := strings.SplitAfter(s, "_")
	for i := range rs2 {
		fmt.Println(rs2[i])
	}
	rs3 := strings.SplitAfterN(s, "_", 2)
	for i := range rs3 {
		fmt.Println(rs3[i])
	}
	// 截取字符串
	fmt.Println("截取字符串")
	fns := func(c rune) bool {
		return strings.ContainsRune(",|/", c)
	}
	res4 := strings.TrimFunc("|/Shirdon Liao,/", fns)
	fmt.Println(res4)
}
