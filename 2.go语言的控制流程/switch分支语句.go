package main

import "fmt"

// switch语句的语法格式
/*
	switch 表达式 {
	case 表达式1:
		代码块1
	case 表达式2:
		代码块2
	default:
		代码块3
	}
*/

// switch语句的执行流程：
// 1. 计算switch表达式的值
// 2. 将switch表达式的值与每个case表达式的值进行比较，找到匹配的case表达式
// 3. 执行匹配的case表达式对应的代码块，如果没有匹配的case表达式，则执行default对应的代码块
// 4.Go 默认执行完一个 case 就退出 switch
// 5.Go 里 switch 不能单独存在，必须写在函数体内
// 判断学生成绩
func Switch() {
	grade := " "
	score := 78.5
	switch { // switch后面没有表达式，表示每个case都是一个布尔表达式
	case score >= 90:
		grade = "A"
	case score >= 80:
		grade = "B"
	case score >= 70:
		grade = "C"
	case score >= 60:
		grade = "D"
	default:
		grade = "E"
	}
	fmt.Printf("你的等级是: %s \n", grade) // 输出: 你的等级是: C
	fmt.Print("最终评价是:")
	switch grade {
	case "A":
		fmt.Println("优秀") // 不会执行
	case "B":
		fmt.Println("良好") // 不会执行
	case "C":
		fmt.Println("中等") // 执行输出
	case "D":
		fmt.Println("及格") // 不会执行
	default:
		fmt.Println("不及格") // 不会执行
	}
}

// 判断某年某月的天数
func Days() {
	year := 2008
	month := 2
	days := 0
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		days = 31
	case 4, 6, 9, 11:
		days = 30
	case 2:
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			days = 29 // 闰年
		} else {
			days = 28 // 平年
		}
	default:
		fmt.Println("月份输入错误")
	}
	fmt.Printf("%d年%d月有%d天 \n", year, month, days) // 输出: 2008年2月有29天
}

// 类型转换
// 判断interface变量中储存的变量类型
func TypeSwitch(x interface{}) {
	switch i := x.(type) {
	case nil:
		fmt.Printf("x的类型:%T", i)
	case int:
		fmt.Printf("x的类型:%T", i)
	case float64:
		fmt.Printf("x的类型:%T", i)
	case func(int) float64:
		fmt.Printf("x是func(int)型")
	case bool, string:
		fmt.Printf("x是bool或string型")
	default:
		fmt.Printf("未知型")
	}

}
