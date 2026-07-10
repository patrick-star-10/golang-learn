package main

import "fmt"

// 打印左上直角三角形
func LeftTriangle() {
	lines := 8
	for i := 0; i <= lines; i++ {
		for j := 8; j >= i; j-- {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

// 打印所有水仙花数（三位数，其各位数字的立方之和等于该数）
func flowernumber() {
	var a, b, c int
	for i := 100; i < 1000; i++ {
		a = i / 100
		b = i / 10 % 10
		c = i % 10
		if a*a*a+b*b*b+c*c*c == i {
			fmt.Println(i)
		}
	}

}
