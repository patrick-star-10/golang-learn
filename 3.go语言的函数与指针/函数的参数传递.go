package main

import "fmt"

// 函数传int型参数
func intparameter() {
	a := 10
	fmt.Printf("1.变量a的内存地址是: %p,值为: %v \n\n", &a, a)
	fmt.Printf("=======int型变量a的内存地址是:%p \n\n", a)
	changeIntVal(a)
	fmt.Printf("2.changeIntVal函数调用后:变量a的内存地址:%p,值为%v \n\n", &a, a)
	changeIntPtr(&a)
	fmt.Printf("3.changeIntVal函数调用后:变量a的内存地址:%p,值为%v \n\n", &a, a)
}
func changeIntVal(a int) {
	fmt.Printf("---------changeIntVal函数内:值参数a的内存地址:%p,值为%v\n", &a, a)
	a = 90
}
func changeIntPtr(a *int) {
	fmt.Printf("---------changeIntPtr函数内:值参数a的内存地址:%p,值为%v\n", &a, a)
	*a = 90
}

// 函数传递slice型参数
func sliceparameter() {
	a := []int{1, 2, 3, 4}
	fmt.Printf("1.变量a的内存地址是:%p,值为:%v\n\n", &a, a)
	fmt.Printf("切片型变量a的地址是:%p \n\n", a)
	changeSliceVal(a)
	fmt.Printf("2.changeSliceVal函数调用后:变量a的内存地址是:%p,值为%v\n\n", &a, a)
	changeSlicePtr(&a)
	fmt.Printf("3.changeSlicePtr函数调用后:变量a的内存地址是:%p,值为%v\n\n", &a, a)
}
func changeSliceVal(a []int) {
	fmt.Printf("------------changeSliceVal函数内:值参数a的内存地址是:%p,值为%v\n", &a, a)
	fmt.Printf("------------changeSlicePtr函数内:值参数a的内存地址是:%p \n", a)
	a[0] = 99
}
func changeSlicePtr(a *[]int) {
	fmt.Printf("------------changeSlicePtr函数内:指针参数a的内存地址是:%p,值为%v\n", &a, a)
	(*a)[1] = 250
}

// 函数传数组
// 数组传参和int传参是一样的原理，无法改变修改原内容数据
// 结构体传参和int传参也是一样的原理
