package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

/*
	sync包提供了互斥锁。除了Once和WaitGroup类型，其余多数适合低水平的程序
	多数情况下，高水平的同步使用channel通信性能会更优一些。sync包类型的值不应该被复制
	前面的案例中一般使用time.Sleep函数，通过睡眠将主goroutine阻塞至所有goroutine结束
	而更好的做法是使用WaitGroup来实现
*/

// WaitGroup方法使用案例
func waitgroup() {
	var wg sync.WaitGroup
	fmt.Printf("%T\n", wg) //sync.WaitGroup
	fmt.Println(wg)
	wg.Add(3)
	rand.Seed(time.Now().UnixNano())
	go printnum(&wg, 1)
	go printnum(&wg, 2)
	go printnum(&wg, 3)
	wg.Wait() //进入阻塞状态，当计数为0时解除阻塞
	defer fmt.Println("main over...")
}
func printnum(wg *sync.WaitGroup, num int) {
	for i := 1; i <= 3; i++ {
		//在每个goroutine前添加多个制表符方便观看打印结果
		pre := strings.Repeat("\t", num-1)
		fmt.Printf("%s 第%d号子goroutine,%d\n", pre, num, i)
		time.Sleep(time.Second)
	}
	wg.Done() //计数器减1
	// Add(3) 表示我要等 3 个 goroutine；每个 goroutine 完成后 Done() 一次；
	// main 在 Wait() 阻塞，直到 3 次 Done() 把计数器减到 0
}
