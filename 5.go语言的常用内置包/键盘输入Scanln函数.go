package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 键盘输入
func Scanln() {
	username := ""
	age := 0
	fmt.Scanln(&username, &age)
	fmt.Println("账号信息为：", username, age)
}

// 随机数+键盘输入案例-----猜数字游戏
func Game() {
	playGame()
}
func playGame() {
	//获取随机数
	target := generateRandNum(10, 100)
	fmt.Println("请输入随机数：")
	fmt.Println("-----------------")

	// 记录猜测次数
	count := 0
	for {
		count++
		yourNum := 0
		fmt.Scanln(&yourNum)

		if yourNum < target {
			fmt.Println("小了X")
		} else if yourNum > target {
			fmt.Println("大了X")
		} else {
			fmt.Println("正确！")
			fmt.Printf("您一共猜了%d次!\n", count)
			fmt.Println("----------------")
			playGame()
		}
		// 错误提示
		alterInfo(count, target)
	}
}

// 错误提示
func alterInfo(count, target int) {
	if count >= 6 {
		fmt.Printf("您一共猜了%d次都没有猜中!", count)
		fmt.Println("正确数字:", target)
		fmt.Println("------------")
		playGame()
	}
}

// 随机生成数
func generateRandNum(min, max int) int {
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r1.Intn(max-min+1) + min
}
