package main

import "fmt"

// 使用闭包实现计数器
func counter() {
	pos := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("i=%d \t", i)
		fmt.Println(pos(i))
	}
	fmt.Println("---------------------------------------")
	for i := 0; i < 10; i++ {
		fmt.Printf("i=%d \t", i)
		fmt.Println(pos(i))
	}
}
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		fmt.Printf("sum1=%d \t", sum)
		sum += x
		fmt.Printf("sum2=%d \t", sum)
		return sum
	}
}

// 案例二
func bibao() {
	myfunc := Counter()
	fmt.Println("myfunc", myfunc)
	// 调用myfunc函数，i变量自增1并返回
	fmt.Println(myfunc())
	fmt.Println(myfunc())
	fmt.Println(myfunc())
	// 创建新的函数nextNumber1,并查看结果
	myfunc1 := Counter()
	fmt.Println("myfunc1", myfunc1)
	fmt.Println(myfunc1())
	fmt.Println(myfunc1())
}

// 计数器，闭包函数
func Counter() func() int {
	i := 0
	res := func() int {
		i += 1
		return i
	}
	fmt.Println("Counter中的内部函数:", res)
	return res
}
