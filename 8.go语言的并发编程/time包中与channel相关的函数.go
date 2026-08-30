package main

import (
	"fmt"
	"time"
)

/*
	Timer结构体
	计时器类型表示单个事件。当计时器过期时，当前时间将被发送到c上(c是一个只读channel<-chan.Time
	该channell中放入的是Timer结构体），除非计时器是After()创建的。
	计时器必须使用NewTimer()或After()创建
	New Timer()函数：
	Newtimer()创建一个新的计时器，它会在至少持续时间d之后将当前时间发送到其channel上
*/
//NewTimer()使用案例
func Timecount() {
	//创建计时器
	timer1 := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())
	data := <-timer1.C
	fmt.Printf("%T\n", timer1.C)
	fmt.Printf("%T\n", data)
	fmt.Println(data)
}
