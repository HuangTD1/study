##### 数组和切片基础技巧
1. 基础
```Go
slice[开始位置:结束位置]
```
* slice: 表示目标切片对象
* 开始位置: 对应目标切片起始索引
* 结束位置: 对应目标切片技术索引
* 切片的创建方式
```Go
make([]Type,length,capacity)
make([]Type,length)
[]Type{}
[]Type{value1,value2,...,valueN}
```
2. 迭代处理数组
