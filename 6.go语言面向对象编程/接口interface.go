package main

import (
	"fmt"
)

/*
	接口概念：
	面向对象语言中，接口用于定义对象的行为，接口只指定对象应该做什么，实现这种行为的方式由对象决定
	在go语言中，接口是一组方法签名，接口指定了类型应该具有哪些方法，类型决定如何实现这些方法，当某个类型为接口
	类型中的所有方法提供了具体的实现细节的时，这个类型就被称为实现了该接口
	接口定义了一组方法，如果某个对象实现了该接口的所有方法，此对象就实现了该接口
	go语言的类型都是隐式实现接口的。任何定义了接口中的所用方法的类型都被称为隐式地实现了该接口

	定义接口的语法格式如下：
	type 接口名字 interface {
		方法1([参数列表])[返回值]
		方法2([参数列表])[返回值]
		.......
		方法n([参数列表])[返回值]
	}
	实现接口的的语法格式如下：
	func (变量名 结构体类型)方法1([参数列表])[返回值]{
		方法体
	}
*/

// 接口具体使用案例
type Phone interface {
	call()
}
type AndroidPhione struct {
}
type IPhone struct {
}

func (a AndroidPhione) call() {
	fmt.Println("我是安卓手机,可以打电话")
}
func (i IPhone) call() {
	fmt.Println("我是苹果手机,可以打电话")
}
func InterfaceUse() {
	//定义接口类型的变量
	var phone Phone
	phone = new(AndroidPhione)
	fmt.Printf("%T,%v,%p\n", phone, phone, &phone)
	phone = AndroidPhione{}
	fmt.Printf("%T,%v,%p\n", phone, phone, &phone)
	phone.call()
	phone = new(IPhone)
	fmt.Printf("%T,%v,%p\n", phone, phone, &phone)
	phone.call()
	phone = IPhone{}
	fmt.Printf("%T,%v,%p\n", phone, phone, &phone)
	phone.call()
}
