package main

import (
	"fmt"
	"math"
)

/*
func(参数列表)（返回参数列表）{
// 函数体
}
*/
// 在定义时调用匿名函数
func anonymous() {
	func(data int) {
		fmt.Println("hello", data)
	}(100) //前面的函数体可以看成func(data) 100是在传参=> func(100)
}

// 将匿名函数赋值给变量
func anonymous1() {
	f := func(data string) {
		fmt.Println(data)
	}
	f("欢迎学习区块链")
}

// 匿名函数作为回调函数
func anonymous2() {
	// 调用回调函数对每个元素进行求平方根
	arr := []float64{1, 9, 16, 25, 30}
	visit(arr, func(v float64) {
		v = math.Sqrt(v)
		fmt.Printf("%.2f \n", v)
	})
	// 调用函数，对每个元素进行求平方和
	visit(arr, func(v float64) {
		v = math.Pow(v, 2)
		fmt.Printf("%.2f \n", v)
	})
}

// 定义visit，遍历切片元素，对每个元素进行处理
func visit(list []float64, f func(float64)) {
	for _, value := range list {
		f(value)// 把参数传给f函数
	}
}

