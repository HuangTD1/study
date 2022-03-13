#### 检查1个key是否在 map中
```go
if _, ok := map[key]; ok {
    // 逻辑代码
}
```

#### json map 转化
```go
jsonStr := `
{
    "name": "lisi",
    "goodat": "Go programming"
}
`
var mapResult map[string]interface{}
err := json.Unmarshal([]byte(jsonStr), &mapResult)
if err != nil {
    fmt.Println("Json TO Map demo errr:%s", err)
}
fmt.Println("mapResult")
```
