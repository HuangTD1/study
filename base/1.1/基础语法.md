### 语言标记
1. go 程序是由`关键字` `标识符` `常量` `字符串` `符号`等多种标记组成

### 行分隔符
1. `fmt.Println("hi,gopher, let's go")`

### 注释
```go
// 单行注释
/*
多行注释
多行注释
*/
```

### 字符串拼接
1. go语言字符串可以用 `"+"` 拼接


### 保留关键字


常量定义关键字 | 类型相关定义关键字 | 函数相关定义关键字 
---- | ----- | ------ 
true,false,iota,nil | int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,uintptr,float32,float64,complex128,complex64,bool,byte,rune,string,error| make,len,cap,new,append,copy,close,delete,comples,real,imag,panic,recover

### 变量
1. Go 语言是静态类型语，因此变量是由明确类型的，编译器也会检查变量类型正确性，声明变量使用`var` 关键字 如下, `var` 是关键字 `name` 是变量名 `type` 是变量类型
```go
var name type
```

2. Go 声明指针类型变量如下,系统会自动赋予该类型的空值
```go
var x,y *int
```

类型| 指针空值
---- | -----
int|0
float|0.0
bool|false
string|""
指针类型|nil

(1) 标准格式
```go
var 变量名 变量类型
```
(2) 批量命名
```go
var (
    userId int
    userName string
    score float32
)
```
(3) 简短格式，简短格式有如下限制
* 只能用来定义变量，同时会显式初始化
* 不能提供数据类型
* 只能用在函数内部，即不能用来声明全局变量
```go
名字 := 表达式
```
### 变量赋值
```go
var name (type) = value

var (
    name1(type1) = value1
    name2(type2) = value2
    name3(type3) = value3
)

var name1,name2,name3 = value1,value2,value3
var (
    userId int = 1
    username string = "Test"
    score float32 = "99.4"
)
```

### 常量
1. 常量的声明方式和变量类似如下
```go
const 常量名 [类型] = 常量值

例子：
const name = "zhangsan"
```
2. 常量生成器`iota`
* 常量声明可以使用常量生成器初始化，用户生成一组规则相似的常量，但不是每一行都初始化下变量，在第一个声明的常量所在的行`iota`被重置为`0` 然后每一个常量行`加1`
```go
type Weekday int
const (
　　Sunday Weekday = iota
　　Monday
　　Tuesday
　　Wednesday
　　Thursday
　　Friday
　　Saturday
)
```

### 运算符

优先级| 分类| 运算符 | 结合性
---- | -----| ------| ------
1|逗号运算符|,| 从左到右
2|赋值运算符|=,+=,-=,*=,/=,%=,>=,<<=,&=,^=,\|=|从右到左
3|逻辑或|\|\||从左到右
4|逻辑与|&&|从左到右
5|按位或| \||从左到右
6|按位异或|^|从左到右
7|按位与|&|从左到右
8|相等/不等|==,!=|从左到右
9|关系运算符|<,<=,>,>=|从左到右
10|位移运算符|<<,>>|从左到右
11|加法/减法|+,-|从左到右
12|乘法/除法/取余|*(乘号),/,%|从左到右
13|单且运算符|!,*(指针),&,++,--,+(正号),-(负号)|从右到左
14|后缀运算符|(),[],->|从左到右


### 流程空值语句
1. `if...else`(分支结构)
```go
age := 12
if age == 12 {
    fmt.Println("ok")
}else {
    fmt.Println("failed")
}
```

2. for 循环
```go
for init; condition;post {}

for condition {}

for {}
```
* init: 一般位赋值表达式,给空值变量赋初始值
* condition: 关系表达式或逻辑表达式，循环空值条件
* post: 一般位赋值表达式，给空值变量增量或减量
* 左大括号 { 必须与for 在同一行
* Go 语言中的for 一样可以使用`continue break` 来终止循环
* 如果循环被`break goto return panic 等语句强制退出，则之后语句都不会执行

3. `for ...range` 循环
* `for...range` 可以遍历`数组 切片 字符串 map 和 通道(channel)`
```go
for key, val := range 复合变量 {
    //....逻辑语句
}
```

4. `switch...case` 语句
* 一分支多值
```go
var user = "zhangsan"
switch user {
    case "zhangsan", "lisi":
    fmt.Println("good name")
}
```
* 分支表达式 `这种情况switch 后不需要加判断变量`
```go
var num int = 12
switch {
    case num > 1 && num < 10:
    fmt.Println(num)
}
```

5. `goto` 语句
* 在Go 语言中goto 语句跳转标签，进行代码间无条件跳转，另外goto在快速调处循环，避免重复退出方面也有一定帮助
```go
func main() {
    for x := 0; x < 5; x ++ {
        for y := 0; y < 5; y++ {
            if y == 2 {
                goto breakTag //跳转标签
            }
        }
    }
    return //手动返回避免执行进入标签
breakTag: // 标签
    fmt.Pringln("done")
}
```

6. `break` 语句
```go
package main

import (
    "fmt"
)
func main() {
    fmt.Println("1")
Exit: //标签
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if i+j > 15 {
                fmt.Print("exit")
                break Exit  //跳到外面去啦，但是不会再进来这个for循环了
            }
        }
    }
    fmt.Println("3")
}
```
7. `continue` 语句
```go
package main

import (
    "fmt"
)

func main() {

    for i := 1; i <= 10; i++ {
        if i < 6 {
            continue     //如果i<6，忽略本次循环，继续下一次迭代
        }
        fmt.Println(i)
    }
}
```
