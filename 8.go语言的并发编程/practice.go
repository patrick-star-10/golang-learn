package main

/*
综合练习：限时抢单配送系统

【背景】
模拟一个外卖平台。
平台有 3 个配送员 goroutine，同时等待平台派发订单。
main goroutine 负责产生订单、派发订单以及接收配送结果。

订单结构体：

type Order struct {
	ID   int
	Name string
}

【要求】

1. 创建 orders channel
   - 用于 main 向配送员发送订单。

2. 创建 result channel
   - 配送员完成订单后，通过 result channel 将配送结果发送给 main。
   - 例如：
     "配送员2 完成订单3：炸鸡"

3. 启动 3 个配送员 goroutine
   - 每个配送员不断从 orders channel 获取订单。
   - 推荐使用：
     for order := range orders
   - 收到订单后输出：
     "配送员1 接到订单3：奶茶"

4. 模拟配送
   - 每个订单配送耗时 1~3 秒。
   - 使用 time.Sleep() 模拟配送过程。
   - 配送完成后，将结果发送到 result channel。

5. main 一共生成 10 个订单
   - 可以自己定义订单名称，例如：
     汉堡、炸鸡、奶茶、披萨等。

6. 派单必须有超时机制
   - 每个订单最多等待 2 秒。
   - 必须使用：
     select
     time.After()

   - 如果 2 秒内有配送员接收订单：
     输出：
     "订单5派送成功"

   - 如果 2 秒内没有配送员接收：
     输出：
     "订单5派送超时，取消订单"

7. 10 个订单全部尝试派发完成后：
   close(orders)

   - 让正在：
     for order := range orders
     的配送员知道不会再有新订单，从而退出。

8. 使用 sync.WaitGroup
   - main 必须等待 3 个配送员全部结束。
   - 每个配送员结束时调用 Done()。

9. main 还需要接收 result channel 中的配送结果。

   注意思考：
   如果 result 是无缓冲 channel，
   配送员执行：

       result <- xxx

   时需要有人接收。

   因此不能不加思考地：
       wg.Wait()
   之后才开始接收 result。

   思考如何让：
   - 派发订单
   - 配送员配送
   - 接收配送结果

   正确地并发工作，避免死锁。

10. 所有配送员结束、所有配送结果处理完成后输出：

    "所有订单处理完毕"
    "main over"

--------------------------------------------------

【必须使用的知识】

- goroutine
- channel
- select
- time.After()
- time.Sleep()
- sync.WaitGroup
- close(channel)
- for range channel

--------------------------------------------------

【额外挑战】

增加：

var completed int

用于记录成功完成的订单数量。

3 个配送员完成订单时都执行：

completed++

由于 completed 是多个 goroutine 共同修改的共享数据，
需要使用：

sync.Mutex

保护 completed++。

最终输出：

"成功完成订单数量：X"

--------------------------------------------------

【思考题】

完成程序后尝试回答：

1. 为什么 orders 适合使用 channel？
2. 为什么 completed++ 需要 Mutex？
3. 为什么要 close(orders)？
4. 为什么 for range orders 在 close 后能够退出？
5. 为什么 WaitGroup 不能替代 channel？
6. 如果 result 是无缓冲 channel，为什么可能发生死锁？
7. select + time.After() 是如何实现订单超时的？
8. 为什么 goroutine 的输出顺序每次运行可能不同？
*/
import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID   int
	Name string
}

