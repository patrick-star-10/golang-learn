package main

import (
	"fmt"
	"time"
)

/*
	select语句的机制有点像switch语句，不同的是，select会随机挑选一个可通信的case来执行
	如果所有case都没有数据到达，则执行default,如果没有default语句，select就会阻塞，直到有case接收数据
*/
//selct分支语句，随机挑选case示例
func Select() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		ch1 <- 100
	}()
	go func() {
		ch2 <- 200
	}()
	time.Sleep(time.Microsecond)
	select {
	case data := <-ch1:
		fmt.Println("ch1中读取数据了: ", data)
	case data := <-ch2:
		fmt.Println("ch2中读取数据了: ", data)
	default:
		fmt.Println("执行了default...")
	}
}

// select的阻塞机制
func selectblock() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		time.Sleep(10 * time.Millisecond)
		data := <-ch1
		fmt.Println("ch1: ", data)
	}()
	go func() {
		time.Sleep(1 * time.Millisecond)
		data := <-ch2
		fmt.Println("ch2: ", data)
	}()
	select {
	case ch1 <- 100:
		close(ch1)
		fmt.Println("ch1写入数据。。")
	case ch2 <- 200:
		close(ch2)
		fmt.Println("ch2写入数据。。。")
	case <-time.After(2 * time.Millisecond):
		fmt.Println("执行延时通道")
	}
	time.Sleep(4 * time.Second)
	fmt.Println("main over")
}

//select 可以同时监听多个 channel 操作和超时事件，程序会阻塞等待，哪个 case 最先满足条件就执行哪个。
// time.After() 常与 select 配合实现超时控制，当规定时间内没有完成 channel 通信时，
// 可执行超时逻辑，这是 Go 并发编程中处理超时最常见的模式
