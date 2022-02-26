#### 读取写入csv 数据
`read-csv.go` 读取csv数据
`write-csv.go` 写入csv数据


#### 按指定函数分割字符串
`split-str.go` 

#### 合并与分割字符串
`split-sr2.go`

1. 合并字符串`Join`
2. 分割字符串`Split() SplitN() SplitAfter() SplitAfterN()` 4个方法
```Go
func Split(s, sep string) []string
// s 为正则分割的字符串 sep 为传入的字符串

func SplitN(s, sep string, n int)[]string
//  s 为正则分割的字符串 sep 为传入的字符串 n 为空值分割的片数n 如果为-1 则不限制如果匹配返回1个字符串切片

func SplitAfter(s, sep string)

func SplitAfterN(s, sep string, n int) []string
```


#### 指定函数截取字符串
```Go
func TrimFunc(s string, f func(rune) bool) string
```