package main

import (
	"fmt"
	"time"
)

// time包提供了时间的显示和测量用过的函数。日历的计算采用的是公历
/*
func Now()Time  返回当前本地时间
func(t Time)Local()time	 将时间转成本地时区，但指向同一时间点
func (t Time)UTC()time 	 将时间转成零区，但指向同一时间点
func Date(year int,month Month,day,hour,min,sec,nsec int,loc*Location)time 	根据指定数值，返回一个时间
func Parse(layout,value string)(Time,error)	 将一个格式化的时间字符串解析成它所代表的时间，就是string类型转time类型
func (t Time)Format(layout string)string 	根据layout指定的格式返回t代表的时间点的格式化文本表示
func (t Time)Uinx()int64	将t表示Unix时间，即从时间点january1,1970 UTC到时间点t所经过的时间（单位秒）
func (t TIme)UnixNano()int64	将t表示为Unix时间，即从时间点january1,1970 UTC到时间点t所经过的时间（单位纳秒）
func (t Time)Equal(u Time)bool 	判断两个时间是否相同，会考虑时区的影响，因此不同时区标准的时间也可以正确比较
func (t Time)Before(u Time)bool  如果t代表的时间点在u之前，返回true，否则返回false
func (t Time)After(u TIme)bool	如果t代表的时间点在u之后，返回true，否则返回false
func (t Time)Date()(year int,month Month,day int) 	返回时间点t对应的年，月，日
func (t Time)Year()int 	返回时间点t对应的年份
func (t Time)Month()Month 	返回时间点t对应那一年对应的月份
func (t Time)Day()int  返回时间点t对应那一月的第几日
func (t Time)Weekday()Weekday 	返回时间点t对应那一周的周几
func (t Time)Clock()(hour,min,sec int) 	返回时间t对应的那一天的时，分，秒
func (t Time)Hour()int 	返回t对应的那一天的第几小时，范围[0,23]
func (t Time)Minute()int 	返回t对应的那一小时的第几分钟，范围[0,59]
func (t Time)Second()int 	返回t对应的那一分钟的第几秒，范围[0,59]
func (t Time)Nanosecond()int  返回t对应的那一秒内的纳秒偏移量，范围[0,999999999]
func (t Time)Sub(u Time)Duration  返回一个时间段t-u
func (d Duration)Hours()float64  将时间段表示为float64类型的小时数
func (d Duration)Minutes()float64  将时间段表示为float64类型的分钟数
func (d Duration)Seconds()float64 	将时间段表示为float64类型的秒数
func (d Duration)Nanoseconds()int64  将时间段表示为int64类型的纳秒数，等价于int64(d)
func (d Duration)String()string  	返回时间段用字符串表示，格式如“72h3m0.5s”
func ParseDuration(s string)(Duration,error) 	解析一个时间段字符串
func (t Time)Add(d Duration)Time  返回时间点t+d
func (t TIme)AddDate(years int,months int,days int)Time  返回增加了给出的年数，月数和天数的时间点
*/
// 通过一个案例演示部分time包的函数
func Time() {
	time1 := time.Now()
	testTime()
	time2 := time.Now()
	fmt.Println(time2.Sub(time1).Seconds())
}
func testTime() {
	t := time.Now()
	fmt.Println("1. ", t)
	fmt.Println("2. ", t.Local())
	fmt.Println("3. ", t.UTC())
	t = time.Date(2026, time.January, 1, 1, 1, 1, 0, time.Local)
	fmt.Printf("4. 本地时间%s,国际统一时间:%s\n", t, t.UTC())
	t, _ = time.Parse("2006-01-02 15:04:05", "2018-07-19 05:47:13")
	fmt.Println("5. ", t)
	fmt.Println("6. ", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("7. ", time.Now().String())
	fmt.Println("8. ", time.Now().Unix())
	fmt.Println("9. ", time.Now().UnixNano())
	fmt.Println("10. ", t.Equal(time.Now()))
	fmt.Println("11. ", t.Before(time.Now()))
	fmt.Println("12. ", t.After(time.Now()))
	year, month, day := time.Now().Date()
	fmt.Println("13. ", year, month, day)
	fmt.Println("14. ", time.Now().Year())
	fmt.Println("15. ", time.Now().Month())
	fmt.Println("16. ", time.Now().Day())
	fmt.Println("17. ", time.Now().Weekday())
	hour, minute, second := time.Now().Clock()
	fmt.Println("18. ", hour, minute, second)
	fmt.Println("19. ", time.Now().Hour())
	fmt.Println("20. ", time.Now().Minute())
	fmt.Println("21. ", time.Now().Second())
	fmt.Println("22. ", time.Now().Nanosecond())
	fmt.Println("23. ", time.Now().Sub(time.Now()))
	fmt.Println("24. ", time.Now().Sub(time.Now()).Hours())
	fmt.Println("25. ", time.Now().Sub(time.Now()).Seconds())
	fmt.Println("26. ", time.Now().Sub(time.Now()).Seconds())
	fmt.Println("27. ", time.Now().Sub(time.Now()).Nanoseconds())
	fmt.Println("28. ", "时间间距:", t.Sub(time.Now()).String())
	d, _ := time.ParseDuration("1h30m")
	fmt.Println("29. ", d)
	fmt.Println("30. ", "交卷时间: ", time.Now().Add(d))
	fmt.Println("31. ", "一年一个月零一天之后的日期:", time.Now().AddDate(1, 1, 1))
}
