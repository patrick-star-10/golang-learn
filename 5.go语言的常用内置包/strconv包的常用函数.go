package main

import (
	"fmt"
	"strconv"
)

// Parese类函数
// func Atoi(s string)(int,error). 将字符类型转化为整型
// func ParseInt(s string,base int,bitSize int)(i int64,err error) 将字符串解析成数字，base表示进制
// func PraseUint(s string,base int,bitSize int)(uint64,error) 类似PraseInt,但是用于无符号数字
// func ParseFloat(s stirng,bitSize int)(float64,error) 解析一个表示浮点数的字符串并返回其值
// func ParseBool(str string)(bool,error)返回字符串表示的布尔值
func prase() {
	TestAtoi()
	TestParseInt()
	TestParseUint()
	TestParseFloat()
	TestParseBool()
}

// 将字符串类型转化为int类型
func TestAtoi() {
	a, _ := strconv.Atoi("100")    //(_第二个返回值是判断是否转化成功，用_接受代表丢弃返回值)
	fmt.Printf("%T,%v \n", a, a+2) // int,102
	fmt.Println("-----------------")
}

// 解释给定基数(2~36)的字符串s并且返回相应的值i
func TestParseInt() {
	num, _ := strconv.ParseInt("-4e00", 16, 64)
	fmt.Printf("%T,%v \n", num, num) // int64,-19968
	num, _ = strconv.ParseInt("01100001", 2, 64)
	fmt.Printf("%T,%v\n", num, num) // int64,97
	num, _ = strconv.ParseInt("-01100001", 10, 64)
	fmt.Printf("%T,%v\n", num, num) // int64,-1100001
	num, _ = strconv.ParseInt("4e00", 10, 64)
	fmt.Printf("%T,%v\n", num, num) // int64,0
	fmt.Println("----------------------")
}

// ParseUint类似ParseInt,但是用于无符号数字
func TestParseUint() {
	num, _ := strconv.ParseUint("4e00", 16, 64)
	fmt.Printf("%T,%v\n", num, num) // uint64,19968
	num, _ = strconv.ParseUint("01100001", 2, 64)
	fmt.Printf("%T,%v\n", num, num) // uint64,97
	num, _ = strconv.ParseUint("-1100001", 10, 64)
	fmt.Printf("%T,%v\n", num, num) // uint64,0
	num, _ = strconv.ParseUint("4e00", 10, 64)
	fmt.Printf("%T,%v\n", num, num) // uint64,0
	fmt.Println("-------------------------")
}

// ParseFloat将字符串s转化为float类型
func TestParseFloat() {
	pi := "3.1415926"
	num, _ := strconv.ParseFloat(pi, 64)
	fmt.Printf("%T,%v\n", num, num*2) // float64,6.2831852
	fmt.Println("---------------------")
}

// 将字符串转化为bool类型
func TestParseBool() {
	flag, _ := strconv.ParseBool("steven")
	fmt.Printf("%T,%v\n", flag, flag) // bool,false
	fmt.Println("----------------------")
}

// Format类函数（主要功能是将其他类型格式化成字符串）
// func Itoa(i int)string Itoa是Formatint(int64(i),10)的缩写，int转化成string
// func FormatInt(i int64,base int)string 返回给定基数的字符串表示，2<=base<=36
// func FormatUint(i uint64,base int)string 返回给定基数的i的字符串表示，用于无符号数字
// func FormatFLoat(f float64,fmt byte,prec,bitSize int)string 函数将浮点数表示为字符串返回
// func FormatBool(b bool)string FormatBool根据b的值返回true 或 false
func FormatString() {
	TestItoa()
	TestFormatInt()
	TestFormatUint()
	TestFormatFloat()
	TestFormatBool()
}

// int转化成string
func TestItoa() {
	s := strconv.Itoa(199)
	fmt.Printf("%T,%v,长度：%d", s, s, len(s)) // string,199，长度：3
	fmt.Println("----------------")
}

// 给定基数的i的字符串表示
func TestFormatInt() {
	s := strconv.FormatInt(19968, 16)
	s = strconv.FormatInt(-40869, 16)
	fmt.Printf("%T,%v,长度：%d", s, s, len(s)) // string,-9fa5,长度；5
	fmt.Println("----------------")
}

// 返回给定基数i的字符串表示
func TestFormatUint() {
	s := strconv.FormatInt(-19968, 16)
	s = strconv.FormatInt(-40869, 16)       //无符号
	fmt.Printf("%T,%v,长度：%d", s, s, len(s)) // string,9fa5,长度；4
	fmt.Println("----------------")
}

// 将浮点数f转化为字符串
func TestFormatFloat() {
	s := strconv.FormatFloat(3.1415926, 'g', -1, 64)
	fmt.Printf("%T,%v,长度：%d", s, s, len(s)) // string,3.1415926,长度：5
	fmt.Println("----------------")
}

// 返回布尔值
func TestFormatBool() {
	a := false
	s := strconv.FormatBool(a)
	fmt.Printf("%T,%v,长度：%d", s, s, len(s)) // string,false,长度:5
}
