package main

import "fmt"

// 打印直角三角形
func RightTriangle() {
	lines := 8
	for i := 0; i < lines; i++ {
		for n := 0; n < 2*i+1; n++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

// 打印九九乘法表
func NineTable() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, j*i)
		}
		fmt.Println()
	}
}

// 使用循环嵌套来输出2～50的素数
func PrimeNumber() {
	var a, b int
	for a = 2; a <= 50; a++ {
		for b = 2; b <= (a / b); b++ { // b <= (a/b)等于b<=根号a
			if a%b == 0 {
				break
			}
		}
		if b > (a / b) { // 当除数已经检查到超过被除数的平方根，仍然没有找到任何可以整除被除数的数，那么这个被除数一定是素数。
			fmt.Printf("%d\t", a)
		}
	}
}
