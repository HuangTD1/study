package main

import (
	"fmt"
	"time"
)

func Hello(c chan string) {
	name := <-c //获取数据
	fmt.Println("Hello", name)
}

func main() {
	// 声明一个字符串 通道
	ch := make(chan string)
	// 开启一个goroutine
	go Hello(ch)
	// 发送数据到通道
	ch <- "World"
	//关闭通道
	close(ch)
	time.Sleep(5 * time.Second) //主进程结束的话，协程也会结束
}
