package main

import "fmt"

// var 变量名 变量类型 = 表达式
var a1 int = 1

// 变量声明时可以省略类型，编译器会根据表达式自动推导出变量类型
var b1 = "hello"

// 批量声明未初始化变量
var (
	a2 int
	b  string
	c  []float32
	d  func() bool
	e  struct {
		x int
		y string
	}
)

// 初始化变量的简短声明格式 (只能在函数内部使用)
// 变量名 := 表达式

func main1() {
	a3 := 1
	fmt.Println(a3)
}

// 匿名变量
func GetData() (int, int) {
	return 10, 20
}
func print() {
	a4, _ := GetData() // 忽略第二个返回值
	_, b3 := GetData() // 忽略第一个返回值
	fmt.Println(a4, b3)
}
