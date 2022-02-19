### 中文字符串截取
```go
package main

import (
    "fmt"
    "unicode/utf8"
)
func main() {
    //声明1个字符串
    str := "在go中可以愉快地玩耍"
    //打印字符串长度
    fmt.Println(utf8.RuneCountInString(str))
    //打印字节长度
    fmt.Println(len(str))
    //获取前10个字符串
    str1 := str[0:9]
    //打印
    fmt.Println(str1)
    //将字符串转为[]rune 类型
    nameRune := []rune(str)
    // 打印转换后长度
    fmt.Println(len(nameRune))
    // 获取截取后的前10个字符串
    fmt.Println("string =", string(nameRune[0:9]))
}
```