package main

import "fmt"

/*
	channel的概述
	channel即go的通道，是协称之间的通信机制。一个channel就是一条通信管道，他可以让一个协程给另一个协程发送数据
	每个channel都需要制定数据类型，即channel可发送数据的类型，如果发送int类型数据，可以写成chan int
	数据发送的方式如同水在管道中的流动
	传统的线程之间可以通过共享内存进行数据交互，不同的线程共享内存的同步问题需要使用锁来解决，这样会导致性能底下
	go语言提倡使用channel的方式来代替共享内存，换而言之，go语言主张通过数据传递来实现共享内存
	而不是通过共享内存来实现数据传递

	创建channel类型
	声明channel类型的语法格式如下：
	{var channel 变量 chan channel 类型}

	chan类型的空值是nil,声明后需要配合make()才能使用
	channel是应引用类型，需要make)()进行创建，语法格式如下
	{channel 示例 := make(chan 数据类型)}

	具体创建语法如下：
	ch1 := make(chan int)
	ch2 := make(chan interface{})
	type Equip struct{}
	ch3 := make(chan *Equip)  //创建一个Equip指针类型的channel,可以存放Equip指针
*/

/*
	使用channel发送数据
	通过channel发送数据需要使用特殊操作符“<-”.将数据通过channel发送的语法格式如下：
	channel 变量 <- 值
	channel发送的值的类型必须与channel的元素类型一致，如果接收方一直没有接收，那么发送操作将持续阻塞
	此时所有的Goroutine,包括main()的Goroutine都处于等待状态
	使用channel要考虑发生死锁(deadlock)的可能。如果goroutine在一个channel上发送数据，
	其他的goroutine应该接收得到数据；如果没有接收，那么程序将在运行的时候出现死锁
	如果goroutine正在等待从channel接收数据，其他一些goroutine将会在该channel上写入数据，如果没有写入，程序将会死锁
*/

/*
	通过channel接收数据
	channel收发操作在不同的两个goroutine间进行，语法格式有四种：
	1.阻塞接收数据
	channel接收同样树勇特殊的操作符"<-",语法格式如下：
	data := <-ch
	执行改语句时channel将会阻塞，直到接收到数据并赋值给data变量

	2.完整写法
	阻塞接收数据的完整写法如下：
	data,ok := <-ch
	data:表示接收到的数据；ok:表示是否接收到数据，通过ok(bool类型)可以判断当前channel是否被关闭

	3.忽略接收数据
	接收任意数据，忽略接收的数据，语法格式如下：
	<-ch
	执行该语句时channel将会阻塞，其目的不在于接收channel中数据，而是为了阻塞goroutine

	4.循环接收数据
	循环接收数据，需要配合使用关闭channel,借助普通for循环和for...range语句循环接收多个元素。
	遍历channel，遍历的结果就是接收到的数据，数据类型就是channel的数据类型
	普通for循环接收channel数据，需要有break循环的条件;for...range会自动判断出channel已关闭，而无需通过判断来终止循环
*/

// channel接收的三种方式
func Channel() {
	ch1 := make(chan string)
	go sendData(ch1)
	// //1.循环接收数据方式1
	// for {
	// 	data := <-ch1
	// 	if data == "" {
	// 		break
	// 	}
	// 	fmt.Println("从通道中读取数据方式1: ", data)
	// 	fmt.Println("-----------------------")
	// }

	//2.循环接收方式2
	// for {
	// 	data, ok := <-ch1
	// 	fmt.Println(ok)
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println("从通道中接收数据方式2: ", data)
	// 	fmt.Println("-----------------------")
	// }

	//3.循环接收数据方式3
	for value := range ch1 {
		fmt.Println("从通道中读取数据方式3: ", value)
	}
}
func sendData(ch1 chan string) {
	defer close(ch1)
	for i := 0; i < 3; i++ {
		ch1 <- fmt.Sprintf("发送数据%d\n", i)
	}
	fmt.Println("发送数据完毕")
	//显示调用close()实现关闭通道
}

/*
	阻塞
	channel默认是阻塞的。当数据被发送到channel时会发生阻塞，知道有其他goroutine从这个channel中读取数据
	当从channel中读取数据时，读取也会被阻塞，直到其他goroutine将数据写入该channel
	这些channel的特性帮助goroutine有效地通信，而不需要使用其他语言中的显式锁或条件变量
*/
//阻塞的基本用法
func block() {
	var ch1 chan int
	ch1 = make(chan int)
	fmt.Printf("%T\n", ch1)
	ch2 := make(chan bool)
	go func() {
		data, ok := <-ch1
		if ok {
			fmt.Println("子goroutine取到数值: ", data)
		}
		ch2 <- true
	}()
	ch1 <- 10
	<-ch2 //阻塞
	fmt.Println("main over...")
}
