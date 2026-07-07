package main

import "fmt"

// 算数运算符
/*
	+	加法
	-	减法
	*	乘法
	/	除法
	%	取模
	++	自增
	--	自减
*/

func arithmetic() {
	var a int = 21
	var b int = 10
	var c int
	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c) // 输出: 31
	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n", c) // 输出: 11
	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c) // 输出: 210
	c = a / b
	fmt.Printf("第四行 - c 的值为 %d\n", c) // 输出: 2
	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", c) // 输出: 1
	a++
	fmt.Printf("第六行 - a 的值为 %d\n", a) // 输出: 22
	a--
	fmt.Printf("第七行 - a 的值为 %d\n", a) // 输出: 21
}

// 关系运算符
/*
	==	等于  检查两个值是否相等，如果相等返回 true，否则返回 false
	!=	不等于  检查两个值是否相等，如果不相等返回 true，否则返回 false
	>	大于  检查左边的值是否大于右边的值，如果是返回 true，否则返回 false
	<	小于  检查左边的值是否小于右边的值，如果是返回 true，否则返回 false
	>=	大于等于  检查左边的值是否大于或等于右边的值，如果是返回 true，否则返回 false
	<=	小于等于  检查左边的值是否小于或等于右边的值，如果是返回 true，否则返回 false
*/
func relation() {
	var a int = 21
	var b int = 10
	if a == b {
		fmt.Printf("第一行 - a 等于 b\n")
	} else {
		fmt.Printf("第一行 - a 不等于 b\n") // 输出: 第一行 - a 不等于 b
	}
	if a < b {
		fmt.Printf("第二行 - a 小于 b\n")
	} else {
		fmt.Printf("第二行 - a 不小于 b\n") // 输出: 第二行 - a 不小于 b
	}
	if a > b {
		fmt.Printf("第三行 - a 大于 b\n") // 输出: 第三行 - a 大于 b
	} else {
		fmt.Printf("第三行 - a 不大于 b\n")
	}
}

// 逻辑运算符
/*
	&&	逻辑与  如果两边的操作数都是 true，则条件为 true，否则为 false
	||	逻辑或  如果两边的操作数有一个为 true，则条件为 true，否则为 false
	!	逻辑非  用于反转操作数的布尔值，如果条件为 true，则逻辑非运算符将其变为 false，否则为 true
*/
func logic() {
	var a bool = true
	var b bool = false
	if (a && b) == true {
		fmt.Printf("第一行 - 条件为 true\n") // 输出: 第一行 - 条件为 true
	}
	if a || b {
		fmt.Printf("第二行 - 条件为 true\n") // 输出: 第二行 - 条件为 true
	}
	// 修改 a 和 b 的值
	a = false
	b = true
	if (a && b) == true {
		fmt.Printf("第三行 - 条件为 true\n")
	} else {
		fmt.Printf("第三行 - 条件为 false\n") // 输出: 第三行 - 条件为 false
	}
	if !(a && b) {
		fmt.Printf("第四行 - 条件为 true\n") // 输出: 第四行 - 条件为 true
	}
}

// 赋值运算符
/*
	=	简单的赋值运算符，将右边的操作数的值赋给左边的操作数
	+=	加法赋值运算符，将左边的操作数与右边的操作数相加，并将结果赋给左边的操作数
	-=	减法赋值运算符，将左边的操作数与右边的操作数相减，并将结果赋给左边的操作数
	*=	乘法赋值运算符，将左边的操作数与右边的操作数相乘，并将结果赋给左边的操作数
	/=	除法赋值运算符，将左边的操作数与右边的操作数相除，并将结果赋给左边的操作数
	%=	取模赋值运算符，将左边的操作数与右边的操作数取模，并将结果赋给左边的操作数
	<<=	左移位赋值运算符，将左边的操作数向左移动右边的操作数指定的位数，并将结果赋给左边的操作数
	>>=	右移位赋值运算符，将左边的操作数向右移动右边的操作数指定的位数，并将结果赋给左边的操作数
	&=	按位与赋值运算符，将左边的操作数与右边的操作数进行按位与运算，并将结果赋给左边的操作数
	^=	按位异或赋值运算符，将左边的操作数与右边的操作数进行按位异或运算，并将结果赋给左边的操作数
	|=	按位或赋值运算符，将左边的操作数与右边的操作数进行按位或运算，并将结果赋给左边的操作数
*/
func assign() {
	var a int = 21
	var c int

	c = a
	fmt.Printf("第一行 - =  运算符实例,c 的值为 %d\n", c) // 输出: 21

	c += a
	fmt.Printf("第二行 - += 运算符实例,c 的值为 %d\n", c) // 输出: 42

	c -= a
	fmt.Printf("第三行 - -= 运算符实例,c 的值为 %d\n", c) // 输出: 21

	c *= a
	fmt.Printf("第四行 - *= 运算符实例,c 的值为 %d\n", c) // 输出: 441

	c /= a
	fmt.Printf("第五行 - /= 运算符实例,c 的值为 %d\n", c) // 输出: 21

	c %= a
	fmt.Printf("第六行 - %= 运算符实例,c 的值为 %d\n", c) // 输出: 0

	c <<= 2
	fmt.Printf("第七行 - <<= 运算符实例,c 的值为 %d\n", c) // 输出: 0

	c >>= 2
	fmt.Printf("第八行 - >>= 运算符实例,c 的值为 %d\n", c) // 输出: 0

	c &= 2
	fmt.Printf("第九行 - &= 运算符实例,c 的值为 %d\n", c) // 输出: 0

	c ^= 2
	fmt.Printf("第十行 - ^= 运算符实例,c 的值为 %d\n", c) // 输出: 2

	c |= 2
	fmt.Printf("第十一行 - |= 运算符实例,c 的值为 %d\n", c) // 输出: 2
}
