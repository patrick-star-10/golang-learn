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

/*
	go语言为开发者提供了专用于拦截运行时panic的内建函数recover()
	recover()可以让进入恐慌流程的Goroutine恢复过来并重新获得流程控制权
	需要注意的是，recover()让程序恢复，必须在延迟函数中执行。换而言之，revover仅在延迟函数中有效
	在正常的程序运行过程中，调用recover()会返回nil,并且没有其他任何效果。
	如果当前的Goroutine陷入恐慌，调用recover()可以捕获panic()的输入值，使程序恢复正常运行
*/
// recover示例
func TEST() {
	funca()
	funcb()
	funcc()
	fmt.Println("main over")
}
func funca() {
	fmt.Println("这是funca")
}
func funcb() {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println("恢复啦,获取recover的返回值", msg)
		}
	}()
	fmt.Println("这是funcb")
	for i := 0; i < 10; i++ {
		fmt.Println("i:", i)
		if i == 5 {
			//panic("funcb恐慌了")
		}
	}
}
func funcc() {
	defer func() {
		fmt.Println("执行延迟函数")
		msg := recover()
		fmt.Println("获取recover的返回值: ", msg)
	}()
	fmt.Println("这是funcc")
	panic("func恐慌了")
}
