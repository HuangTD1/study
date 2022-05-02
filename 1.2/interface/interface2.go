package main

import "fmt"

// Message 是一个定义了通知类行为的接口
type Message interface {
	sending()
}

// Programmer 定义Programmer及Programmer.notify 方法
type Programmer struct {
	name  string
	phone string
}

func (u *Programmer) sending() {
	fmt.Printf("Sending Programmer phone to %s<%s>\n", u.name, u.phone)
}

// Student 定义student及student.message 方法
type Student struct {
	name  string
	phone string
}

func (s *Student) sending() {
	fmt.Printf("Sending Student phone to %s<%s>\n", s.name, s.phone)
}

//sendMessage接受一个实现了message接口的值并发送通知
func sendMessage(n Message) {
	n.sending()
}
func main() {
	// 创建一个Programmer值传递给sendMessage
	bill := Programmer{"Jack", "jack@gmail.com"}
	sendMessage(&bill)
	// 穿件一个Student值传递给sendMessage
	lisa := Student{name: "wade", phone: "wade@gmail.com"}
	sendMessage(&lisa)
}
