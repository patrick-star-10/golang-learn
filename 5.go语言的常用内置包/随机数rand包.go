package main

import (
	"fmt"
	"math/rand"
	"time"
)

// “math/rand"包实现了伪随机数生成器，能够生成整型和浮点型的随机数。
// 使用随机数生成器需要放入种子，可以使用seed()函数生成一个不确定的种子放入碎机函数生成器
// 这样每次运行随机数生成器都会生成不同的序列
// 如果没有在随机数生成器中放入种子，则默认使用具有确定性状态的种子，此时可以理解为种子的值是一个常数1，即seed(1)

/*
rand包中核心方法
func NewSource(seed int64)Source  使用给定的种子创建一个伪随机资源
func New(src Source)*Rand  返回一个使用src产生的伪随机数来生成其他各种分布的随机函数值*Rand
func (r*Rand)Seed(seed int64)  使用给定的seed来初始化生成器到一个确定的状态
func (r*Rand)Int()int  返回一个非负的伪随机int值
func (r*Rand)Intn(n int)int  返回一个取值范围在[0,n]的伪随机int值，如果n<=0会panic
func (r*Rand)Folat64()folat64  返回一个取值范围在[0.0,1.0]的伪随机float64值
*/

/*
获取随机素的几种方式，通过默认的随机数种子获取随机素，具体方法如下：
rand.Int()
rand.Float64()
rand.Init(n) // 获取0-n的随机数
这样总是生成固定的随机数。默认情况下，随机数种子都是1

动态随机数种子生产随机资源，产生随机对象来获取随机数，具体方法如下：
s1 := rand.NewSource(time.Now().UnixNano())
r1 := rand.New(s1)
randnum := r1.Intn(n)  获取0-n随机数

简写形式：动态变化随机数种子来获取随机数，具体形式如下：
(1)获取0～10随机数
rand.Seed(time.Now().UnixNano())
rand.Intn(10)

(2)获取浮点型0.0～1.0随机数
rand.Seed(time.Now().UnixNano())
rand.Float64()

(3)获取m~n的随机数
rand.seed(time.Now().UnixNano())
随机数 = rand.Inta(n-m+1)+m
例如，获取[5,11]的随机数，语法格式如下：
rand.Intn(7)+5
*/

// 案例演示部分rand包的核心函数
func Rand() {
	randTest()
	randAnswer()
}

// 生成随机数
func randTest() {
	fmt.Println(rand.Int())
	fmt.Println(rand.Intn(50))
	fmt.Println(rand.Float64())
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	randnum := r1.Intn(10)
	fmt.Println(randnum)
	rand.Seed(time.Now().UnixNano()) // 不会影响r1的种子 	注：（最新版官方放弃rand.Seed这种修改全局种子的方法）
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Float64())
	num := rand.Intn(7) + 5
	fmt.Println(num)
}

// 随机获取应答
func randAnswer() {
	answers := []string{
		"It is certain",
		"It is decidedly so",
		"without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it yes",
		"Most Likely",
		"Outlook good",
		"Yes",
		"Signs,point to yes",
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randnum := r.Intn(len(answers))
	fmt.Println("随机互答: ", answers[randnum])
}
