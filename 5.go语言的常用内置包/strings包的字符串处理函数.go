package main

import (
	"fmt"
	"strings"
	"unicode"
)

// 常用字符串检索方法
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

// 返回字符串中满足函数f(r)==true的字符最后一次出现的位置（返回的是字节索引
func TestLastIndexFunc() {
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c) // 判断是否是汉字
	}
	fmt.Println(strings.LastIndexFunc("Hello,世界", f))       // 9
	fmt.Println(strings.LastIndexFunc("Hello,world中国人", f)) // 17
}

// 获取文件后缀名
func GetFileSuffix(fileName string) string {
	arr := strings.Split(fileName, ".") // 以.分割字符串，返回一个切片
	return arr[len(arr)-1]
}

// 分割字符串
// func Fields(s string) []string // 按照空格分割字符串，返回一个切片
// func FieldsFunc(s string, f func(rune) bool) []string // 按照函数f分割字符串，返回一个切片
// func Split(s, sep string) []string // 按照指定的分隔符分割字符串，返回一个切片
// func SplitAfter(s, sep string) []string // 按照指定的分隔符分割字符串，返回一个切片，保留分隔符
// func SplitN(s, sep string, n int) []string // 按照指定的分隔符分割字符串，返回一个切片，最多分割n次
// func SplitAfterN(s, sep string, n int) []string // 按照指定的分隔符分割字符串，返回一个切片，最多分割n次，保留分隔符

func SplitString() {
	TestFields()
	TestFieldsFunc()
	TestSplitAfterN()
	TestSplit()
	TestSplitAfter()
	TestSplitN()
}

// 将字符串按照空格分割，返回一个切片
func TestFields() {
	fmt.Println(strings.Fields(" abc 123 ABC xyz XYZ")) // [abc 123 ABC xyz XYZ]
}

// 将字符串以满足函数f(r)==true的字符分割，返回一个切片
func TestFieldsFunc() {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) // 判断是否是字母或数字
	} // 不是字母或数字的字符都作为分隔符
	fmt.Println(strings.FieldsFunc(" abc@123*ABC&xyz%XYZ", f)) // [abc 123 ABC xyz XYZ]
}

// 将字符串以sep分割，分割后字符最后去掉sep
func TestSplit() {
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))                        // ["a" "b" "c"]
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a ")) // ["a man " "plan " "canal panama"]
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))                         // [" " "x" "y" "z" " "]
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))            // [""]
}

// 将字符串s以sep作为分隔符分割，分割后字符最后附上sep,n决定返回的切片数
func TestSplitN() {
	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2)) // ["a" "b,c"]
	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 1)) // ["a,b,c"]
}

// 将字符串s以sep作为分隔符分割，分割后字符最后附上sep
func TestSplitAfter() {
	fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ",")) // ["a," "b," "c"]
}

// 将字符串s以sep作为分隔符分割，分割后字符最后附上sep,n决定返回的切片数
func TestSplitAfterN() {
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 2)) // ["a," "b,c"]
}

// 大小写转化
// func Title(s string) string // 将字符串s的首字母大写，其余小写
// func ToLower(s string) string // 将字符串s转化为小写
// func ToLowerSpecial(s string, caseMapping unicode.SpecialCase) string // 将字符串s转化为小写，使用指定的大小写映射规则
// func ToTitle(s string) string // 将字符串s转化为大写
// func ToTitleSpecial(s string, caseMapping unicode.SpecialCase) string // 将字符串s转化为大写，使用指定的大小写映射规则
// func ToUpper(s string) string // 将字符串s转化为大写
// func ToUpperSpecial(s string, caseMapping unicode.SpecialCase) string // 将字符串s转化为大写，使用指定的大小写映射规则

func UpLower() {
	TestTitle()
	TestToTitle()
	TestToLower()
	TestToUpper()
}

// 将字符串s每个单词的首字母大写，其余小写
func TestTitle() {
	fmt.Println(strings.Title("her name is steven")) // Her Name Is Steven
}

// 将字符串s转化为大写
func TestToTitle() {
	fmt.Println(strings.ToTitle("louD noises")) // LOUD NOISES
}

// 将字符串s转化为小写
func TestToLower() {
	fmt.Println(strings.ToLower("Gopher")) // gopher
}

// 将字符串s转化为大写
func TestToUpper() {
	fmt.Println(strings.ToUpper("Gopher")) // GOPHER
}

