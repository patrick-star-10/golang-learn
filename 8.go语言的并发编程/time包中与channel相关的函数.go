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
	/*
	   Timer（计时器）本质上是基于 Channel 实现的。time.NewTimer(d) 创建一个计时器，
	   并返回包含只读 C 通道的 Timer 对象。程序执行到 <-timer.C 时会阻塞等待，直到计时结束，
	   Timer 自动向 C 通道发送一个 time.Time 类型的数据（触发时间），
	   程序收到数据后继续执行。因此，Timer 可以看作是利用 Channel 实现的“时间通知器”
	*/
}

/*
	After()函数
	After()函数相当于NewTimer(d).C
*/
//After()函数案例
func after() {
	//使用After(),返回值<-chan Time,同Timer.C
	ch1 := time.After(5 * time.Second)
	fmt.Println(time.Now())
	data := <-ch1
	fmt.Printf("%T\n", data)
	fmt.Println(data)
}

//time.After() 可以理解为 time.NewTimer(d).C 的简化写法。
// 只需要等待一次就用 After()；如果需要停止、重置或管理计时器，就使用 NewTimer()。
