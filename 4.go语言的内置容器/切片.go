package main

import (
	"fmt"
	"strconv"
)

// 切片与数组的定义区别就是切片[]里不需要写内容，切片为可变长度序列
// 声明切片 1: var indentifier [] type
// 用make()函数来创建切片：var slicel [] type = make([]type,len)
// 创建切片时其中capacity(容量)为可选参数：make([]T,length,capacity)

// 切片属性
func slice1() {
	var numbers = make([]int, 3, 5)
	fmt.Printf("%T\n", numbers)                                                   // 输出[]int
	fmt.Printf("len=%d cap=%d slice = %v\n", len(numbers), cap(numbers), numbers) //输出len=3 cap=5 slice = [0 0 0]
}

// 切片语法左闭右开（右边的值取不到）
func slice2() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)
	fmt.Println("numbers==", numbers)
	fmt.Println("numbers[1:4]==", numbers[1:4])
	fmt.Println("numbers[:3]==", numbers[:3])
	fmt.Println("numbers[4:]==", numbers[4:])
	numbers2 := numbers[:2]
	printSlice(numbers2)
	numbers3 := numbers2[2:5]
	printSlice(numbers3)
}
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

// len()和cap()函数
// len()获取长度，cap()获取容量
// 切片使用细节
func sliceCap() {
	arr0 := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	fmt.Println("cap(arr0)=", cap(arr0), arr0)
	s01 := arr0[2:8]
	fmt.Printf("%T\n", s01)
	fmt.Println("cap(s01)=", cap(s01), s01)
	s02 := arr0[4:7]
	fmt.Println("cap(s02)=", cap(s02), s02)
	s03 := s01[3:9]
	fmt.Println("截取s01[3:9]后形成s03:", s03)
	s04 := s02[4:7]
	fmt.Println("截取s02[4:7]后形成s04:", s04)
	s04[0] = "x" //切片是引用类型
	fmt.Print(arr0, s01, s02, s03, s04)
}

// 切片是引用类型
func slice3() {
	a := [4]float64{67.7, 89.8, 21, 78}
	b := []int{2, 3, 5}
	fmt.Printf("变量a-地址:%p,类型:%T,数值:%v,长度:%d\n", &a, a, a, len(a))
	fmt.Printf("变量b-地址:%p,类型:%T,数值:%v,长度:%d\n", &b, b, b, len(b))
	c := a
	d := b
	fmt.Printf("变量c-地址:%p,类型:%T,数值:%v,长度:%d\n", &c, c, c, len(c))
	fmt.Printf("变量d-地址:%p,类型:%T,数值:%v,长度:%d\n", &d, d, d, len(d))
	a[1] = 200
	fmt.Println("a=", a, "c=", c)
	d[0] = 100
	fmt.Println("b=", b, "d=", d)
	// 输出结果可以看到数组是复制的，改变a不会影响c,切片是引用的，改变d会直接改变b
}

// 修改切片数值
func changeslice() {
	arr := [3]int{1, 2, 3}
	// 根据数组截取切片
	nums1 := arr[:]
	nums2 := arr[:]
	fmt.Println("arr=", arr)
	nums1[0] = 100
	fmt.Println("arr=", arr)
	nums2[1] = 200
	fmt.Println("arr=", arr)
	// 由输出可以看到切片是直接引用数组里的数据，改变切片就会改变数组
}

// append()和copy()函数
// append用于追加新元素，如果容量不够会创建新的内存地址来储存元素
// copy会复制切片元素，其中一修改不会影响另一个

func example() {
	fmt.Println("1.-----------------")
	numbers := make([]int, 0, 20)
	printSlices("numbers:", numbers)
	numbers = append(numbers, 1)
	printSlices("numbers:", numbers)
	numbers = append(numbers, 2, 3, 4, 5, 6, 7)
	printSlices("numbers:", numbers)
	fmt.Println("2.-------------------")
	// 追加一个切片
	s1 := []int{100, 200, 300, 400, 500, 600, 700}
	numbers = append(numbers, s1...)
	printSlices("numbers:", numbers)
	fmt.Println("3.----------------------")
	// 切片删除第一个元素
	numbers = numbers[1:]
	printSlices("numbers:", numbers)
	//切片删除最后一个元素
	numbers = numbers[:len(numbers)-1]
	printSlices("numbers:", numbers)
	// 删除中间一个元素
	a := int(len(numbers) / 2)
	fmt.Println("中间数: ", a)
	numbers = append(numbers[:a], numbers[a+1:]...)
	printSlices("numbers:", numbers)
	fmt.Println("4.-------------------")
	// 创建切片number1是之前容量的两倍
	numbers1 := make([]int, len(numbers), cap(numbers)*2)
	//复制numbers的内容到numbers1
	count := copy(numbers1, numbers)
	fmt.Println("复制个数：", count)
	printSlices("numbers1:", numbers1)
	numbers[len(numbers)-1] = 99
	numbers1[0] = 100
	// numbers1与number不存在联系
	printSlices("numbers1:", numbers1)
	printSlices("numbers:", numbers)
}
func printSlices(name string, x []int) {
	fmt.Print(name, "\t")
	fmt.Printf("addr:%p\t len=%d \t cap=%d\t slice=%v\n", x, len(x), cap(x), x)
}

// 对比案例2
func example2() {
	// 思考：使用哪种初始化切片的方式更高效
	var sa []string
	// sa := make([]string,0,20)
	printSliceMsg(sa)
	for i := 0; i < 15; i++ {
		sa = append(sa, strconv.Itoa(i))
		printSliceMsg(sa)
	}
	printSliceMsg(sa)
}
func printSliceMsg(sa []string) {
	fmt.Printf("addr:%p\t len:%d\t cap:%d\t value:%v\n",
		sa, len(sa), cap(sa), sa)
}
