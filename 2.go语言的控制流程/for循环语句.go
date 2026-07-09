package main

import "fmt"

// 语法结构
/*
	for 初始语句init;条件表达式condition; 结束语句post {
	//	循环体代码
	}
*/

// for循环
func selfadd() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
}

// 省略初始语句
func selfadd2() {
	i := 0
	for ; i <= 10; i++ {
		fmt.Printf("%d", i)
	}
}

// 省略条件表达式（会默认形成无限循环）
func selfadd3() {
	i := 0
	for ; ; i++ {
		if i > 20 {
			break
		}
		fmt.Printf("%d", i)
	}
}

// 语法形式二
func selfadd4() {
	var i int
	for i <= 10 {
		fmt.Print(i)
		i++
	}
}

// 语法形式三
func selfadd5() {
	var i int
	for {
		if i > 10 {
			break
		}
		fmt.Print(i)
		i++
	}
}

// 语法形式四
func for4() {
	str := "123abcABC你好"
	for i, value := range str {
		fmt.Printf("第%d位的ASCII的值=%d,字符是%c \n", i, value, value)
	}
}

// 使用案例
// 求1～100的和
func addsum() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum) //输出5050
}

// 求1～40所有3的倍数的和
func addsum3() {
	sum := 0
	i := 1
	for i < 40 {
		if i%3 == 0 {
			sum += i
		}
		i += 1
	}
	fmt.Println(sum)
}

// 截竹杆，32米竹杆，每次减少1.5米，求至少几次之后剩余长度不足4米
func zhuzi() {
	sum := 0
	for i := 32.0; i >= 4; i -= 1.5 {
		sum += 1
	}
	fmt.Println(sum)
}

