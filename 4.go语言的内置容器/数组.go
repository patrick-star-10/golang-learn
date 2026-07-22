package main

import "fmt"

// 数组的语法 var 变量名 [数组长度] 数据类型
// 数组定义长度后就不能修改
// 如果忽略[]中的数字不设置长度，go语言会根据元素的个数来设置数组的长度，可以忽略声明的长度并且替换为...
// 数组的长度
func NumLength() {
	a := [4]float64{67.7, 89.9, 21, 78}
	b := [...]int{2, 3, 5}
	fmt.Printf("数组a的长度为%d,数组b的长度为%d\n", len(a), len(b))
}

// 遍历数组
func NumRange() {
	a := [4]float64{67.7, 89.9, 21, 78}
	b := [...]int{2, 3, 5}
	//遍历方式1
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i], "\t")
	}
	//遍历方式2
	for _, value := range b {
		fmt.Print(value, "\t")
	}
}

// 多维数组
// var variable_name [size1][size2]....[sizen] variable_type

// 二维数组的嵌套循环输出
func twoDArray() {
	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	fmt.Println(len(a))
	fmt.Println(len(a[0]))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}

// 数组是值类型
// 数组被分配个一个新的变量
func array() {
	a := [...]string{"USA", "China", "India", "Germany", "France"}
	b := a
	b[0] = "Singapore"
	fmt.Println("a:", a)
	fmt.Println("b: ", b)
}
