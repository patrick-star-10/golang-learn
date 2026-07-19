package main

import "fmt"

// 如果一个函数的参数，类型一定，但是个数不定，可以使用可变参数
/*语法格式：
func 函数名（参数名...类型[(返回值列表)]{
	// 函数体
}
*/

// 计算学员考试的总成绩及平均成绩
func score() {
	sum, avg, count := GetScore(90, 82.5, 73, 64.8)
	fmt.Printf("学员共有%d门成绩,总成绩为：%.2f,平均成绩为:%.2f", count, sum, avg)
	fmt.Println()
	scores := []float64{92, 72.5, 93, 74.5, 89, 87, 74}
	sum, avg, count = GetScore(scores...)
	fmt.Printf("学员共有%d门成绩,总成绩为：%.2f,平均成绩为:%.2f", count, sum, avg)
}
func GetScore(scores ...float64) (sum, avg float64, count int) {
	for _, value := range scores {
		sum += value
		count++
	}
	avg = sum / float64(count)
	return
}
