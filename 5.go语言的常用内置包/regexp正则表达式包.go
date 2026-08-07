package main

import (
	"fmt"
	"regexp"
)

// 正则表达式就是由元字符组成的一种字符串匹配模式，可以实现对文本内容解析，校验，替换
// 正则表达式中主要元字符：
/*

\  将下一下字符标记为一个特殊字符，或定义一个原义字符，或一个向后引用，或一个八进制转义值
^  匹配输入字符串的开始位置
.  匹配除\n之外的任何单个字符
$  匹配输入字符串的结束位置
x|y  匹配x或y
*  匹配前面的子表达式0次或多次，zo*匹配z以及zoo,等价于{0,}
+  匹配前面的子表达式1次或者多次
？ 匹配前面的子表达式0次或1次
{n}  例如，o{2}匹配food中的两个o
{n,m}  最少匹配n次，最多匹配m次
[xyz]  匹配所包含的任意一个字符
[^xyz]  负值字集合。匹配未包含的任意字符
[a-z]  字符范围。匹配指定范围内的任意字符
[^a-z]  负值字符范围。匹配不在指定范围内的任意字符
\b  匹配一个单词边界，也就是指单词和空格间的位置
\B  匹配非单词边界。er\B能匹配verb中的er,但不能匹配never中的er
\cx 匹配由x指明的控制字符
\d  匹配一个数字。等价于[0-9]
\D  匹配一个非数字。等价于[^0-9]
\f  匹配一个换页符
\n  匹配一个换行符
\r  匹配一个回车符
\s  匹配任何空白字符，包括空格，制表符，换页符
\S  匹配任何非空白字符
\t  匹配一个制表符
\v  匹配匹配一个垂直制表符
\w  匹配包括下画线的任何单词字符
\W  匹配任何非单词字符
\num  匹配num,num是一个正整数
\xn  匹配n,其中n是为十六进制转义值
\un  匹配n,其中n是一个用4个十六进制数子表示的unicode字符
(pattern)  匹配括号内pattern所代表的表达式，是组成匹配
(?=pattern)  正向预查
(?!pattern)  负向预查
*/

/*
regexp包中核心函数以及方法介绍
1.检查正则表达式与字节数组是否匹配
func Match(pattern string,b[]byte)(matched bool,err error)
flag,_ := regexp.Match("^\\d{6,15}$",[]byte("123456789"))
返回结果：true
flag,_ := regexp.Match("^\\d{6,7}$",[]byte("123456789"))
返回结果：false

2.检查正则表达式与字符串是否匹配
func MatchString(pattern string,s string)(matched bool,err error)
flag,_ := regexp.MatchString("\\d{6,15}$","123456789")
返回结果：true
flag,_ := regexp.MatchString("^\\d{6,7}$",[]byte("123456789"))
返回结果：false

3.将正则表达式字符串编译成正则表达式对象
func Compile(expr string)(*Regexp,error)
myRegexp,_ := regexp.Compile("\\d{6}\\D{2}$")

4.MustCompile()用法同Compile(),但是不返回error
func MustComoile(str string)*Regexp
Myregexp := regexp.MustCompile("[\u4e00-\u9fa5]+$")

5.判断Regexp正则对象是否与给定的字节数组匹配
func (re *Regexp) Match(b []byte) bool
MyRegexp := regexp.MustCompile("^[\u4e00-\u9fa5]+$")
flag = MyRegexp.Match([]byte("一丁丐"))
返回结果：true

6.判断Regexp正则对象是否与给定的字符串匹配
func (re *Regexp) MatchString(s string) bool
MyRegexp,_ := regexp.Compile("\\d{6}\\D{2}$")
flag = RegExp.MatchString("123456ab")
返回结果：true

7.ReplaceALl()将src中符合正则表达式的部分全部替换成指定内容
func (re *Regexp) ReplaceAll(src,repl[]byte)[]byte
text := "将字符串123中符合正则表达式的内容3 4 5 全部替换成 56 78 指定的内容"
Myregexp := Regexp.MustCompile("[\\d\\s]+")
result := string(MyRegexp.ReplaceAll([]byte(text),[]byte("")))
本案例中将返回结果转位string,结果为：“将字符串中符合正则表达式的内容全部替换成指定内容”

8.将字符串按照正则表达式分割成子字符串组成的切片，如果切片长度超过指定参数n,则不再分割
func (re *Regexp) Split(s string,n int)[]string
text := "第一部分#第二部分##第三部分###第四部分#第五部分##第六部分"
MyRegexp := regexp.MustCompile("#+")
arr := MyRegexp.Split(text,5)
返回结果：[第一部分 第二部分 第三部分 第四部分 第五部分##第六部分]
*/

// 通过一个案例演示部分正则表达式包的函数
func Regexp() {
	testRegexp()
}
func testRegexp() {
	//1.Match(pattern string,b[]byte)(matched bool,err,error)检查正则表达式是否与字节切片符合
	flag, _ := regexp.Match("^\\d{6,7}$", []byte("123456789"))
	fmt.Println(flag) // false
	//2.MatchString(pattern string,s string)(matched bool,err,error)
	flag, _ = regexp.MatchString("\\d{6,7}$", "123456789")
	fmt.Println(flag) // false
	//3.Compile(expr string)(*Regexp,error)
	RegExp, _ := regexp.Compile("\\d{6}\\D{2}$")
	//4.MustCompile(str,string)*Regexp 不返回error
	RegExp2 := regexp.MustCompile("^[\u4e00-\u9fa5]+$")
	//5.Match(b []byte)bool
	flag = RegExp2.Match([]byte("一丁丐"))
	fmt.Println("xxx:", flag) // xxx:true
	//6.MatchString(s string)bool
	flag = RegExp.MatchString("123456ab")
	fmt.Println(flag) // true
	//7.ReplaceAll(src,repl,[]byte)[]byte
	text := "将字符串123中符合正则表达式的内容 3 4 5 全部替换成 56 78 指定的内容"
	regExp3 := regexp.MustCompile("[\\d\\s]+")
	result := string(regExp3.ReplaceAll([]byte(text), []byte("-")))
	fmt.Println("替换后的字符串为：", result) // 替换后的字符串为：将字符串-按照正则表达式-分割成子字符串-组成的切片
	text = "第一部分#第二部分##第三部分###第四部分#第五部分##第六部分"
	MyRegexp := regexp.MustCompile("#+")
	arr := MyRegexp.Split(text, 5)
	fmt.Println(arr)
}
