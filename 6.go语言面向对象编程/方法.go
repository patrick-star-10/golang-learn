package main

import (
	"fmt"
	"math"
)

/*
	方法的本质是函数
	1.含义不同：
	函数是一段具有独立功能的代码，可以被反复多次调用，从而实现代码复用，
	而方法(method)是一个类的行为功能，只有该类的对象才能调用

	2.方法有接受者，而函数无接受者：
	go语言的方法(method)是一种作用于特定类型的变量的函数。
	这种特点类型变量叫做接受者(receiver),接受者的概念类似于传统面向对象语言中的this或self关键字。
	go语言的接受者强调了方法具有作用对象，而且=函数没有作用对象，一个方法就是一个包含了接受者的函数
	接受者可以是结构体，也可以是结构体类型外的其他任何类型

	3.函数不可以重名，但是方法可以：
	只要接受者不同，方法就可以相同

	方法的基本语法：
	func(接受者变量 接受者类型)方法名(参数)(返回值){}
*/

// 方法与函数对比
type Employee struct {
	name, currency string
	salary         float64
}

func FuncAndMethod() {
	emp1 := Employee{"Daniel", "$", 2000}
	emp1.printSalary() // method调用
	printSalary(emp1)  // func调用
}

// printSalary()方法
func (e Employee) printSalary() {
	fmt.Printf("员工姓名：%s,薪资：%s%.2f\n", e.name, e.currency, e.salary)
}

// printSalary()函数
func printSalary(e Employee) {
	fmt.Printf("员工姓名：%s,薪资：%s%.2f\n", e.name, e.currency, e.salary)
}

/*
	一段程序可以用函数来写，却还要使用方法，主要有以下两个原因：
	1.go不是一种纯粹面向对象的语言，它不支持类，因此其方法旨在实现类似类的行为
	2.相同名称的方法可以在不同的类型上定义，而且有相同名称的函数是不允许的，假设有一个
	正方形和一个圆形，可以分别在正方形和圆形上定义一个名为Area的求面积的方法
*/
// 相同方法名字案例
type Rectangle struct {
	width, height float64
}
type Circle struct {
	radius float64
}

func sameNameMethod() {
	r1 := Rectangle{10, 4}
	r2 := Rectangle{12, 5}
	c1 := Circle{1}
	c2 := Circle{10}
	fmt.Println("r1面积:", r1.Area())
	fmt.Println("r2面积:", r2.Area())
	fmt.Println("c1面积:", c1.Area())
	fmt.Println("c2面积:", c2.Area())
}

// 定义Rectangle的方法
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// 定义Cricle的方法
func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

// 若方法的接受者不是指针，实际知识获取了一个拷贝，而不能真正改变接受者原来的数据
// 当指针作为接受者时，情况如例：
func PointerReceiver() {
	r1 := Rectangle{5, 8}
	r2 := r1
	// 打印对象的内存地址
	fmt.Printf("r1的地址:%p\n ", &r1)
	fmt.Printf("r2的地址:%p\n ", &r2)
	r1.setValue()
	fmt.Println("r1.height=", r1.height)
	fmt.Println("r2.height=", r2.height)
	fmt.Println("-------------")
	r1.setValue2()
	fmt.Println("r1.height=", r1.height)
	fmt.Println("r2.height=", r2.height)
	fmt.Println("-------------")

}
func (r Rectangle) setValue() {
	fmt.Printf("setValue方法中r的地址:%p\n", &r)
	r.height = 10
}
func (r *Rectangle) setValue2() {
	fmt.Printf("setValue2中r的地址:%p \n", &r)
	r.height = 20
	fmt.Printf("r指向的地址: %p\n", r)
}

/*
	方法继承
	方法是可以继承的，如果匿名字段实现了一个方法，那么包含这个匿名字段的struct也能调用匿名字段中的方法
*/
// 方法继承案例
type Human struct {
	name, phone string
	age         int
}
type Student2 struct {
	Human  //匿名字段
	school string
}
type Employee2 struct {
	Human   //匿名字段
	company string
}

func inheritMethod() {
	s1 := Student2{Human{"Daniel", "123456789", 13}, "野鸡小学"}
	e1 := Employee2{Human{"steven", "987654321", 34}, "1000phone"}
	s1.SayHi()
	e1.SayHi()
}
func (h *Human) SayHi() {
	fmt.Printf("大家好!我是%s,我%d岁,我的联系方式是:%s\n", h.name, h.age, h.phone)
}

/*
	方法重写：
	在go语言中，方法重写是指一个包含了匿名字段的struct也实现了该匿名字段实现的方法
*/
//方法重写案例
func MethodRewrite() {
	s1 := Student2{Human{"Daniel", "123456789", 13}, "野鸡小学"}
	e1 := Employee2{Human{"steven", "987654321", 34}, "1000phone"}
	s1.SayHi()
	e1.SayHi()
}
func (s *Student2) SayHi() {
	fmt.Printf("大家好!我是%s,我%d岁,我在%s上学,联系方式:%s\n", s.name, s.age, s.school, s.phone)
}
func (e *Employee2) SayHi() {
	fmt.Printf("大家好!我是%s,我%d岁,我在%s工作,联系方式:%s\n", e.name, e.age, e.company, e.phone)
}

// 结构体会优先调用自己实现的同名方法；
// 如果没有实现，则会调用匿名嵌入（模拟继承）结构体的方法。
