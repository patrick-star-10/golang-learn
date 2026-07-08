package main

import "fmt"

/*
	%v 值的默认格式表示(value)
	%+v 在 %v 的基础上，对结构体会添加字段名
	%#v 值的 Go 语法表示
	%T 值的类型的 Go 语法表示(type)
	%t 布尔型
	%d 十进制表示
	%5d 输出宽度至少为 5 个字符

*/

func fmtprint() {
	str := "hello"
	fmt.Printf("%T,%v\n", str, str) // 输出: string,hello
	var a rune = '-'
	fmt.Printf("%T,%v\n", a, a) // 输出: int32,45
}

// 布尔型打印格式
func boolprint() {
	var flag bool
	fmt.Printf("%T,%t \n", flag, flag) // 输出: bool,false
	flag = true
	fmt.Printf("%T,%t \n", flag, flag) // 输出: bool,true
}

// 整型打印格式
/*%b 二进制表示
%o 八进制表示
%c 该值对应的 Unicode 码点所表示的字符
%b 科学计数法表示(无小数点)
*/
func intprint() {
	fmt.Printf("%T,%d \n", 123, 123)   // 输出: int,123
	fmt.Printf("%T,%5d \n", 123, 123)  // 输出: int,  123 (宽度为5，右对齐)
	fmt.Printf("%T,%05d \n", 123, 123) // 输出: int,00123 (宽度为5，右对齐，前面补0)
	fmt.Printf("%T,%b \n", 123, 123)   // 输出: int,1111011 (二进制)
	fmt.Printf("%T,%o \n", 123, 123)   // 输出: int,173 (八进制)
	fmt.Printf("%T,%c \n", 123, 123)   // 输出: int,{ (字符)
}

// 浮点型打印格式
/*%b 二进制表示
%e 科学计数法表示（有6位小数）
%f 小数点表示
*/
func floatprint() {
	fmt.Printf("%b \n", 123.123456)   // 输出: 12312345600000000p-50 (科学计数法)
	fmt.Printf("%f \n", 123.123456)   // 输出: 123.123456 (小数形式)
	fmt.Printf("%.2f \n", 123.123456) // 输出: 123.12 (保留2位小数)
	fmt.Printf("%e \n", 123.123456)   // 输出: 1.231235e+02 (科学计数法)
	fmt.Printf("%g \n", 123.123456)   // 输出: 123.123456 (根据数值大小选择 %e 或 %f)
}

// 复数打印格式
func complexprint() {
	var value complex64 = 2.1 + 12i
	value2 := complex(2.1, 12) //complex(real, imag) 复数构造函数
	fmt.Println(real(value))   // 输出: 2.1 (实部)
	fmt.Println(imag(value))   // 输出: 12 (虚部)
	fmt.Println(value2)        // 输出: (2.1+12i) (复数)
	fmt.Println(value)         // 输出: (2.1+12i) (复数)
}

// 字符串与字节数组的打印格式
/*
	%s 字符串表示
	%q 带引号的字符串表示
	%x 十六进制表示（小写）
	%X 十六进制表示（大写）
*/
func stringprint() {
	arr := []byte{'x', 'y', 'z', 'Z'}
	fmt.Printf("%s \n", "学习区块链") // 输出: 学习区块链
	fmt.Printf("%q \n", "学习区块链") // 输出: "学习区块链" (带引号的字符串表示)
	fmt.Printf("%x \n", "学习区块链") // 输出: e5ad a6e7 9fbae5 8c85 e9bb 91e7 9fbae5 8c85 e9bb 91 (十六进制表示)
	fmt.Printf("%X \n", "学习区块链") // 输出: E5AD A6E7 9FBA E58C 85E9 BB91 E79F BAE5 8C85 E9BB 91 (十六进制表示)
	fmt.Printf("%s \n", arr)     // 输出: xyzZ (字节数组表示)
	fmt.Printf("%q \n", arr)     // 输出: "xyzZ" (带引号的字节数组表示)
	fmt.Printf("%x \n", arr)     // 输出: 78797a5a (十六进制表示)
	fmt.Printf("%X \n", arr)     // 输出: 78797A5A (十六进制表示)
}
