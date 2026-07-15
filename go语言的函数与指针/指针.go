package main

import "fmt"

// 指针是存储另一个变量的内存地址的变量
// 变量是一=一种使用方便的占位符，变量都指向计算机的内存地址，一个指针可以指向任何一个值的内存地址
// go语言的指针不能运算
// 取地址，go中用&来获取变量的地址
func GetAddress() {
	a := 10
	fmt.Printf("变量的地址：%x", &a) //变量的地址：510a4dfd40c0
}

// 声明指针，*T是指针变量的类型，它指向T类型的值
// var 指针变量名 *指针类型 （*号用于指定变量的一个指针）
// var ip *int (指向整型的指针)
// var fp *float32 (指向浮点型的指针)

// 指针示例1
func pointer1() {
	var a int = 120
	var ip *int
	ip = &a // 将a的地址复制给ip
	fmt.Printf("a的类型是%T,值时%v \n", a, a)
	fmt.Printf("&a 的类型是%T,值是%v \n", &a, &a)
	fmt.Printf("ip 的类型是%T,值是%v \n", ip, ip)
	fmt.Printf("*ip 的类型是%T,值是%v \n", *ip, *ip)
	fmt.Printf("*&a 的类型是%T,值是%v \n", *&a, *&a)
	fmt.Println(a, &a, *&a)
	fmt.Println(ip, &ip, *ip, *(&ip), &(*ip))
}

// 指针示例2
func pointer2() {
	type Student struct {
		name    string
		age     int
		married bool
		sex     int8
	}
	var s1 = Student{"steven", 35, true, 1}
	var s2 = Student{"sunny", 20, false, 0}
	var a *Student = &s1
	var b *Student = &s2
	fmt.Printf("s1的类型为%T,值为%v \n", s1, s1)
	fmt.Printf("a的类型为%T,值为%v \n", a, a)
	fmt.Printf("b的类型为%T,值为%v \n", b, b)
	fmt.Printf("*a的类型为%T,值为%v \n", *a, *a)
	fmt.Printf("*b的类型为%T,值为%v \n", *b, *b)
	fmt.Println(s1.name, s1.age, s1.married, s1.sex)
	fmt.Println(a.name, a.age, a.married, a.sex)
	fmt.Println(s2.name, s2.age, s2.married, s2.sex)
	fmt.Println(b.name, b.age, b.married, b.sex)
	fmt.Println((*a).name, (*a).age, (*a).married, (*a).sex)
	fmt.Println((*b).name, (*b).age, (*b).married, (*b).sex)
}

// 空指针（当一个指针没有被分配到任何变量的时候，它的值为nil）
// if(ptr != nil) 	ptr不是空指针
// if（ptr == nil）  ptr是空指针

// 通过使用指针来修改变量的值
func usepointer() {
	b := 3185
	a := &b
	fmt.Printf("b 的地址为：", a)
	fmt.Println("*a 的值为:", *a) //输出 3185
	*a++
	fmt.Println("b 的新值为:", b) // 输出 3186
}

// 使用指针作为函数参数
func parameter() {
	a := 58
	fmt.Println("函数调用之前a的值:", a)
	fmt.Printf("%T \n", a)
	var b *int = &a
	change(b)
	fmt.Println("函数调用之后的a的值为:", a)
}
func change(val *int) {
	*val = 15
}

// 指针数组：就是元素为指针类型的数组
// var ptr [3]*string

// 指针数组
const COUNT int = 4

func pointerNum() {
	a := [COUNT]string{"abc", "ABC", "123", "一二三"}
	i := 0
	var ptr [COUNT]*string
	fmt.Printf("%T,%v \n", ptr, ptr)
	for i = 0; i < COUNT; i++ {
		ptr[i] = &a[i] //将数组中的每个元素的地址赋值给指针数组
	}
	fmt.Println(ptr[0])
	for i = 0; i < COUNT; i++ {
		fmt.Printf("a[%d] = %s \n", i, *ptr[i])
	}
}

// 指针的指针（如果一个指针存放的是另一个指针变量的地址，则这个指针变量称为指向指针的指针变量）
// var ptr **int
func Ppointer() {
	var a int
	var ptr *int
	var pptr **int
	a = 1234
	ptr = &a
	fmt.Println("ptr", ptr)
	pptr = &ptr
	fmt.Println("pptr", pptr)
	fmt.Printf("变量a = &d \n", a)
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **ptr = %d\n", **pptr)
}
