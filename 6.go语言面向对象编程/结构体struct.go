package main

import (
	"fmt"
	"math"
)

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
	fmt.Println("d5修改后: ", d5) //d5修改后:  &{多多 金色 1 金毛}
	fmt.Println("d4: ", d4)    //&{多多 金色 1 金毛}
}

// 结构体作为函数的参数以及返回值：值传递和引用传递
// 结构体对象与指针在函数中的传递
type Flower struct {
	name, color string
}

func Rose() {
	//1.结构体作为参数的用法
	f1 := Flower{"玫瑰", "红"}
	fmt.Printf("f1:%T,%v,%p \n", f1, f1, &f1)
	fmt.Println("-------------------")
	//将结构体对象作为参数
	changeInfo1(f1)
	fmt.Printf("f1:%T,%v,%p \n", f1, f1, &f1)
	fmt.Println("-------------------")
	// 将结构体指针作为参数
	changeInfo2(&f1)
	fmt.Printf("f1:%T,%v,%p \n", f1, f1, &f1)
	fmt.Println("-------------------")
	//2.结构体作为返回值的用法
	f2 := getFlower1()
	f3 := getFlower1()
	fmt.Println("更改前", f2, f3)
	fmt.Printf("f2地址为%p,f3地址为%p\n", &f2, &f3) //地址发生改变，对象发生了复制
	f2.name = "杏花"
	fmt.Println("更改后:", f2, f3)
	//结构体指针作为返回值
	f4 := getFlower2()
	f5 := getFlower2()
	fmt.Println("更改前:", f4, f5)
	f4.name = "桃花"
	fmt.Println("更改后:", f4, f5)
}

// 返回结构体对象
func getFlower1() (f Flower) {
	f = Flower{"牡丹", "白"}
	fmt.Printf("函数getFlower1内f:%T,%v,%p \n", f, f, &f)
	return
}

// 返回结构体指针
func getFlower2() (f *Flower) {
	temp := Flower{"芙蓉", "红"}
	fmt.Printf("函数getFlower2内temp:%T,%v,%p \n", temp, temp, &temp)
	f = &temp
	fmt.Printf("函数getFlower2内f:%T,%v,%p,%p \n", f, f, f, &f)
	return
}

// 传结构体对象
func changeInfo1(f Flower) {
	f.name = "月季"
	f.color = "粉"
	fmt.Printf("changeInfo1内f:%T,%v,%p \n", f, f, &f)
}

// 传结构体指针
func changeInfo2(f *Flower) {
	f.name = "蔷薇"
	f.color = "紫"
	fmt.Printf("changeInfo1内f:%T,%v,%p ,%p\n", f, f, f, &f)
}

/*
	匿名结构体和匿名字段
	匿名结构体就是没有名字的结构体，无需通过type关键字来定义就可以直接使用。
	创建匿名结构体时，同时要创建对象，格式如下：
	变量名 ：= struct{
		定义成员属性
	}{初始化成员属性}
*/
//匿名结构体
func UnameStruct() {
	//匿名函数
	res := func(a, b float64) float64 {
		return math.Pow(a, b)
	}(2, 3)
	fmt.Println(res)
	// 匿名结构体
	addr := struct {
		province, city string
	}{"陕西省", "西安市"}
	fmt.Println(addr)
	cat := struct {
		name, color string
		age         int8
	}{
		name:  "绒毛",
		color: "黑白",
		age:   1,
	}
	fmt.Println(cat)
}

/*
	结构体的匿名字段
	匿名字段就是在结构体中的字段没有名字，只包含一个没有字段名的类型。
	如果字段没有名字，那么默认使用类型作为字段名，同一个类型只能有一个匿名字段。
	结构体嵌套中采用匿名结构体字段可以模拟继承关系
*/
//匿名字段
type User struct {
	string
	byte
	int8
	float64
}

func user() {
	// 实例化结构体
	user := User{"Steven", 'm', 35, 177.5}
	fmt.Println(user)
	//如果想依次输出姓名，年龄，性别，身高
	fmt.Printf("姓名:%s\n", user.string)
	fmt.Printf("身高:%2f \n", user.float64)
	fmt.Printf("性别：%c \n", user.byte)
	fmt.Printf("年龄:%d\n", user.int8)
}

