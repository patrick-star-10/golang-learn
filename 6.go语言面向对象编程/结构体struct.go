package main

import "fmt"

/*
结构体的定义格式：
type 类型名 struct{
	成员属性1，类型1
	成员属性2，类型2
	成员属性，类型3
	....
}
使用结构体的过程中注意以下3点：
1.类型名称是标识结构体的名称，在同一个包内不能重复
2.结构体的属性，也叫字段，必须唯一
3.同类型的成员属性可以写在一行
结构体的定义只是一种内存布局的描述，只有当结构体示例化，才会真正分配内存。因为定义后要实例化后才能使用
实例化就是根据结构体的定义格式创建一份与格式一致的内存区域。结构体每个实例的内存是完全独立的
*/

// 结构体实例化
type Teacher struct {
	name string
	age  int8
	sex  byte
}

func teacher() {
	//1.var声明方式实例化结构体，初始化方式为；对象.属性
	var t1 Teacher
	fmt.Println(t1)                         // {0,0}
	fmt.Printf("t1:%T,%v,%q\n", t1, t1, t1) //t1:main.Teacher,{ 0 0},{"" '\x00' '\x00'}
	t1.name = "Steven"
	t1.age = 15
	t1.sex = 1
	fmt.Println(t1) // {Steven 15 1}
	fmt.Println("--------------")
	//2.变量间断声明格式实例化结构体
	t2 := Teacher{}
	t2.name = "David"
	t2.age = 30
	t2.sex = 1
	fmt.Println(t2) // {David 30 1}
	fmt.Println("-----------")
	//3.变量简短声明格式实例化结构体，类似map的用法
	t3 := Teacher{
		name: "josh",
		age:  28,
		sex:  1,
	}
	fmt.Println(t3) //{josh 28 1}
	t3 = Teacher{name: "josh2", age: 27, sex: 1}
	fmt.Println(t3) // {josh2 27 1}
	fmt.Println("--------------")
	//4.变量简短声明格式实例化结构体，不写属性名，俺属性顺序只写属性值
	t4 := Teacher{"Rudy", 30, 0}
	fmt.Println(t4) //{Rudy 30 0}
	fmt.Println("-------------")
}

/*
结构体语法糖
通常来说语法糖能够提升程序的可读性，从而减少程序代码出错的机会，结构体和数据结构都有语法糖
使用内置函数new()对结构体进行实例化，结构体实例化后形成指针类型的结构体，new()内置函数会分配内存。
第一个参数是类型，而不是值，返回的值是指向该类型新分配的零值的指针。该数用来创建某个类型的指针
*/
// 使用方式：语法糖
type Emp struct {
	name string
	age  int8
	sex  byte
}

func EmpStruct() {
	// 使用new()内置函数实例化struct
	emp1 := new(Emp)
	fmt.Println(emp1)                               // &{0,0}打印的是指针类型
	fmt.Printf("emp1:%T,%v,%p\n", emp1, emp1, emp1) //emp1:*main.Emp,&{ 0 0},0x595524c2c000
	(*emp1).name = "Davied"
	(*emp1).age = 30
	(*emp1).sex = 1
	fmt.Println(*emp1) //{Davied 30 1}
	// 语法糖写法：emp1.name 等价于 (*emp1).name
	emp1.name = "David2"
	emp1.age = 31
	emp1.sex = 1
	fmt.Println(emp1) // &{David2 31 1}
	fmt.Println("-------------")
	SyntacticSugar()
}
func SyntacticSugar() {
	// 数组中的语法糖
	arr := [4]int{10, 20, 30, 40}
	arr2 := &arr
	fmt.Println((*arr2)[len(arr)-1]) //40
	fmt.Println(arr2[0])             //10（语法糖会自动解引用）

	// 切片中的语法糖
	arr3 := []int{100, 200, 300, 400}
	arr4 := &arr3
	fmt.Println((*arr4)[len(arr)-1]) //400
	//fmt.Println(arr4[0]) 会报错，不存在语法糖
}

// 结构体是值类型，在函数中对参数进行修改，不会影响到实际参数
// 证明结构体是值类型
type human struct {
	name string
	age  int8
	sex  byte
}

func StructType() {
	//1.初始化Human
	h1 := human{"Steven", 35, 1}
	fmt.Printf("h1:%T,%v,%p \n", h1, h1, &h1)
	fmt.Println("----------------")
	//2.复制结构体对象
	h2 := h1
	h2.name = "David"
	h2.age = 30
	fmt.Printf("h2修改后:%T,%v,%p \n", h2, h2, &h2)
	fmt.Printf("h1:%T,%v,%p \n", h1, h1, &h1)
	fmt.Println("-------------------")
	//3.结构体对象作为参数传递
	changeName(h1)
	fmt.Printf("h1:%T,%v,%p \n", h1, h1, &h1)
}
func changeName(h human) {
	h.name = "Daniel"
	h.age = 13
	fmt.Printf("函数体内修改后:%T,%v,%p \n", h, h, &h) // h是新的结构体，不影响h1
}

// 结构体的深拷贝和浅拷贝
// 值类型是深拷贝，深拷贝就是为新的对象非配了内存。引用类型就是浅拷贝，浅拷贝只是复制了对象的指针。
type Dog struct {
	name  string
	color string
	age   int8
	kind  string
}

func DeepAndShallowCopy() {
	//1.实现结构体深拷贝
	// struct是值类型，默认的复制就是身拷贝
	d1 := Dog{"一二", "白色", 2, "熊猫"}
	fmt.Printf("d1:%T,%v,%p \n", d1, d1, &d1)
	d2 := d1 //深拷贝
	fmt.Printf("d2:%T,%v,%p \n", d2, d2, &d2)
	d2.name = "布布"
	fmt.Println("d2修改后: ", d2)
	fmt.Println("d1: ", d1)
	fmt.Println("------------")
	//2.实现结构体浅拷贝：直接赋值指针地址
	d3 := &d1
	fmt.Printf("d3:%T,%v,%p \n", d3, d3, &d3)
	d3.name = "球球"
	d3.color = "白色"
	d3.kind = "萨摩耶"
	fmt.Println("d3修改后:", d3)
	fmt.Println("d1: ", d1)
	fmt.Println("-------------")
	//3.实现结构体浅拷贝：通过new()函数来实例化对象
	d4 := new(Dog)
	d4.name = "多多"
	d4.color = "棕色"
	d4.age = 1
	d4.kind = "巴哥犬"
	d5 := d4 // new() 返回的是 *Dog，因此这里复制的是指针（地址），d4 和 d5 指向同一个 Dog 对象
	fmt.Printf("d4:%T,%v,%p \n", d4, d4, d4)
	fmt.Printf("d5:%T,%v,%p \n", d5, d5, d5)
	fmt.Println("--------------")
	d5.color = "金色"
	d5.kind = "金毛"
	fmt.Println("d5修改后: ", d5)
	fmt.Println("d4: ", d4)
}
