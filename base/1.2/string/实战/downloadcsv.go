package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// 生成csv 文件
func GenerateCsv(filePath string, data [][]string) {
	fp, err := os.Create(filePath) //创建文件
	if err != nil {
		log.Fatalf("创建文件["+filePath+"]句柄失败, % v", err)
		return
	}
	defer fp.Close()

	fp.WriteString("\xEF\xBB\xBF") // 写入utf-8 BOM
	w := csv.NewWriter(fp)
	w.WriteAll(data)
	w.Flush()

}

// 编写处理器函数
func Welcome(w http.ResponseWriter, r *http.Request) {
	// 定义下载文件名字
	filename := "exportUsers.csv"
	// 定义二维数组数据
	column := [][]string{
		{"手机号", "用户UID", "Email", "用户名"},
		{"188888888888", "1", "lisi@163.com", "LiSi"},
		{"199999999999", "2", "zhangsan@163.com", "ZhangSan"},
		{"166666666666", "3", "wangwu@163.com", "WangWu"},
	}
	// 导出
	GenerateCsv(filename, column)
	// 读取文件
	file, err := os.Open(filename)
	content, err := ioutil.ReadAll(file)
	// 下载文件
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("content-disposition", "attachment;filename=\""+filename+"\"")
	if err != nil {
		fmt.Println("Read file errr:", err.Error())
	} else {
		w.Write(content)
	}
}
func main() {
	http.HandleFunc("/down", Welcome)
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		panic(err)
	}
}
