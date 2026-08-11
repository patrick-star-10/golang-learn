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