func Deliver() {
	orders := make(chan Order)
	result := make(chan string)

	var mu sync.Mutex
	var wg sync.WaitGroup
	var wg2 sync.WaitGroup
	var completed int

	wg.Add(3)
	wg2.Add(1)
	for i := 1; i <= 3; i++ {
		go worker(i, orders, result, &wg, &mu, &completed)
	}
	// time.Sleep(100 * time.Millisecond)

	go func() {
		defer wg2.Done()
		for res := range result {
			fmt.Println(res)
		}
	}()

	foods := []string{"汉堡", "炸鸡", "奶茶", "披萨", "薯条", "可乐", "蛋挞", "鸡翅", "冰淇淋", "米线"}

	for i := 1; i <= 10; i++ {
		order := Order{
			ID:   i,
			Name: foods[i-1],
		}
		fmt.Printf("\n=======开始派发订单%d:%s========\n", order.ID, order.Name)

		select {
		case orders <- order:
			fmt.Printf("订单%d派送成功", order.ID)
		case <-time.After(2 * time.Second):
			fmt.Printf("订单%d派送超时,取消订单\n", order.ID)
		}
	}
	close(orders)
	fmt.Println("\n所有订单已派发完毕,关闭订单通道")

	wg.Wait()
	close(result)
	wg2.Wait()
	fmt.Println("所有订单派送结束")
	fmt.Printf("成功完成订单数量:%d\n", completed)
	fmt.Println("main over...")

}
func worker(id int, orders <-chan Order, result chan<- string, wg *sync.WaitGroup, mu *sync.Mutex, completed *int) {
	defer wg.Done()

	for order := range orders {
		fmt.Printf("配送员%d 接到订单%d:%s\n", id, order.ID, order.Name)

		delay := time.Duration(rand.Intn(3)+1) * time.Second
		time.Sleep(delay)

		res := fmt.Sprintf("配送员%d 完成订单%d:%s", id, order.ID, order.Name)
		result <- res

		mu.Lock()
		*completed++
		mu.Unlock()
	}
	fmt.Printf("配送员%d 收到通道关闭信号，退出工作\n", id)
}

