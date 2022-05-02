package main

// 创建定时通知
import (
	"fmt"
	"time"
)

func TimeLong(d time.Duration) <-chan struct{} {
	ch := make(chan struct{}, 1)
	go func() {
		time.Sleep(d)
		ch <- struct{}{}
	}()
	return ch
}

func main() {
	fmt.Println("hello")
	<-TimeLong(time.Second)
	fmt.Println("1s后继续提示")
	<-TimeLong(time.Second)
	fmt.Println("再过1s后提示")
}
