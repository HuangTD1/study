package main

import (
	"encoding/csv"
	"os"
)

/*
go 语言提供 encoding/csv 来处理csv 数据 生成csv 数据顺序如下
os.Create() 创建csv 文件
调用文件对象*File的WriteString()方法设置写入文件的字符串类型
再次调用encoding/csv 中的NewWriter()函数来实例化并返回*Writer 对象
调用*Writer中提供的Write()方法写入数据到csv 中
*/
func main() {
	// 创建1个csv 文件
	f, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") // 吸入UTF-8 BOM
	//NeWriter() 函数返回1个Writer 对象
	w := csv.NewWriter(f)
	// 调用Write 方法写入
	w.Write([]string{"学号", "姓名", "分数", "性别"})
	w.Write([]string{"1", "李四", "50", "男"})
	w.Write([]string{"2", "王五", "70", "男"})
	w.Write([]string{"3", "张三", "80", "男"})
	w.Write([]string{"4", "赵敏", "100", "女"})
	w.Write([]string{"5", "张无忌", "90", "男"})
	w.Write([]string{"6", "杨蓉", "40", "女"})
	w.Flush()
}
