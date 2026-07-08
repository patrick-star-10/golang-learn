package main

import "fmt"

//不需要使用括号将条件包含起来
//大括号{}必须存在，即使只有一条语句
// else 语句必须与 if 语句在同一行，或者 else 语句前面必须有一个闭合的大括号

// 判断奇数偶数
func oddEven() {
	num := 20
	if num%2 == 0 {
		fmt.Println(num, "是偶数") // 输出: 20 是偶数
	} else {
		fmt.Println(num, "是奇数") // 不会执行
	}
}

// 判断学生成绩
func score() {
	score := 88
	if score >= 90 {
		fmt.Println("优秀") // 不会执行
	} else if score >= 80 {
		fmt.Println("良好") // 执行输出
	} else if score >= 70 {
		fmt.Println("中等") // 不会执行
	} else if score >= 60 {
		fmt.Println("及格") // 不会执行
	} else {
		fmt.Println("不及格") // 不会执行
	}
}

// 特殊写法，if语句还有一个变体
/*
	if statement; condition {
	代码块
	}
*/

func oddEven2() {
	if num := 20; num%2 == 0 { // num定义在if语句中，只在if语句的作用域内有效
		fmt.Println(num, "是偶数") // 输出: 20 是偶数
	} else {
		fmt.Println(num, "是奇数") // 不会执行
	}
}
