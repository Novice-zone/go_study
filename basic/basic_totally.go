package main

import "fmt"

func foo_01() (int, string) {
	return 14, "shuzhi"
}

func test01() {
	//变量
	/*var name string
	var age int
	var isOK bool

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isOK)

	var (
		a int16   = -10
		b uint16  = 10
		c float32 = 15.53
		d bool    = true
	)
	fmt.Println(a, b, c, d)*/

	/*x, _ := foo_01()
	_, y := foo_01()
	fmt.Println("x=", x)
	fmt.Println("y=", y)*/

	//常量

	/*
		//const pi=3.14159
		//const e=2.7182
		const (
			pi = 3.1415
			e  = 2.7182
		)
		fmt.Println(pi, e)
		const (
			n1 = 10
			n2
			n3
		)
		fmt.Println(n1, n2, n3)*/

	//iota
	/*const (
		n0 = iota
		n1
		_
		n3
		n4
	)
	fmt.Println(n0,n1, n3,n4)*/

	/*const (
		n0 = iota
		n1
		n2 = 200
		n3 = iota
		n4
		_
		n6
	)
	fmt.Println(n0, n1, n2, n3, n4, n6)*/

	/*	const (
			a, b = iota + 1, iota + 2
			c, d
			e, _
		)
		fmt.Println(a, b)
		fmt.Println(c, d)
		fmt.Println(e)*/
}

func test02() {
	//基本数据类型
	//数字字面量语法

	//不同格式定义数字
	/*	v1 := 0b001010 //二进制，相当于十进制10
		v2 := 0o000033 //八进制，相当于十进制27
		v3 := 0x00001c //十六进制，相当于十进制22
		fmt.Printf("%b \n", v1)
		fmt.Printf("%d \n", v1)

		fmt.Printf("%o \n", v2)
		fmt.Printf("%d \n", v2)

		fmt.Printf("%x \n", v3)
		fmt.Printf("%d \n", v3)*/

	//控制打印浮点数
	//fmt.Printf("%f\n", math.Pi)
	//fmt.Printf("%.4f\n", math.Pi)
	//
	////布尔值，只有true和false，不允许将整型强制转换为布尔型
	////无法参与数值运算
	//a := true
	//b := false
	//fmt.Println(a)
	//fmt.Println(b)
	//
	////使用转义符打印文件路径
	//fmt.Println("本文件路径为：D:\\go\\GoWrokstation\\src\\go_study\\basic_tatally.go")
	//fmt.Println("c:\\Code\\lesson1\\hehe'") //单引号'不用转义、
	//
	////多行字符串
	///*str := "first\n" +
	//"second\n" +
	//"third\n"*/
	//str := `我这样就可以\直接
	//输出我在代码区写的格式//
	//不用任何修饰”“
	//使用反引号即可`
	//fmt.Println(str)
	//
	////byte和rune(int32)类型
	////普通字符是byte类型，中文等复合字符是rune类型，遍历字符串不能按字节遍历
	//s := "树脂66"
	//for i := 0; i < len(s); i++ {
	//	fmt.Printf("%c ", s[i])
	//}
	//fmt.Println()
	//for _, r := range s {
	//	fmt.Printf("%c ", r)
	//}
	//修改字符串
	//先把字符串转为[]byte或[]rune，按数组形式使用下标进行修改
	//再把类型转换为string类型
	//s1 := "big"
	//// 强制类型转换
	//byteS1 := []byte(s1)
	//byteS1[0] = 'p'
	//fmt.Println(string(byteS1))
	////fmt.Println(byteS1)
	//
	//s2 := "白萝卜"
	//runeS2 := []rune(s2)
	//runeS2[0] = '红'
	//fmt.Println(string(runeS2))

	/*//test
	a := 1
	b := 1.1
	c := true
	d := "树脂"
	fmt.Printf("%T,%d\n", a, a)
	fmt.Printf("%T,%.1f\n", b, b)
	fmt.Printf("%T,%t\n", c, c)
	fmt.Printf("%T,%s\n", d, d)

	str := "hello,go语言,这个是我的测试"
	count := 0
	for _, r := range str {
		if unicode.Is(unicode.Han, r) {
			count++
		}
	}
	fmt.Println("中文个数：%d", count)*/
}

func test03() {
	//运算符
	//test--找出一堆数字(每个数字出现两次，有一个出现一次)里只出现一个的那一个
	//a^0=a
	//a^a=0
	a := 5
	b := 0
	c := a ^ b
	d := a ^ a
	fmt.Println(c)
	fmt.Println(d)
	arr := [9]int{1, 2, 3, 4, 5, 5, 4, 3, 1}
	ret := 0
	for _, r := range arr {
		ret ^= r
	}
	fmt.Printf("单独的数字是：%d", ret)
}

func test04() {
	//流程控制符号
	//test打印99乘法表
	//for i := 1; i < 10; i++ {
	//	for j := 1; j < 10; j++ {
	//		fmt.Printf("%dx%d=%-3d", i, j, i*j)
	//	}
	//	fmt.Println()
	//}
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dx%d=%-3d", i, j, i*j)
		}
		fmt.Println()
	}
}
func main() {
	test04()
}
