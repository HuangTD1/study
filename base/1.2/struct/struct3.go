package main

import (
	"encoding/json"
	"fmt"
)

// 基础
type Item struct {
	Title string
	URL   string
}

// 定义相应结构体
type Response struct {
	Data struct {
		Children []struct {
			Data Item
		}
	}
}

func main() {
	jsonStr := `{"data": {"children": [{"data": {"title": "This is test response","url": "https://www.baidu.com"}}]}}`
	res := Response{}
	err := json.Unmarshal([]byte(jsonStr), &res)
	fmt.Println(err)
	fmt.Println(res)
}
