##### 结构体介绍
1. 结构体介绍
(1) 结构体是由一系列具有相同类型或不同类型数据构成的集合，结构体是由0个或多个任意类型的值聚合成的实体，每个值都可以成为结构体的成员
结构体成员也可以被称为`字段` 有一下特性
* 字段拥有自己的类型和值
* 字段名称必须唯一
* 字段的类型也可以是结构体 甚至是字段所在结构体的类型
(2) 结构体定义
```go
type 类型名称 struct {
    字段1 类型1
    字段2 类型2
    //... ...
}
```
* 类型名：表示自定义结构体的名称，在同一个包内名字不可重复
* struct{}: 表示结构体类型
* 字段1 字段2 : 表示结构体字段名，结构体中的字段名必须唯一
* 类型1 类型2： 别是结构体各个字段的类型
(3) 访问权限
* 在go语言中函数名的手写字母大小非常重要，它贝莱实现控制对方法的访问权限，当方法的首字母为大写时，这个方法对于包都是公开的，其他包可以随意调用，，当方法的首字母为小写的这个方法是私有的 只能包内访问
(4)结构体排序
* 在go语言中可以使用`sort.Slice()` 对结构体进行排序
```go
package main
import (
    "fmt"
    "sort"
)
type Programmer struct {
    FirstName string
    GoodAt string
}
func main() {
    members := []Programmer{
        {"Jack": "PHP"},
        {"Jane": "JAVA"},
        {"Barry": "GO"},
    }
    fmt.Println("members")
    sort.Slice(members.func(i,j int) bool {
        return members[i].GoodAt < members[j].GoodAt || members[i].FirstName < members[j].FirstName})
    fmt.Println(members)
}
// out
[ {Jack PHP} {Jane JAVA} {Barry GO}]
[ {Barry GO} {Jane JAVA} {Jack PHP}]
```
*sort 不保证排序的稳定性 两个相同的值排序之后相对位置不变排序的稳定性由`sort.Stable()`函数保证*
```go
package main
import (
    "fmt"
    "sort"
)
type Programmer struct {
    Name string
    Age string
}
type ProgrammerSlice []Programmer
func (s ProgrammerSlice) Len() int { return len(s) }
func (s ProgrammerSlice) Swap(i, j int) { s[i],s[j] = s[j],s[i] }
func (s ProgrammerSlice) Less(i,j int) bool { return s[i].Age < s[j].Age }
func main() {
    a := ProgrammerSlice{
        {
            Name: "Barry",
            Age: 30,
        },
        {
            Name: "Jace",
            Age: 22,
        },
        {
            Name: "Jim",
            Age: 18,
        },
    }
    sort.Stable(a)
    fmt.Println(a)
}
// out
[ {Jim 18} {Jack 22} {Barry 30} ]
```

2. 结构体初始化
* go 语言中数组只能存储同一种类型数据，但在结构体重可以为不同项定义不同的类型
```go
package main
import "fmt"
// 定义
type Programmer struct {
    Name string
    GoodAt string
}
func main() {
    // 方法1
    var person Programmer
    person.Name = "Jack"
    person.GoodAt = "JAVA"
    fmt.Println(person)
    // 方法二
    person1 := Programmer{"Jack", "PHP"}
    fmt.Println(person)
    // 方法三 此处(*person2).Name 等同于person2.Name
    var person2 = new(Programmer)
    person2.Name = "Barry"
    (*person2).GoodAt = "Python"
    fmt.Println(*person2)
    // 方法四 此处(*Programmer).Name 等同于person3.Name 
    var person3 = &Programmer{}
    (*person3).Name = "Shirdon"
    (*person3).GoodAt = "Go"
    fmt.Println(*person3)
}
```
3. 结构体嵌套`struct3.go`
4. 结构体字段标签
* 结构体字段标签Tag 是指结构体字段的额外信息，常用于对自担进行说明，在进行json 序列化即对象关系映射时，都会用到结构体字段标签，标签信息都是静态的，无序实例化结构体，可以通过反射获得。标签在结构体字段后边书写，格式由1个或者多个键值对组成，键值对之间使用空格隔开
```go
package main
import (
    "fmt"
    "reflect"
)
type Programmer struct {
    Name string `json:"name" level:"12"` // 标签
}
func main(){
    var pro Programmer = Programmer{}
    // 反射获取标签信息
    typeOfPro := reflect.TypeOf(pro)
    proFieldName, ok := typeOfPro.FieldByName("Name")
    if ok {
        // 打印标签信息
        fmt.Println(proFieldName.Tag.Get("json"), proFieldName.Tag.Get("level"))
    }
}
//out
//name 12
```