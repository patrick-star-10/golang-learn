package main

import "fmt"

// 声明方式： const 变量名 变量类型 = 表达式

// 枚举
func sex() {
	const (
		Unkown = 0
		Female = 1
		Male   = 2
	)

	fmt.Println(Unkown, Female, Male) // 输出: 0 1 2
}

// 常量数组中如果不指定类型和初始值，则与上一行非空常量的值相等
func constarray() {
	const (
		a = 1
		b
		c
	)
	fmt.Println(a, b, c) // 输出: 1 1 1
}

// iota 是一个常量计数器，只能在常量的表达式中使用。iota 在 const 关键字出现时将被重置为 0
// const 中每新增一行常量声明将使 iota 计数一次（即 iota 的值在 const 语句中自动加 1）
func iotaarray() {
	const (
		a = iota // 0
		b = iota // 1
		c = iota // 2
	)
	fmt.Println(a, b, c) // 输出: 0 1 2
}

// 因为常量数组中如果不指定类型和初始值，则与上一行非空常量的值相等，所以可以简写为：
func iotaarray2() {
	const (
		a = iota // 0
		b        // 1
		c        // 2
	)
	fmt.Println(a, b, c) // 输出: 0 1 2
}
