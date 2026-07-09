package main

import "fmt"

// break语句，break语句用来终止当前正在执行的for循环，并且开始执行循环之后的语句
func BreakExample() {
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break
		}
		fmt.Printf("%d", i)
	}
	fmt.Printf("\nline after for loop")
}

// continue语句 跳过当前的循环，而开始执行下一次循环
func ContinueExample() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
}

// goto语句 goto语句可以无条件地转移到程序制定的行
// 借助goto跳转来输出1~50的素数
func PrimeGoto() {
	var C, c int
	C = 1
LOOP:
	for C < 50 {
		C++
		for c = 2; c*c <= C; c++ {
			if C%c == 0 {
				goto LOOP
			}
		}
		fmt.Printf("%d \t", C)
	}
}
