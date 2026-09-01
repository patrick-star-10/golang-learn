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
			mutex.Unlock() //票没了，但是前面gotoutine进来的时候上锁了，如果不unlock,程序就会阻塞，报错
			break
		}
		//解锁
		mutex.Unlock()
	}
}

/*
	读写互斥锁
	RWMutex是读写互斥锁。该锁可以同时被多个读取者持有或被唯一一个写入者持有
	RWMutex可以创建为其结构体的字段；零值为解锁状态，和goroutine也无关，可以由不同的goroutine加读取锁/写入锁
	读写锁的使用中，写操作都是互斥的，读和写诗互斥的，读和读不互斥
	该规则可以理解为，多个goroutine可以同时读取数据，但是只能一个goroutine写入数据
	具体方法：
	func(rw*RWMutex)Lock()
	Lock()方法将rw锁定义为写入状态，禁止其他goroutine读取或写入
	func(rw*RWMutex)Unlock()
	Unlock()方法解除rw的写入锁，如果rw未加写入锁会导致运行错误
	func(rw*RWMutex)RLock()
	Rlock()方法将rw锁定为读取状态，禁止其他goroutine写入，但不禁止读取
	func(rw*RWMutex)RUnlock()
	RUnlock()方式解除rw的读取锁，如果rw未加读取锁会导致运行时错误
	func(rw*RWMutex)RLocker() Locker
	Rlocker()方法返回一个读写锁,通过调用rw.Rlock()和rw.RUnlock()实现了Locker接口
*/
//读写锁使用案例
func rwMutex() {
	var rwm sync.RWMutex
	for i := 1; i <= 3; i++ {
		go func(i int) {
			fmt.Printf("goroutine%d,尝试读锁定。\n", i)
			rwm.RLock()
			fmt.Printf("goroutine%d,已经读锁定了\n", i)
			time.Sleep(5 * time.Second)
			fmt.Printf("goroutine %d,读解锁。。\n", i)
			rwm.RUnlock()
		}(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("main..尝试写锁定")
	rwm.Lock()
	fmt.Println("main。。已经写锁定了。。")
	rwm.Unlock()
	fmt.Println("main..写解锁。。")
	// 只要还有读锁 RLock() 没有被 RUnlock() 释放，写锁 Lock() 就无法获得，会阻塞等待
	// 如果写锁已经被某个 goroutine 持有，那么其他 goroutine 想 RLock() 读取，也会阻塞，直到写锁释放。
}
