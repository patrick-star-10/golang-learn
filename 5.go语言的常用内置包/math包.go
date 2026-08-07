package main

import (
	"fmt"
	"math"
)

// math包提供了基本的数学常数和数学函数，使用时需要impoet "math"
/*
func IsNaN(f float64)(is bool) 	报告f是否表示一个NaN(Not A Number)值
func Ceil(x float64)float64  返回不小于x的最小整数(浮点数)
func Floor(x float64)float64 	返回不大于X的最小整数(浮点数)
func Trunc(x float64)float64 	返回X的整数部分（的浮点值）
func Abs(x float64)float64 	返回X的绝对值
func Max(x,y float64)float64  返回X和Y中的较大值
func Mix(x,y float64)float64  返回X和Y中的较小值
func Dim(x,y,float64)float64  返回X-Y和0中的较大值
func Mod(x,y,float64)float64  取余运算，
func Sqrt(x float64)float64  返回X的二次方根
func Cbrt(x float64)float64  返回X的三次方根
func Hypot(p,q float64)  返回Sqrt(p*p+q*q)
func Pow(x,y,float64)float64  返回X^Y
func Sin(x float64)float64  求正弦
func Cos(x float64)float64  求余弦
fun Tan(x float64)float64  求正切
func Log(x float64)float64  求自然对数
func Log2(x float64)  求2为底的对数
func Log10(x float64)float64  求10为底的对数
*/
// math包部分函数演示
func Math() {
	fmt.Println(math.IsNaN(3.4))      //false
	fmt.Println(math.Ceil(1.000001))  // 2
	fmt.Println(math.Floor(1.999999)) // 1
	fmt.Println(math.Trunc(1.999999)) // 1
	fmt.Println(math.Abs(-1.3))       // 1.3
	fmt.Println(math.Max(-1.3, 0))    // 0
	fmt.Println(math.Min(-1.3, 0))    // -1.3
	fmt.Println(math.Dim(-12, -19))   // 7
	fmt.Println(math.Dim(-12, 19))    // 0
	fmt.Println(math.Mod(9, 4))       // 1
	fmt.Println(math.Sqrt(9))         // 3
	fmt.Println(math.Cbrt(8))         // 2
	fmt.Println(math.Hypot(3, 4))     // 5
	fmt.Println(math.Pow(2, 8))       // 256
	fmt.Println(math.Log(1))          // 0
	fmt.Println(math.Log2(16))        // 4
	fmt.Println(math.Log10(1000))     // 3
}
