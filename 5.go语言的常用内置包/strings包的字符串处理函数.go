package main

import (
	"fmt"
	"strings"
	"unicode"
)

// ======== 包含判断（bool）========
// func Contains(s, substr string) bool // 判断字符串 s 是否包含子字符串 substr
// func ContainsAny(s, chars string) bool // 判断字符串 s 是否包含 chars 中任意一个字符
// func ContainsRune(s string, r rune) bool // 判断字符串 s 是否包含 Unicode 字符 r
// func HasPrefix(s, prefix string) bool // 判断字符串 s 是否以 prefix 开头
// func HasSuffix(s, suffix string) bool // 判断字符串 s 是否以 suffix 结尾

// ======== 统计 ========
// func Count(s, sep string) int // 返回字符串 s 中 sep 出现的次数

// ======== 第一次出现（Index）========
// func Index(s, sep string) int // 返回 sep 在 s 中第一次出现的位置
// func IndexAny(s, chars string) int // 返回 chars 中任意字符第一次出现在 s 中的位置
// func IndexByte(s string, c byte) int // 返回字节 c 第一次出现在 s 中的位置
// func IndexFunc(s string, f func(rune) bool) int // 返回满足 f(r)==true 的字符第一次出现的位置
// func IndexRune(s string, r rune) int // 返回 Unicode 字符 r 第一次出现在 s 中的位置

// ======== 最后一次出现（LastIndex）========
// func LastIndex(s, sep string) int // 返回 sep 在 s 中最后一次出现的位置
// func LastIndexAny(s, chars string) int // 返回 chars 中任意字符最后一次出现在 s 中的位置
// func LastIndexByte(s string, c byte) int // 返回字节 c 最后一次出现在 s 中的位置
// func LastIndexFunc(s string, f func(rune) bool) int // 返回满足 f(r)==true 的字符最后一次出现的位置
func Strings() {
	TestContains()
	TestCount()
	TestIndex()
	TestIndexFunc()
	TestLastIndex()
	TestLastIndexFunc()
	res := GetFileSuffix("abc.xyz.lmn.jpg")
	fmt.Println(res)
}

// 判断是否包含子字符串
func TestContains() {
	fmt.Println(strings.Contains("seafood", "foo"))   // true
	fmt.Println(strings.Contains("seafood", "bar"))   // false
	fmt.Println(strings.Contains("seafood", ""))      // true
	fmt.Println(strings.Contains("", ""))             // true
	fmt.Println(strings.Contains("steven王2008", "王")) // true
}

// 判断字符串是否包含另一个字符串的任一字符
func TestContainsAny() {
	fmt.Println(strings.ContainsAny("team", "i"))        // false
	fmt.Println(strings.ContainsAny("failure", "u & i")) // true
	fmt.Println(strings.ContainsAny("foo", ""))          // false
	fmt.Println(strings.ContainsAny("", ""))             // false
}

// 判断字符串是否包含指定的 Unicode 字符
func TestContainsRune() {
	fmt.Println(strings.ContainsRune("一丁亏", '丁'))   // true
	fmt.Println(strings.ContainsRune("一丁亏", 19969)) // false
}

// 返回字符串包含另一个字符的个数
func TestCount() {
	fmt.Println(strings.Count("cheese", "e")) // 3
	fmt.Println(strings.Count("one", ""))     // 4(空字符串存在于每两个字符之间，也存在于开头和结尾)
}

// 判断字符串s是否有前缀字符串
func TestHasPrefix() {
	fmt.Println(strings.HasPrefix("1000phone news", "1000"))  // true
	fmt.Println(strings.HasPrefix("1000phone news", "1000a")) // false
}

// 判断字符串是否有后缀字符串
func TestHasSuffix() {
	fmt.Println(strings.HasSuffix("1000phone news", "news")) // true
	fmt.Println(strings.HasSuffix("1000phone news", "news")) // false
}

// 返回字符串中子字符串第一次出现的位置
func TestIndex() {
	fmt.Println(strings.Index("chicken", "ken")) // 4
	fmt.Println(strings.Index("chicken", "dmr")) // -1（表示未找到）
}

// 返回字符串中的任一unicode码值首次出现的位置
func TestIndexAny() {
	fmt.Println(strings.IndexAny("abcABC120", "教育基地A")) // 3
}

// 返回字符串中字符首次出现的位置
func TestIndexByte() {
	fmt.Println(strings.IndexByte("123abc", 'a')) // 3
}

// 判断字符串是否包含unicode码值
func TestIndexRune() {
	fmt.Println(strings.IndexRune("abcABC120", 'C')) // 4
	fmt.Println(strings.IndexRune("IT培训教育", '教'))    // 4
}

// 返回字符串中满足函数f(r)==true的字符首次出现的位置
func TestIndexFunc() {
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c) // 判断是否是汉字
	}
	fmt.Println(strings.IndexFunc("Hello123中国", f)) // 8
}

// 返回字符串中子字符串最后一次出现的位置
func TestLastIndex() {
	fmt.Println(strings.LastIndex("Steven learn english", "e")) // 15
	fmt.Println(strings.LastIndex("go gopher", "go"))           // 3
	fmt.Println(strings.Index("go gopher", "go"))               // 0
	fmt.Println(strings.LastIndex("go gopher", "rodent"))       // -1
}

// 返回字符串中任意一个unicode码值最后一次出现的位置
func TestLastIndexAny() {
	fmt.Println(strings.LastIndexAny("chicken", "aeiouy")) // 5
	fmt.Println(strings.LastIndexAny("crwth", "aeiouy"))   // -1
}

// 返回字符串中字符最后一次出现的位置
func TestLastIndexByte() {
	fmt.Println(strings.LastIndexByte("abcABCA123", 'A')) // 6
}

// 返回字符串中满足函数f(r)==true的字符最后一次出现的位置
func TestLastIndexFunc() {
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c) // 判断是否是汉字
	}
	fmt.Println(strings.LastIndexFunc("Hello123中国", f)) // 8
}

// 获取文件后缀名
func GetFileSuffix(fileName string) string {
	arr := strings.Split(fileName, ".") // 以.分割字符串，返回一个切片
	return arr[len(arr)-1]
}
