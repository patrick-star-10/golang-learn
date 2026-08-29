package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

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
	// <-ch2的作用就是阻塞，等待匿名函数的goroutine运行结束，防止主函数退出导致goroutine提前退出
}

/*
	关闭channel
	发送方如果数据写入完毕，需要关闭channel，用于通知接收方数据传递完毕。
	通常情况是发送方主动关闭channel,接收方通过多重返回值判断channel是否关闭，返回false表示关闭
	往关闭的channel中写入数据会报错，但是可以从关闭的channel中读取数据，返回数据的默认值是false
*/
//channel关闭以后是否可以写入数据案例
func CloseChannel() {
	// channel关闭后是否可以写入数据？
	ch1 := make(chan int)
	go func() {
		ch1 <- 100
		ch1 <- 200
		close(ch1)
		ch1 <- 10 //关闭的channel，无法写入数据
	}()
	data, ok := <-ch1
	fmt.Println("main读取数据: ", data, ok)
	data, ok = <-ch1
	fmt.Println("main读取数据: ", data, ok)
	data, ok = <-ch1
	fmt.Println("main读取数据: ", data, ok)
	data, ok = <-ch1
	fmt.Println("main读取数据: ", data, ok)
	data, ok = <-ch1
	fmt.Println("main读取数据: ", data, ok)
	//由运行结果可知，向已关闭的channel写入数据会导致程序崩溃
}

/*
	缓冲channel
	默认创建的都是非缓冲channel，读写都是即使阻塞。
	缓存channel自带一块缓冲区，可以暂时存储数据，如果缓冲区满了，就会发生阻塞
*/
// 缓存channel案例
func Channel2() {
	//1.非缓冲channel
	ch1 := make(chan int)
	fmt.Println("非缓冲通道", len(ch1), cap(ch1))
	go func() {
		data := <-ch1
		fmt.Println("获得数据: ", data)
	}()
	ch1 <- 100
	time.Sleep(time.Second)
	fmt.Println("赋值ok", "main over...")
	//2.非缓冲通道
	ch2 := make(chan string)
	go sendData2(ch2)
	for data := range ch2 {
		fmt.Println("\t读取数据", data)
	}
	fmt.Println("main over...")
	//3.缓冲通道，缓冲区满了才会阻塞
	ch3 := make(chan string, 6)
	go sendData2(ch3)
	for data := range ch3 {
		fmt.Println("\t读取数据", data)
	}
	fmt.Println("main over...")
}
func sendData2(ch chan string) {
	for i := 1; i <= 3; i++ {
		ch <- fmt.Sprintf("data%d", i)
		fmt.Println("往通道放入数据: ", i)
	}
	defer close(ch)
	// 无缓冲 Channel 能保证数据的发送和接收必须完成交接，但不能保证交接完成后，发送方和接收方各自的 Println 谁先执行。
	// 缓冲 Channel 因为发送不必等待接收，所以发送方常常能连续跑完，看起来更有序，
	// 但这种打印顺序本身依然不是 Channel 保证的
}

// 缓冲channel模拟生产者和消费者
func CHANNEL() {
	//用channel来传递数据，不再需要自己去加锁维护一个全局的阻塞列队
	ch1 := make(chan int)
	ch_bool1 := make(chan bool) //判断结束
	ch_bool2 := make(chan bool) //判断结束
	ch_bool3 := make(chan bool) //判断结束
	rand.Seed(time.Now().UnixNano())
	// 生产者
	go producer(ch1)
	//消费者
	go consumer(1, ch1, ch_bool1)
	go consumer(2, ch1, ch_bool2)
	go consumer(3, ch1, ch_bool3)
	<-ch_bool1
	<-ch_bool2
	<-ch_bool3
	defer fmt.Println("main...over")
}

// 生产者
func producer(ch1 chan int) {
	for i := 1; i <= 10; i++ {
		ch1 <- i
		fmt.Println("生成蛋糕,编号为: ", i)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
	defer close(ch1)
}

// 消费者
func consumer(num int, ch1 chan int, ch chan bool) {
	for data := range ch1 {
		pre := strings.Repeat("--------", num)
		fmt.Printf("%s %d号购买%d号蛋糕 \n", pre, num, data)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
	ch <- true
	defer close(ch)
	/*
	   生产者—消费者模型：使用 channel 作为生产者和消费者之间的数据通道。生产者负责向 channel 发送数据，
	   多个消费者并发地从 channel 接收数据，每条数据只会被一个消费者处理。close(channel)
	   用于通知消费者数据已全部发送完毕，消费者退出后再通知主 goroutine，从而实现多个 goroutine 之间安全、
	   高效的协作，无需使用互斥锁（Mutex）维护共享队列
	*/
}

/*
	单向channel
	channel默认都是双向的，即可读可写。定向channel也叫单向channel，只读，或只写
	只读channel使用方式如下所示：
	make(<- chan type)
	<-chan
	只写channel使用方式如下：
	make(chan<- type)
	chan <- data
	直接创建单向channel没有任何意义。通=通常的做法是创建双向channel，然后以单向channel的
	方式进行函数传递
*/

func SingalChannel() {
	//双向通道
	ch1 := make(chan string)
	go fun1(ch1)
	data := <-ch1
	fmt.Println("main,接收到数据: ", data)
	ch1 <- "区块链"
	ch1 <- "以太坊"
	go fun2(ch1)
	go fun3(ch1)
	time.Sleep(1 * time.Second)
	fmt.Println("main over")

}
func fun1(ch1 chan string) {
	ch1 <- "我是steven老师"
	data := <-ch1
	data2 := <-ch1
	fmt.Println("回应: ", data, data2)
}

// 功能；只有写入数据
func fun2(ch1 chan<- string) {
	//只能写入
	ch1 <- "how are you"
}

// 功能：只有读取数据
func fun3(ch1 <-chan string) {
	data := <-ch1
	fmt.Println("只读: ", data)
}

/*
单向 Channel：Go 可以将双向 channel 限制为发送专用（chan<- T）或接收专用（<-chan T）。
发送专用只能写入数据，接收专用只能读取数据，违反限制会在编译时报错。单向 channel 并不会创建新的通道，
而是对同一个 channel 进行权限限制，使函数职责更清晰、代码更安全，避免误操作
*/
