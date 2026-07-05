package main

import "fmt"

// 布尔型无法直接转换为其他类型，其他类型也无法直接转换为布尔型
func FloatAndIntconvert() {
	chinese := 90
	english := 80.9
	avg := (chinese + int(english)) / 2      //将浮点型转换为整型
	avg2 := (float64(chinese) + english) / 2 //将整型转换为浮点型
	fmt.Printf("%T,%d \n", avg, avg)         // 输出: int,85
	fmt.Printf("%T,%f \n", avg2, avg2)       // 输出: float64,85.450000
}

// 整数转字符串类型
// 不允许直接将整型转换为字符串类型，必须使用strconv包中的Itoa()函数将整型转换为字符串类型
// 但是将整型转换为字符串类型时，转换的是该整型对应的 Unicode 码点所表示的字符
func IntAndStringconvert() {
	a := 97
	x := 19958
	result := string(a)  //将整型转换为字符串类型
	fmt.Println(result)  // 输出: a
	result2 := string(x) //将整型转换为字符串类型
	fmt.Println(result2) // 输出: ䷶
}
