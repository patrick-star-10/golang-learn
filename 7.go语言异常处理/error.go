package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

/*
	error接口
	错误是指程序中出现不正常的情况，从而导致程序无法正常运行。假设尝试打开一个错误文件，文件系统中不存在
	这是一个异常情况，它表示一个错误
	go语言通过内置的错误类型提供了非常简单的错误处理机制，即error接口。定义如下：
	type error interface{
		Error() string
	}
	error本质上是一个接口类型，其中包含Error()方法，错误值可以存储在变量中，通过函数放回。
	它必须是函数返回的最后一个值
	在go语言中处理错误的方式通常是将返回的错误和nil进行比较，nil值表示那样发生错误，而非nil值表示出现错误
	如果不是nil,需打印输出错误
*/

// 使用error接口示例
func ErrorInterface() {
	// 异常情况1
	res := math.Sqrt(-100)
	fmt.Println(res)
	res, err := Sqrt(-100)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	//异常情况2
	res, err = Divide(100, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	//异常情况3打开不存在的文件
	f, err := os.Open("/abc.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f.Name(), "该文件被成功打开！")
	}
}

// 定义平方根运算函数
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("负数不可以获取平方根")
	} else {
		return math.Sqrt(f), nil
	}
}

// 定义除法运算函数
func Divide(dividee float64, divider float64) (float64, error) {
	if divider == 0 {
		return 0, errors.New("出错:除数不可以为0!")
	} else {
		return dividee / divider, nil
	}
}
