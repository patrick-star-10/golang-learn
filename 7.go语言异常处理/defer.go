package main

import (
	"fmt"
)

/*
	defer
	关键字defer用来延迟一个函数或者方法(或者当前所创建的匿名函数)的执行
	defer只能出现在函数或者方法的内部
	函数中使用defer
	在函数中可以提添加多个defer语句，若果有很多调用defer,当函数执行到最后时，这些defer语句会按照逆序执行
	报错的时候也会执行，最后该函数返回
*/
// defer执行顺序案例
func deferTurn() {
	defer funcA()               //最后执行
	funcB()                     //第一个执行
	defer funcC()               //第三个执行
	fmt.Println("main over...") //第二个执行
}
func funcA() {
	fmt.Println("这是funcA")
}
func funcB() {
	fmt.Println("这是funcB")
}
func funcC() {
	fmt.Println("这是funcC")
}

/*
	defer语句经常被用来处理成对的操作，如打开-关闭，连接-断开连接，加锁-释放锁
	特别是在执行打开资源的操作时，遇到错误需要提前返回，在返回前需要关闭相应的资源，不然很容易造成资源泄漏等问题
*/
//defer在函数中的使用案例
func deferTurn2() {
	s1 := []int{78, 100, 2, 400, 324}
	getLargest(s1)
}
func finished() {
	fmt.Println("结束")
}
func getLargest(s []int) {
	defer finished()
	fmt.Println("开始寻找最大数值")
	max := s[0]
	for _, v := range s {
		if v > max {
			max = v
		}
	}
	fmt.Printf("%v中最大数为:%v\n", s, max)
}

// 方法中使用defer
type person struct {
	firstName, lastName string
}

func (p person) fullName() {
	fmt.Printf("%s %s\n", p.firstName, p.lastName)
}
func deferInmethod() {
	p := person{"steven", "Wang"}
	defer p.fullName()
	fmt.Print("Wellcome ")
}

/*
defer参数
延迟函数的参数在执行延迟语句时被执行，而不是在执行实际的函数调用时执行
*/
func DeferParameter() {
	a := 5
	b := 6
	defer printAdd(a, b, true)
	a = 10
	b = 7
	printAdd(a, b, false)
}
func printAdd(a, b int, flag bool) {
	if flag {
		fmt.Printf("延迟执行函数printAdd(),参数a,b分别为%d,%d,两数之和为:%d\n", a, b, a+b)
	} else {
		fmt.Printf("未延迟执行函数printAdd(),参数a,b分别为%d,%d,两数之和为:%d\n", a, b, a+b)
	}
}
