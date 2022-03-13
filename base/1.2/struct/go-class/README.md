### Go 语言面向对象编程
* 封装: 隐藏对象的属性和实现细节，仅对外提供访问方式
* 继承: 使得子类具有父类的属性和方法，或者重新定义，追加属性方法等
* 多态: 不同对象中同种行为的不同实现方式

1. 封装
(1) 属性
*在go语言中可以使用结构体对属性进行封装，结构体就像是类的一种简化方式例如定义一个三角形，每个三角形都有底和高，例子:*
```go
type Triangle struct {
    Bottom float32
    Height float32
}
```
(2) 方法
*既然有了类，那类的方法呢？ go 语言中也有方法Methods 方法是作用在接受者Receiver上的一个函数，接受者是某种类型的变量，一次方法是一种特殊类型的函数*
```go
func (recv recv_type) methodName (parameter_list) (return_value_list){...}
```
```go
//例子定义1个圆，在定义个计算面积方法:
package main

import "fmt"
// 圆形结构体
type Circle struct {
    Radius float32
}
// 计算面积
func (t *Circle) Area() float32 {
    return t.Radius * t.Radius * 3.14
}

func main() {
    r := Circle{10}
    // 调用类方法
    fmt.Println(r.Area())
}
//out
// 314 
```
(3) 访问权限
*在go 语言中没有public prviate protected 这样的关键字，是通过首字母大小写来去方私有和公共的*
*在go 与验证也有实现获取和设置属性的方式*
* 对于设置方法使用Set前缀
* 对于获取方法只是用成员名
```go
//例子
package main
type Student struct {
    name string
    score float32
}

// 获取name
func (s *Student) GetName() string {
    return s.name
}

// 设置name
func (s *Student) SetName(newName string) {
    s.name = newName
} 

func main() {
    s := new(Student)
    s.SetName("Jack")
    fmt.Println(s.GetName())
}
// out
// Jack
```
(2) 继承
*在go语言中没有extends关键字 它使用在结构体中内嵌匿名类型的方法实现继承，例子:*
```go
type Engine interface {
    Run()
    Stop()
}

type Bus struct {
    Engine //包含Engine 类型的匿名字段
}
```
*此时匿名字段Engine 上的方法晋升成为外层类型Bus的方法，可以构建如下例子：*
```go
func (c *Bus) Working() {
    c.Run() // 开动汽车
    c.Stop() // 停止汽车
}
```
(3)多态
*在面向对象编程中，多态的特征是指不同对象中同种行为的不同实现方式，在go语言中可以使用接口实现这个特征*
```go
//正方形结构体
type Square struct {
    sideLen float32
}
//三角形结构体
type Triangle struct {
    Bottom float32
    Height folat32
}
```
*然后计算出这两个几何图形的面积，由于他们计算方式不同，所以需要定义两个不同的Area方法，于是可以定义1个包含Area方法的接口Shape，让square和triangle 都能调用这个额借口里的Area方法*
```go
// 计算三角形面积
runc (t *Triangle) Area() float32 {
    return (t.Bottom * t.Height)/2
}
// 接口 Shape
type Shape interface {
    Area() float32
}

// 计算正方体面积
func (sq *Square) Area() float32 {
    return sq.sideLen * sq.sideLen
}
func main() {
    t := &Triangle{6,8}
    s := &Square{8}
    shapes := []Shape{t,s} // 创建1个Shape 类型的数组
    for n, _ := range shapes { // 迭代数组中每1个元素并调用Area方法
        fmt.Println("图形数据是:", shapes[n])
        fmt.Println("它的面积是:", shapes[n].Area())
    }
}
// out
// 图形数据是: &{6 8}
// 它的面积是: 24
// 图形数据是: &{8}
// 它的面积是: 64
```