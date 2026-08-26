package main

import (
	"fmt"
	"time"
)

/*
	go语言的协程叫做Goroutine
	Goroutine是由Go程序运行时调度和管理，Go程序会智能地将Goroutine中的任务合理地分配给每个CPU。
	创建Goroutine的成本很小，每个Goroutine的堆栈只有几kb,且堆栈可以根据应用程序的需要增长和收缩

	Goroutine属于抢占式任务处理，和现有的多线程和多进程任务处理非常类似
	应用程序对CPU的控制最终由操作系统来管理，如果操作系统发现一个应用程序占用CPU，那么用户有权终止这个任务
*/
/*
	普通函数创建goroutine
	在函数或方法前面加上关键字go,将会同时运行一个新的Goroutine
	使用go关键字创建goroutine时，被调用的函数往往没有返回值，如果有返回值也会被忽略
	如果需要在goroutine中返回数据，必须使用channel,通过channel把数据从goroutine中作为返回值传出
	Go程序执行过程中是：创建和启动主Goroutine，初始化操作，执行main函数，当main函数结束，主goroutine随之结束
*/
// Goroutine案例1
func hello() {
	fmt.Println("hello world goroutine")
}
func example1() {
	go hello() //不会执行
	fmt.Println("main function")
	// go hello() 的作用只是”启动一个 Goroutine”，不会等待它执行完成。
	// 只要主 Goroutine（通常是 main）结束，整个程序就退出，其他所有 Goroutine 都会被终止。
}

// Goroutine案例2
func example2() {
	go hello()
	time.Sleep(50 * time.Microsecond)
	fmt.Println("main function")
	/*
	   上面这个程序中，已经调用了时间包的Sleep方法，程序会在运行过程中“睡觉”，这种情况下
	   main的goroutine被用来睡觉50毫秒。现在go hello()有足够的时间在main Goroutine终止之前执行
	   这个程序首先打印“hello world gotoutine",等待50毫秒，然后打印“main function”
	*/
}

// 案例3
func example3() {
	go running()
	var input string
	fmt.Scanln(&input)
}
func running() {
	var times int
	for {
		times++
		fmt.Println("tick", times)
		time.Sleep(time.Second)
	}
	// 控制台不断输出tick,同时还可以接收用户输入，两个环节同时进行
}

/*
匿名函数创建gotoutine
go关键字后也可以是匿名函数或闭包
*/
func example4() {
	go func() {
		var times int
		for {
			times++
			fmt.Println("tick", times)
			time.Sleep(time.Second)
		}
	}()
	var input int
	fmt.Scanln(&input)
}

// 启动多个Goroutine
func example5() {
	go printNum()
	go printLetter()
	time.Sleep(3 * time.Second)
	fmt.Println("\nmain over....")
}
func printNum() {
	for i := 1; i < 5; i++ {
		time.Sleep(250 * time.Microsecond)
		fmt.Printf("%d", i)
	}
}
func printLetter() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Microsecond)
		fmt.Printf("%c", i)
	}
} //多个goroutine随机调度，打印的结果是数字与字母交叉输出(无序)
