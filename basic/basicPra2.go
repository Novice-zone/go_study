package main

import (
	"fmt"
	"strings"
)
import "math"

//基本数据类型

// 十进制 二进制 八进制 十六进制
// %d    %b    %o    %x
// 10     1010  012   0x0a
func test1() {
	n := 10
	fmt.Printf("%d \n", n)
	fmt.Printf("%b \n", n)
	fmt.Printf("%o \n", n)
	fmt.Printf("%x \n", n)
	fmt.Printf("%p \n", &n)
}

// 浮点型
func test2() {
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
	fmt.Printf("%f \n", math.Pi)
	fmt.Printf("%.2f \n", math.Pi)
	fmt.Printf("%.7f \n", math.Pi)
}

// 复数complex
func test3() {
	var c1 complex64 //实部虚部分别32位
	c1 = 1 + 2i
	var c2 = 2 + 3i
	//var c2 complex128
	fmt.Println(c1, c2)
}

// 布尔值bool
// 默认值只有true和false,不能强转
func test4() {
	var isOk bool = true
	var isFalse bool = false
	fmt.Println(isOk, isFalse)
}

// 字符串
// 常用转义字符
// \n \r \t \' \" \\
func test5() {
	fmt.Printf("哈哈\n我的路径是：D:\\go\\GoWorkstation\\src\\go_study\\basicPra2.go\n")
	fmt.Printf("制表符\t,单引号\\',换行符\n")
	//``反引号不是单引号
	s1 := `
	多行字符串
	在两个反引号之间的内容原样输出
	不会\n转义
	`
	fmt.Println(s1)
	s2 := "这是一个" +
		"多行字符串的示例。" +
		"可以使用加号拼接。"
	fmt.Println(s2)
}

// 字符串常见操作
func test6() {
	/*//字符串长度
	s1 := "hello."
	fmt.Println(len(s1))
	s2 := "你好golang."
	fmt.Println(len(s2))

	//拼接字符串
	fmt.Println(s2 + s1)
	s3 := fmt.Sprintf(s1 + s2)
	fmt.Println(s3)
	s4 := "哈哈," + "这是第二句话," + "拜拜!"
	fmt.Println(s4)
	*/
	/*
		//字符串分割 strings.Split
		s5 := "how are you man how are you"
		fmt.Println(strings.Split(s5, " "))
		fmt.Printf("%T\n", strings.Split(s5, "")) //字符串切片类型

		//判断是否包含 strings.Contains
		//包含返回true 否则返回false
		var isContain bool
		isContain = strings.Contains(s5, "man")
		fmt.Println(isContain)

		//判断前缀 strings.HasPrefix
		//判断后缀 strings.HasSuffix
		//返回true 否则返回false
		fmt.Println(strings.HasPrefix(s5, "how"))
		fmt.Println(strings.HasSuffix(s5, "man!"))

		//判断字串位置
		//第一个  string.Index
		//最后一个string.LastIndex
		fmt.Println(strings.Index(s5, "how"))
		fmt.Println(strings.LastIndex(s5, "how"))
	*/

	//join
	//strings.Join([]string, string)
	s6 := []string{"how", "do", "you", "do"}
	fmt.Println(s6)
	fmt.Println(strings.Join(s6, " "))
}

// 字符
func test7() {
	//byte  uint8别名 ASCII码
	//rune  int32别名
	//编译器会把它们变成真正的类型,byte rune只是用于提高可读性
	var c1 byte = 'a'
	var c2 rune = 'a'
	fmt.Println(c1, c2)
	fmt.Printf("c1:%T c2:%T\n", c1, c2)

	s := "hello go语言"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", s[i])
	} //中文乱码，因为其他语言的一个字符可能不止占用字节

	fmt.Printf("\n")

	//中英文混杂，使用for range
	for _, r := range s {
		fmt.Printf("%c", r)
	}
}
func main() {
	test7()
}