// 修建字符串
// func Trim(s string, cutset string) string // 去掉字符串s首尾的cutset中的字符
// func TrimFunc(s string, f func(rune) bool) string // 去掉字符串s首尾满足函数f(r)==true的字符
// func TrimLeft(s string, cutset string) string // 去掉字符串s左边的cutset中的字符
// func TrimLeftFunc(s string, f func(rune) bool) string // 去掉字符串s左边满足函数f(r)==true的字符
// func TrimPrefix(s, prefix string) string // 去掉字符串s的前缀prefix
// func TrimRight(s string, cutset string) string // 去掉字符串s右边的cutset中的字符
// func TrimRightFunc(s string, f func(rune) bool) string // 去掉字符串s右边满足函数f(r)==true的字符
// func TrimSpace(s string) string // 去掉字符串s首尾的空格
// func TrimSuffix(s, suffix string) string // 去掉字符串s的后缀suffix
func TrimString() {
	TestTrim()
	TestTrimFunc()
	TestTrimLeft()
	TestTrimLeftFunc()
	TestTrimPrefix()
	TestTrimRight()
	TestTrimRightFunc()
	TestTrimSpace()
	TestTrimSuffix()
}

// 将字符串s首尾的cutset中的字符去掉
func TestTrim() {
	fmt.Println(strings.Trim(" steven wang ", " ")) // steven wang
}

// 将字符串s首尾满足函数f(r)==true的字符去掉
func TestTrimFunc() {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) // 判断是否是字母或数字
	}
	fmt.Println(strings.TrimFunc("!@#$%steven wang%$#@!", f)) // steven wang
}

// 将字符串s左边的cutset中的字符去掉
func TestTrimLeft() {
	fmt.Println(strings.TrimLeft(" steven wang ", " ")) // steven wang
}

// 将字符串s左边满足函数f(r)==true的字符去掉
func TestTrimLeftFunc() {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) // 判断是否是字母或数字
	}
	fmt.Println(strings.TrimLeftFunc("!@#$%steven wang%$#@!", f)) // steven wang%$#@!
}

// 将字符串s右边的cutset中的字符去掉
func TestTrimRight() {
	fmt.Println(strings.TrimRight(" steven wang ", " ")) // steven wang
}

// 将字符串s右边满足函数f(r)==true的字符去掉
func TestTrimRightFunc() {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) // 判断是否是字母或数字
	}
	fmt.Println(strings.TrimRightFunc("!@#$%steven wang%$#@!", f)) // !@#$%steven wang
}

// 将字符串s首尾的空格去掉
func TestTrimSpace() {
	fmt.Println(strings.TrimSpace(" \t\n a lone gopher \n\t\r\n")) // a lone gopher
}

// 将字符串s的前缀prefix去掉
func TestTrimPrefix() {
	var s = "Goodbye, world!"
	s = strings.TrimPrefix(s, "Goodbye")
	fmt.Println(s) // , world!
}

// 将字符串s的后缀suffix去掉
func TestTrimSuffix() {
	var s = "Helllo,goodbye,etc"
	s = strings.TrimSuffix(s, "goodbye,etc")
	fmt.Println(s) // Helllo,
}

// 比较字符串
// func Compare(a,b string)int // 按字典顺序比较a和b，如果a==b返回0，如果a<b返回-1，如果a>b返回1
// func EqualFold(a,b string)bool // 比较a和b是否相等，忽略大小写
// func Repeat(s string, count int) string // 将字符串s重复count次
// func Replace(s, old, new string, n int) string // 将字符串s中的old替换为new，n表示替换的次数，如果n<0表示全部替换
// func join(a []string, sep string) string // 将字符串切片a中的元素用sep连接起来，返回一个字符串
func CompareString() {
	TestCompare()
	TestEqualFold()
	TestRepeat()
	TestReplace()
	TestJoin()
}

// 按字典顺序比较a和b，如果a==b返回0，如果a<b返回-1，如果a>b返回1
func TestCompare() {
	fmt.Println(strings.Compare("abc", "bcd")) // -1
	fmt.Println("abc" < "bcd")                 // true
}

// 比较a和b两个UTF-8字符串是否相等，忽略大小写
func TestEqualFold() {
	fmt.Println(strings.EqualFold("Go", "go")) // true
}

// 将字符串s重复count次
func TestRepeat() {
	fmt.Println("g" + strings.Repeat("o", 8) + "le") // gooooooooogle
}

// 替换字符串s中old字符串为new字符串，n表示替换的次数，如果n<0表示全部替换
func TestReplace() {
	fmt.Println(strings.Replace("王老大 王老二 王老三", "王", "李", 2))  // 李老大 李老二 王老三
	fmt.Println(strings.Replace("王老大 王老二 王老三", "王", "李", -1)) // 李老大 李老二 李老三
}

// 将a中的所有字符串连接成一个字符串，使用字符串sep作为分隔符
func TestJoin() {
	s := []string{"abc", "ABC", "123"}
	fmt.Println(strings.Join(s, ",")) // abc,ABC,123
	fmt.Println(strings.Join(s, ""))  // abcABC123
}
