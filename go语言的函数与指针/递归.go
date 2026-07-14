package main

import "fmt"

//如果一个函数在内部调用自身，那么这个函数就是递归函数
/*
递归函数满足的两个条件：
1.在一次调用自己时，必须（在某种意义上）更接近于解
2.必须有一个终止处理或计算的准则
*/

// 通过递归来实现阶乘
func recursion() {
	fmt.Println(factorial(5))
	fmt.Println(getMultiple(5))
}
func factorial(n int) int {
	if n == 0 {
		return 1 //递归结束条件
	}
	return n * factorial(n-1) //通过递归实现阶乘
}
func getMultiple(num int) (result int) {
	result = 1
	for i := 1; i <= num; i++ {
		result *= i
	}
	return result //通过循环来实现阶乘
}
