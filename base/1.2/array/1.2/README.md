##### 数组删除某个元素
```go
// 初始化1个切片
sed := []string{"i", "Lova", "go", "advanced", "programming"}

// 指定删除 位置
index := 2

// 输出删除位置之前和之后的元素
fmt.Println(sed[:index], sed[index + 1:]) // 结果[i Lova] [advanced programming]
// sed[index +1:]  表示将整个后半部分添加到前段中

// 将删除部分前后的要元素拼接起来
sed = append(sed[:index], sed[index+1:]...)

// 输出连接后的切片
fmt.Println(sed)
// 结果

[i Lova advanced programming]
```


##### 将数组转换为字符串
1. 将数组每1个元素转化为字符串
```go
array := make([]string, 5)
array[0] = "ILovaGO"
string := array[0] // 将数组的1个元素赋值给字符串
fmt.Printf("======>:%s\n", string)
```
2. 将数组转换为字符串
```go
func arrayToString(arr []string) string {
    var result string
    for _, i := range arr { // 遍历所有元素追加
        result + = i
    }
    return result
}
```
3. 检查某个值是否在数组中 示例:`array1.go`
4. 查找1个元素在数组中的位置 示例:`array2.go`
5. 查找数组中最大值和最小值 示例:`array3.go`
6. 随机打乱数组 示例:`array4.go`
7. 删除数组中重复元素 示例:`array5.go`
