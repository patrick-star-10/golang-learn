package main

import (
	"fmt"
	"math"
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

// duck typing
type ISayHello interface {
	SayHello() string
}
type Duck struct {
	name string
}
type Person2 struct {
	name string
}

func (d Duck) SayHello() string {
	return d.name + "叫:ga ga ga"
}
func (p Person2) SayHello() string {
	return p.name + "说:你好"
}
func DuckTyping() {
	//定义实现接口的对象
	duck := Duck{"Yaya"}
	person := Person2{"steven"}
	fmt.Println(duck.SayHello())
	fmt.Println(person.SayHello())
	fmt.Println("----------------")
	//定义接口类型的变量
	var i ISayHello
	i = duck
	fmt.Printf("%T,%v,%p", i, i, &i)
	fmt.Println(i.SayHello())
	i = person
	fmt.Printf("%T,%v,%p", i, i, &i)
	fmt.Println(i.SayHello())
}

/*
	多态：
	如果有几个相似而完全不同的对象，有时人们要求在向他们发出同一个消息时，他们的反应各有不同，
	分别执行不同的操作，这种情况就是多态
	多态就是事物的多种形态，go语言中的多态性时在接口的帮助下实现的--定义接口，创建实现接口的结构体对象
	定义接口类型的对象，可以保存实现该接口的任何类型的值。go语言接口变量的这个特性实现了go语言中的多态性
	接口类型对象，不能访问其实现类中的属性字段
*/

// 多态
type Income interface {
	calculate() float64 //计算收入总额
	source() string     // 说明收入来源
}

// 固定账单项目
type FixedBilling struct {
	projectName  string  // 	工程项目
	biddedAmount float64 // 项目招标总额
}

// 定时生成项(定时和材料项目)
type TimeAndMaterial struct {
	projectName string
	workHours   float64 // 工作时长
	hourlyRate  float64 //每小时工资率
}

// 固定收入x项目
func (f FixedBilling) calculate() float64 {
	return f.biddedAmount
}
func (f FixedBilling) source() string {
	return f.projectName
}

// 定时收入项目
func (t TimeAndMaterial) calculate() float64 {
	return t.workHours * t.hourlyRate
}
func (t TimeAndMaterial) source() string {
	return t.projectName
}

// 通过广告点击收入
type Advertisement struct {
	adNane         string
	clickCount     int
	incomePerclick float64
}

func (a Advertisement) calculate() float64 {
	return float64(a.clickCount) * a.incomePerclick
}
func (a Advertisement) source() string {
	return a.adNane
}
func SUM() {
	p1 := FixedBilling{"项目1", 5000}
	p2 := FixedBilling{"项目2", 10000}
	p3 := TimeAndMaterial{"项目3", 100, 40}
	p4 := TimeAndMaterial{"项目4", 250, 20}
	p5 := Advertisement{"广告1", 10000, 0.1}
	p6 := Advertisement{"广告2", 20000, 0.05}
	ic := []Income{p1, p2, p3, p4, p5, p6}
	fmt.Println(calculateNetIncome(ic))
}

// 计算净收入
func calculateNetIncome(ic []Income) float64 {
	netincome := 0.0
	for _, income := range ic {
		fmt.Printf("收入来源:%s,收入金额:%.2f\n", income.source(), income.calculate())
		netincome += income.calculate()
	}
	return netincome
}

/*
	空接口：
	空接口中没有任何方法，任意类型都可以实现该接口，空接口这样定义：interface{},
	也就是包含0个方法的接口
	空接口常用于以下情形：
	1.println参数就是空接口
	2.定义一个map:key是string,value时任意数据类型
	3.定义一个切片，其中存储任意类型的数据
*/
//空接口案例
type A1 interface {
}
type Cat struct {
	name string
	age  int
}
type Person3 struct {
	name string
	sex  string
}

func spareInterface() {
	var a1 A1 = Cat{"Mimi", 1}
	var a2 A1 = Person3{"steven", "男"}
	var a3 A1 = "Learn golang"
	var a4 A1 = 100
	var a5 A1 = 3.14
	showInfo(a1)
	showInfo(a2)
	showInfo(a3)
	showInfo(a4)
	showInfo(a5)
	fmt.Println("-------------")
	//1.fmt.Println参数就是空接口
	fmt.Println("println的参数就是空接口,可以是任何数据类型", 100, 3.14, Cat{"万才", 2})
	//2.定义map,value时任何数据类型
	map1 := make(map[string]interface{})
	map1["name"] = "Daniel"
	map1["age"] = 13
	map1["height"] = 1.71
	fmt.Println(map1)
	fmt.Println("-------------")
	// 3.定义一个切片，其中储存任意数据类型
	slice1 := make([]interface{}, 0, 10)
	slice1 = append(slice1, a1, a2, a3, a4, a5)
	fmt.Println(slice1)
}
func showInfo(a A1) {
	fmt.Printf("%T,%v\n", a, a)
}

/*
	接口对象转型
	接口对象转型第一种方式实例如下：
	instance,ok := 接口对象.(实际类型)
	如果该接口对象时对应的实际类型，那么instance就是转型之后的对象，ok的值为true,配合if...else if..语句使用
	接口对象转型第二章方式示例如下：
	接口对象.(type)，此方式配合switch...case语句使用
*/
//接口对象转型案例
//1.定义接口
type Shape interface {
	perimeter() float64
	area() float64
}

// 2.矩形
type Rectangle2 struct {
	a, b float64
}

// 3.三角形
type Triangle struct {
	a, b, c float64
}

// 4.原型
type Circle2 struct {
	radius float64
}

// 定义实现接口的方法
func (r Rectangle2) perimeter() float64 {
	return (r.a + r.b) * 2
}
func (r Rectangle2) area() float64 {
	return r.a * r.b
}
func (t Triangle) perimeter() float64 {
	return t.a + t.b + t.c
}
func (t Triangle) area() float64 {
	//海伦公式
	p := t.perimeter()
	return math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
}
func (c Circle2) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func (c Circle2) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

// 接口对象转型方式1(判断接口里的数据类型)
// instance,ok := 接口对象.(实际类型)		(断言语法)
func getType(s Shape) {
	if instance, ok := s.(Rectangle2); ok {
		fmt.Printf("矩形:长度%.2f,宽度%.2f, ", instance.a, instance.b)
	} else if instance, ok := s.(Triangle); ok {
		fmt.Printf("三角形:三边分别:%.2f,%.2f,%.2f,", instance.a, instance.b, instance.c)
	} else if instance, ok := s.(Circle2); ok {
		fmt.Printf("圆形:半径%.2f, ", instance.radius)
	}
}

// 接口对象转型方式2
// 接口对象.(type)，配合switch和case语句来断言
func getType2(s Shape) {
	switch instance := s.(type) {
	case Rectangle2:
		fmt.Printf("矩形:长度%.2f,宽度%.2f, ", instance.a, instance.b)
	case Triangle:
		fmt.Printf("三角形:三边分别:%.2f,%.2f,%.2f,", instance.a, instance.b, instance.c)
	case Circle2:
		fmt.Printf("圆形:半径%.2f, ", instance.radius)
	}
}
func getResult(s Shape) {
	getType2(s)
	fmt.Printf("周长:%.2f,面积:%.2f\n", s.perimeter(), s.area())
}
func Assert() {
	var s Shape
	s = Rectangle2{3, 4}
	getResult(s)
	showInfo2(s)
	s = Triangle{3, 4, 5}
	getResult(s)
	showInfo2(s)
	s = Circle2{1}
	getResult(s)
	showInfo2(s)
	x := Triangle{3, 4, 5}
	fmt.Println(x)
}
func (t Triangle) String() string { //实现了系统接口，最后打印部分会改变
	return fmt.Sprintf("Triangle对象,属性分别为:%.2f,%.2f,%.2f", t.a, t.b, t.c)
}
func showInfo2(s Shape) {
	fmt.Printf("%T,%v\n", s, s)
	fmt.Println("---------------")
}
