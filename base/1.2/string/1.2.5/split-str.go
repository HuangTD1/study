package main

import (
	"fmt"
	"strings"
)

func main() {
	// ,  |/ 都是分隔符是或者关系
	fn := func(c rune) bool {
		return strings.ContainsRune(",|/", c) // 定义分隔符
	}
	str := strings.FieldsFunc("Python,Jquery|JavaScript,Go,C++/c", fn)
	fmt.Println(str)

}