/*
	结构体嵌套
	将一个结构体作为另一个结构体的属性(字段)，这种结构就是结构体嵌套
	结构体嵌套可以模拟面向对象编程中的以下两种关系：
	1.聚合关系：一个类作为另一个类的属性
	2.继承关系：一个类作为另一个类的子类，子类和父类的关系
*/
//聚合关系
type Address struct {
	province, city string
}
type Person struct {
	name    string
	age     int
	address *Address
}

func Aggregate() {
	//模拟结构体对象之间的聚合关系
	p := Person{}
	p.name = "steven"
	p.age = 35
	//赋值方式1
	addr := Address{}
	addr.province = "北京市"
	addr.city = "海淀区"
	p.address = &addr
	fmt.Println(p)
	fmt.Println("姓名: ", p.name, "年龄: ", p.age, "省: ", p.address.province, "市区:", p.address.city)
	fmt.Println("--------------------")
	//修改Person对象的数据，是否会影响Address对象的数据？
	p.address.city = "昌平区"
	fmt.Println("姓名: ", p.name, "年龄: ", p.age, "省: ", p.address.province, "市区:", p.address.city)
	fmt.Println("---------------------")
	// 修改Address对象的数据，是否会影响Person对象的数据？
	addr.city = "大兴区"
	fmt.Println("姓名: ", p.name, "年龄: ", p.age, "省: ", p.address.province, "市区:", p.address.city)
	fmt.Println("---------------------")
	//赋值方式2
	p.address = &Address{ // 创建了一个新的address对象
		province: "陕西省",
		city:     "西安市",
	}
	fmt.Println(p)
	fmt.Println("姓名: ", p.name, "年龄: ", p.age, "省: ", p.address.province, "市区:", p.address.city)
	fmt.Println("---------------------")
}

/*
	继承
	子类可以有自己的属性和方法，也可以重写父类已有的方法，子类可以直接访问父类的所有属性和方法
	在结构体中，属于匿名结构体的字段称为提升字段，他们可以被访问，匿名结构体就像是该结构体的父类
	采用匿名字段的形式就是模拟继承关系
*/
//结构体嵌套模拟继承关系
type Person1 struct {
	name string
	age  int
	sex  string
}
type Student struct {
	Person1
	schoolName string
}

func inherit() {
	//1.实例化并且初始化Person2
	p1 := Person1{"steven", 18, "男"}
	fmt.Println(p1)
	fmt.Println("--------------")
	//2.实例化并且初始化Student
	// 写法1：
	s1 := Student{p1, "北航软件学院"}
	printInfo(s1)
	//写法2:
	s2 := Student{Person1{"john", 30, "男"}, "北京大学"}
	printInfo(s2)
	//写法3；
	s3 := Student{Person1: Person1{
		name: "penn",
		age:  19,
		sex:  "男",
	},
		schoolName: "清华",
	}
	printInfo(s3)
	//写法4:体现继承关系
	s4 := Student{}
	s4.name = "Daniel"
	s4.sex = "男"
	s4.age = 12
	s4.schoolName = "野鸡大学"
	printInfo(s4)
}
func printInfo(s1 Student) {
	fmt.Println(s1)
	fmt.Printf("%+v\n", s1)
	fmt.Printf("姓名: %s, 年龄: %d, 性别:%s, 学校: %s\n", s1.name, s1.age, s1.sex, s1.schoolName)
	fmt.Println("---------------------")
}

// 结构体嵌套时，可能存在相同的成员名，成员重名会导致成员名字冲突
type A struct {
	a, b int
}
type B struct {
	a, d int
}
type C struct {
	A
	B
}

func conflict() {
	c := C{}
	c.A.a = 1
	c.B.a = 2 //如果调用c.a = 2则会提示“引起歧义的参数”
	c.b = 3
	c.d = 4
	fmt.Println(c)
}
