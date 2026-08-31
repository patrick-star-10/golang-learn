package main

import (
	"fmt"
	"math/rand"
	"strconv"
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

/*
	互斥锁
	Mutex是一个互斥锁，可以创建为其他结构体的字段；零值为解锁状态
	Mutex类型的锁和goroutine无关，可以由不同的goroutine加锁和解锁
	Mutex中的方法如下：
	func(m *Mutex)Lock()
	Lock()方法锁住m,如果m已经加锁，则阻塞直到m解锁：func(m *Mutex)Unlock()
	Unlock()方法解锁m,如果m未加锁就会导致运行时错误
*/
//通过互斥锁实现售票
var tickets int = 20
var wg sync.WaitGroup
var mutex sync.Mutex

func MutexLock() {
	wg.Add(4)
	go saleTickets("1号窗口", &wg)
	go saleTickets("2号窗口", &wg)
	go saleTickets("3号窗口", &wg)
	go saleTickets("4号窗口", &wg)
	wg.Wait()
	defer fmt.Println("所有车票售空")
}
func saleTickets(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		mutex.Lock()
		if tickets > 0 {
			time.Sleep(1 * time.Second)
			//获取窗口的编号
			num, _ := strconv.Atoi(name[:1])
			pre := strings.Repeat("-------", num)
			fmt.Println(pre, name, tickets)
			tickets--
		} else {
			fmt.Printf("%s结束售票\n", name)
			mutex.Unlock()
			break
		}
		//解锁
		mutex.Unlock()
	}
}
