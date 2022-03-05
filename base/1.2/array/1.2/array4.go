package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() { // 初始化方法，程序运行时会先加载，init 可以有多个，以init 结尾即可
	rand.Seed(time.Now().Unix())

}
func main() {
	str := []string{ // 初始化1个数组
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	}
	a, _ := RandomInt(str, 6)
	fmt.Println(a)
}

// 生成随机验证码
func RandomInt(strings []string, length int) (string, error) {
	if len(strings) <= 0 {
		return "", errors.New("字符串长度不能小于0")
	}
	if length <= 0 || len(strings) <= length {
		return "", errors.New("参数长度超出数组长度")
	}
	for i := len(strings) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		strings[i], strings[num] = strings[num], strings[i]
	}
	str := ""
	for i := 0; i < length; i++ {
		str += strings[i]
	}
	return str, nil
}
