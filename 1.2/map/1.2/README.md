#### map排序
1. map简单排序`map-sort1.go` 在go语言map 中，进行for 循环时遍历的键，值是使用同1块地址这个地址是临时分配的，虽然地址没有变化但是内容一直在变化，遍历的顺序都是随机的
2. map类型切片 ，必须使用两次`make()`函数 第一次分配，第二次分配切片中每个map的元素
```go
// 创建map
slice := make([]map[int]int,3)
for i := range slice {
    slice[i] = 88
    slice[i] = 66
}
fmt.Printf("Value of slice: %s\n", slice)
// Value of sliece: [map[1:88 2:66] map[1:88,2:66] map[1:88,2:66]]
```

3. map 高级技巧
* 多维map声明遍历
```go
var m map[keyType_1]map[keyType_2]...map[keyType_n]ValueType
// map 如果没有初始化默认值为nil 并且不能对其进行赋值
// 定义1个map
var m map[string]int
//初始化方法1
m := make(map[string]int)
//方法2
var m map[string]int = map[string]int{"barray":12, "tony":22}
// 或者初始化1个空map
m = map[string]int{}
```
* map 多层嵌套的使用和遍历`map1.go`

#### map 排序技巧
1. 按key 排序`key-sort-map.go`
2. 按value 排序`value-sort-map.go`
3. 按map字符出现频率排序`map-str-count.go`