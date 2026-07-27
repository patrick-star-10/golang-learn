package main

import (
	"fmt"
)

// map是无序的，每次打印都会不一样，不能通过index获取，只能通过key获取
// map的value可以是任意类型，map和切片一样是引用类型
// 声明：var 变量名 map[key类型] value 类型
// 使用make函数：变量名:= make(map(key类型)value类型)

// map的初始化赋值和遍历
func Map() {
	var country = map[string]string{
		"China":  "Beijing",
		"Japan":  "Tokyo",
		"India":  "New Delhi",
		"France": "Paris",
		"Italy":  "Romr",
	}
	fmt.Println(country)
	// 短变量声明初始化方式
	rating := map[string]float64{"c": 5, "Go": 4.5, "Python": 4.5, "C++": 3}
	fmt.Println(rating)
	// 创建map后再赋值
	countryMap := make(map[string]string)
	countryMap["China"] = "Beijing"
	countryMap["Japan"] = "Tokyo"
	countryMap["India"] = "New Delhi"
	countryMap["France"] = "pairs"
	countryMap["Italy"] = "Rome"
	// 遍历map(无序)
	// 1. key value
	for key, value := range countryMap {
		fmt.Println("国家", key, "首都", value)
	}
	fmt.Println("-----------------")
	// 2.value
	for _, value := range countryMap {
		fmt.Println("国家", "首都", value)
	}
	fmt.Println("-----------------")
	// 3.key
	for key := range countryMap {
		fmt.Println("国家", key, "首都", countryMap[key])
	}
}

// 查看元素集合:通过map[key]获取对应的value值，当key不存在时会得到value的默认值
func check() {
	countryMap := make(map[string]string)
	countryMap["China"] = "Beijing"
	countryMap["Japan"] = "Tokyo"
	countryMap["India"] = "New Delhi"
	countryMap["France"] = "pairs"
	countryMap["Italy"] = "Rome"
	// 查看元素是否在map里存在
	// map[key]会返回两个值
	value, ok := countryMap["England"]
	fmt.Printf("%q\n", value)
	fmt.Printf("%T,%v\n", ok, ok)
	if ok {
		fmt.Println("首都:", value)
	} else {
		fmt.Println("首都信息未检索到！")
	}
	// 或者
	if value, ok := countryMap["USA"]; ok {
		fmt.Println("首都:", value)
	} else {
		fmt.Println("首都信息未检索到！")
	}
}

// delete()函数
// delete(map,key)用于删除集合的某个元素
// delete函数方法
func Delete() {
	map1 := map[string]string{
		"element":   "div",
		"width":     "100px",
		"height":    "200px",
		"border":    "solid",
		"backgroud": "none",
	}
	// 根据key删除map中的某个元素
	fmt.Println("删除前:", map1)
	if _, ok := map1["background"]; ok {
		delete(map1, "background")
	}
	fmt.Println("删除后；", map1)
	// 清空map
	map1 = make(map[string]string)
	fmt.Println("清空后：", map1)
}
