package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建玩家，设置玩家速度
	var p = games.NewPlayer(0.6)
	fmt.Println(p.Speed)
	// 设置玩家目标位置
	p.MoveTo(games.Vector{3, 9})
	p.CurrentVector = games.Vector{9, 13}
	fmt.Println(p.TargetVector)
	for !p.IsArrived() {
		// 更新位置
		p.Update()
		// 打印位置
		fmt.Println(p.Position())
		// 1s 更新1次
		time.Sleep(time.Second)
	}
	fmt.Println("到达目的地结束")
}