/*
==================== 本题并发调试总结 ====================

这道题主要遇到了 3 个并发问题，核心都和：
“goroutine 的生命周期、channel 的阻塞、WaitGroup 的等待顺序”
有关。

--------------------------------------------------
问题1：result 接收 goroutine 启动太晚
--------------------------------------------------

一开始把：

	go func() {
		for res := range result {
			fmt.Println(res)
		}
	}()

写在了 10 个订单全部派发完成之后。但是 result 是无缓冲 channel。

worker 配送完成后会执行：

	result <- res

无缓冲 channel 的发送必须等待接收方。所以最开始的运行过程实际上是：
	3 个 worker 各接到一个订单
		↓
	配送完成
		↓
	result <- res
		↓
	此时 result 还没有接收者
		↓
	3 个 worker 全部阻塞

worker 一旦阻塞在 result <- res，
就无法重新回到：

	for order := range orders

继续接收新订单。

因此：
前 3 个订单被 3 个 worker 接走，
后面的 7 个订单没有 worker 接收，
最后全部触发 time.After 超时。

错误流程：

	main 派订单
	    ↓
	worker 配送
	    ↓
	result <- res
	    ↓
	没有接收者
	    ↓
	worker 阻塞
	    ↓
	无法继续接单

解决方法：

在开始派发订单之前，就启动 result 接收 goroutine。

正确结构：

	main
	 ↓
	orders channel
	 ↓
	3 个 worker
	 ↓
	result channel
	 ↓
	result 接收 goroutine

这样 worker 配送完成后，
result 会被及时接收，
worker 就可以继续接下一单。

--------------------------------------------------
问题2：把 result goroutine 也放进同一个 WaitGroup
--------------------------------------------------

曾经尝试：

	wg.Add(4)

其中：

	3 个 worker
	+
	1 个 result 接收 goroutine

全部使用同一个 wg。

然后 main 中：

	wg.Wait()
	close(result)

result goroutine 是：

	go func() {
		defer wg.Done()

		for res := range result {
			fmt.Println(res)
		}
	}()

这里会形成死锁。

原因：

result goroutine 想结束，
必须先退出：

	for res := range result

但是 for range channel 只有在：

	result 被 close
	并且剩余数据全部读取完

之后才能退出。

所以 result goroutine 的结束条件是：

	close(result)
	    ↓
	for range 退出
	    ↓
	wg.Done()

但是 main 又是：

	wg.Wait()
	    ↓
	等待 result goroutine 的 Done()
	    ↓
	然后才 close(result)

于是形成循环等待：

	main
	 ↓
	wg.Wait()
	 ↓
	等待 result goroutine 结束
	 ↑
	 |
	result goroutine
	 ↓
	等待 close(result)
	 ↑
	 |
	close(result) 又在 wg.Wait() 后面

最终：

	main 等 result goroutine
	result goroutine 等 main close(result)

形成死锁。

--------------------------------------------------
问题3：把 close(result) 放在 wg.Wait() 前面
--------------------------------------------------

后来尝试：

	close(result)
	wg.Wait()

这样也不行。

因为 wg.Wait() 等待的是 3 个 worker。

在 worker 没全部结束之前，
它们仍然可能执行：

	result <- res

如果 main 提前：

	close(result)

那么某个 worker 配送完成后再执行：

	result <- res

就会发生：

	panic: send on closed channel

因此：

不能在发送方还可能继续发送数据时关闭 channel。

本题中 result 的发送方是：

	worker1
	worker2
	worker3

所以必须先确认 3 个 worker 全部结束：

	wg.Wait()

这时才能确定：

	以后再也不会有 worker 执行 result <- res

然后才能安全：

	close(result)

--------------------------------------------------
最终正确的生命周期
--------------------------------------------------

本题使用两个 WaitGroup：

	wg
		负责等待 3 个 worker

	wg2
		负责等待 result 接收 goroutine

运行结束顺序：

	1. main 派发全部订单

	2. close(orders)
	   告诉 worker：
	   不会再有新订单

	3. worker 把已经接到的订单处理完成

	4. worker 的 for range orders 发现 orders 已关闭
	   自动退出循环

	5. 三个 worker 分别执行 wg.Done()

	6. wg.Wait() 等到计数器归零
	   此时可以确定：
	   所有 worker 都已经结束
	   不会再有人往 result 发送数据

	7. close(result)

	8. result goroutine 的：

	   for res := range result

	   读取完剩余数据以后自动退出

	9. result goroutine 执行 wg2.Done()

	10. wg2.Wait() 返回

	11. main 最后输出：

		所有订单派送结束
		成功完成订单数量
		main over

完整依赖关系：

	close(orders)
	     ↓
	等待 workers
	     ↓
	wg.Wait()
	     ↓
	close(result)
	     ↓
	等待 result 接收者
	     ↓
	wg2.Wait()
	     ↓
	main over

--------------------------------------------------
WaitGroup 复习
--------------------------------------------------

WaitGroup 可以理解为一个任务计数器。

	wg.Add(3)

表示：

	还有 3 个任务需要等待

内部计数：

	3

每执行一次：

	wg.Done()

相当于：

	wg.Add(-1)

例如：

	3
	↓ Done()
	2
	↓ Done()
	1
	↓ Done()
	0

wg.Wait() 的作用：

	阻塞调用 Wait() 的 goroutine，
	直到 WaitGroup 的计数器变成 0。

注意：

wg.Wait() 只会阻塞当前调用它的 goroutine，
不会阻塞其他 goroutine。

例如：

	main：
		wg.Wait()  ← main 阻塞

	worker1：继续运行
	worker2：继续运行
	worker3：继续运行

直到三个 worker 都 Done()，
main 才继续往下执行。

--------------------------------------------------
WaitGroup 不会绑定具体 goroutine
--------------------------------------------------

例如：

	wg.Add(4)

并不是：

	Add 的第1个对应 worker1
	Add 的第2个对应 worker2
	Add 的第3个对应 worker3
	Add 的第4个对应 result goroutine

WaitGroup 不知道谁是谁。

它只维护一个数字：

	4

任何一次：

	wg.Done()

都只是：

	4 -> 3 -> 2 -> 1 -> 0

所以问题不在于：

	“某个 Done 没有对应某个 Add”

而在于：

	“这个 goroutine 能不能在正确的时间执行 Done”

--------------------------------------------------
为什么本题用两个 WaitGroup 更清楚
--------------------------------------------------

本题实际上存在两个结束阶段：

第一阶段：

	3 个 worker 全部结束

第二阶段：

	result goroutine 处理完所有结果并结束

而中间必须执行：

	close(result)

因此：

	workerWg.Wait()
	     ↓
	close(result)
	     ↓
	resultWg.Wait()

比强行把所有 goroutine 塞进一个 WaitGroup 更清晰。

--------------------------------------------------
本题最重要的经验
--------------------------------------------------

并发程序出现问题时，
不要只盯着某一行代码。

应该画出：

	谁在等谁？
	谁在发送？
	谁在接收？
	谁负责 close？
	谁必须先结束？
	哪个 goroutine 还可能继续使用 channel？

尤其注意：

1. 无缓冲 channel：
   没有接收者时，发送方会阻塞。

2. close(channel)：
   必须确保以后不会再有发送者发送数据。

3. WaitGroup：
   只负责计数和等待，
   不负责关闭 channel。

4. for range channel：
   channel 被关闭并且数据读完后才会退出。

5. goroutine 的打印顺序：
   不等于真实业务事件发生顺序，
   因为不同 goroutine 的调度顺序是不确定的。

==================================================
*/
