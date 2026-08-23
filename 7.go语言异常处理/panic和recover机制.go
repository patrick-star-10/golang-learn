package main

import (
	"fmt"
)

/*
	panic
	让当前的程序进入恐慌，中断程序的执行。panic()是一个内建函数，可以中断原有的控制流程
*/
//panic 示例1
func TestA() {
	fmt.Println("func TestA()")
}
func TestB() {
	panic("func TestB():panic")
}
func TestC() {
	fmt.Println("func TestC()")
}
func Test() {
	TestA()
	TestB() //TestB()发生异常，中断程序
	TestC()
}

/*
	通常情况下，向程序使用报告错误状态的方式可以是返回一个额外的error类型值
	但是，当遇到不可恢复的错误状态时，如数组访问越界，空指针引用等，这些运行时错误会引起panic异常
	这时候，上述错误的处理方式显示就不合适了
	需要注意的是，不应该通过调用panic()函数来报告普通的错误，而应该只把它作为报告致命错误的一种方式
	当某些不应该发生的场景发生时调用panic()
*/
//内置的panic()函数引发的panic异常示例
func Test1() {
	fmt.Println("func Test1()")
}
func Test2(x int) {
	var a [100]int
	a[x] = 1000 //x值为101时，数组越界
}
func Test3() {
	fmt.Println("func Test3()")
}
func MAIN() {
	Test1()
	Test2(101) //Test2()发生异常，中断程序
	Test3()
}
