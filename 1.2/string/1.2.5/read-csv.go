package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./test.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// 返回reader 结构体
	reader := csv.NewReader(file)
	//reader.FieldsPerRecord = 1
	// 设置分隔符
	reader.Comma = ';'
	for {
		record, e := reader.Read()
		if e != nil {
			fmt.Println(e)
			break
		}
		fmt.Println(record)
	}

}
