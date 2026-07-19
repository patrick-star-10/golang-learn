package main

import (
	"fmt"
	"strings"
)

// 变量的作用域
// 声明全局变量
var a1 int = 7
var b1 int = 9

func scope() {
	a1, b1, c1 := 10, 20, 0
	fmt.Printf("scope函数中 a1 = %d\n", a1) //输出10
	fmt.Printf("scope函数中 b1 = %d\n", b1) //输出20
	fmt.Printf("scope函数中 c1 = %d\n", c1) //输出0
	c1 = sum(a1, b1)
	fmt.Printf("scope函数中 c1 = %d\n", c1) // 输出33

}

// 定义两数字相加
func sum(a1, b1 int) (c1 int) {
	a1++
	b1 += 2
	c1 = a1 + b1
	fmt.Printf("sum函数中 a1 = %d\n", a1) //输出11
	fmt.Printf("sum函数中 b1 = %d\n", b1) //输出22
	fmt.Printf("sum函数中 c1 = %d\n", c1) //输出33
	return c1
}

// 从结果看出同变量名的局部变量和全局变量，用一个作用域内优先使用局部变量

// 函数变量（函数作为值）
func FuncVariable() {
	result := StringToLower("AbcdefGHujklMNOPqrstUVWxyz", processCase)
	fmt.Println(result)
	result = StringToLower2("AbcdefGHujklMNOPqrstUVWxyz", processCase)
	fmt.Println(result)
}

// 处理字符串，奇数偶数一次显示为大小写
func processCase(str string) string {
	result := ""
	for i, value := range str {
		if i%2 == 0 {
			result += strings.ToUpper(string(value))
		} else {
			result += strings.ToLower(string(value))
		}
	}
	return result
}
func StringToLower(str string, f func(string) string) string {
	fmt.Printf("%T\n", f)
	return f(str) // f(str)就是processCase(str) f是函数类型，所以能f(参数)这样写
}

type caseFunc func(string) string //声明了一个函数类型，caseFunc是一种新类型

func StringToLower2(str string, f caseFunc) string {
	fmt.Printf("%T\n", f) // 打印变量f的类型
	return f(str)
}

// 函数变量案例二
type processFunc func(int) bool //声明一个函数类型
func FuncVariable2() {
	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd) // 函数当作值来传递
	fmt.Println("奇数元素： ", odd)
	even := filter(slice, isEven)
	fmt.Println("偶数元素： ", even)
}

// 判断元素是否为偶数
func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

// 判断元素是否为偶数
func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

// 根据函数来处理切片，根据元素技术偶数分组，返回新的切片
func filter(slice []int, f processFunc) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}
