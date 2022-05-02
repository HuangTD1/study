#### golang `interface`
1. 接口样式
```go
type 接口名称 interface {
	method1 (参数列表) 返回值列表
	method2 (参数列表) 返回值列表
	........
        methodn (参数列表) 返回值列表
}
```
2. 在go 语言中接口本身不能创建实例，但是可以指向一个实现了该接口的自定义类型的变量来实现
```go
type Programmer struct {
	Name string
}

func (stu Programmer) Write() {
	fmt.Println("Programmer Write")
}

type SkillInterface interface {
	Write()
}

func main() {
    var pro Programmer //结构体实现了Write()方法，实现了SillInterface 接口
	var a SkillInterface = pro
	a.Write()
}
```

3. 接口赋值

(1) 将实现接口的对象实例赋值给接口
```go
type Number int

func (x Number) Equal(i Number) bool {
	return x == i
}
func (x Number) LessThan(i Number) bool {
    return x < i
}
func (x Number) MoreThan(i Number) bool {
    return x > i
}
func (x * Number) Multiple(i Number) {
	* x = * x * i
}
func (x * Number) Divide(i Number) {
* x = * x * i
}

type NumberI interface {
	Equal(i Number) bool
	LessThan(i Number) bool
	MoreThan(i Number) bool
	Multiple(i Number)
	Divde(i Number)
}
```
(2) 接口组合
```go
// 接口1
type Interface1 interface{
	Write(p []byte) (n int, err error)
}
// 接口2
type Interface2 interface{
    Close() error	
}
// 接口组合
type InterfaceCombine interface{
    Interface1
	Interface2
}
```

#### 接口常见应用
1. 类型推断 `interface1.go`
2. 实现多态功能 `interface2.go`