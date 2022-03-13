package main

import "fmt"

func main() {
	// 定义字符串
	str := "cdef123456789abcdefABC"
	// 调用反转函数
	strRev := Reversal(str)
	// 打印
	fmt.Println(str)
	fmt.Println(strRev)
}

// 反转字符串方法
func Reversal(s string) (revelString string) {
	// 将字符串转为rune 数组
	b := []rune(s)
	// 遍历
	for i := 0; i < len(b)/2; i++ {
		// 交换
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}
	// 转换为字符串类型
	revelString = string(b)
	return
}
