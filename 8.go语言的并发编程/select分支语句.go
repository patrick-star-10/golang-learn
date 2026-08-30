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
