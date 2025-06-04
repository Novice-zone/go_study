package main

import (
	"fmt"
	"sort"
)

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
	/*//流程控制符号
	//switch
	switch n := 0; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 0, 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}
	age := 20
	switch {
	case age < 22:
		fmt.Println("考虑就业/读研")
	case age > 22 && age < 25:
		fmt.Println("工作能力提升/学历提升")
	case age > 25:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("退休享受")
	default:
		fmt.Println("活着真好")
	}*/
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

func test05() {
	/*//数组的学习
	var a [10]int
	var b [5]string
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)*/
	//a == b 二者是不同类型，不能拷贝
	//初始化1.
	/*var testArray [3]int                        //数组会初始化为int类型的零值
	var numArray = [3]int{1, 2}                 //使用指定的初始值完成初始化
	var cityArray = [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化
	fmt.Println(testArray)                      //[0 0 0]
	fmt.Println(numArray)                       //[1 2 0]
	fmt.Println(cityArray)                      //[北京 上海 深圳]*/
	/*//2.自行推导长度
	var testArray [3]int
	var numArray = [...]int{1, 2}
	var cityArray = [...]string{"北京", "上海", "深圳"}
	fmt.Println(testArray)                          //[0 0 0]
	fmt.Println(numArray)                           //[1 2]
	fmt.Printf("type of numArray:%T\n", numArray)   //type of numArray:[2]int
	fmt.Println(cityArray)                          //[北京 上海 深圳]
	fmt.Printf("type of cityArray:%T\n", cityArray) //type of cityArray:[3]string
	//3.按索引值初始化
	arr := [...]int{0: 9, 5: 1}
	fmt.Println(arr)
	fmt.Printf("type of a:%T\n", arr)*/
	//遍历
	//var arr = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//for _, val := range arr {
	//	fmt.Println(val)
	//}
	//二维数组
	/*a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	for _, v1 := range a {

		fmt.Printf("%s\t", v1)
		for index, v2 := range v1 {
			fmt.Printf("%d:%s ", index, v2)
		}

		fmt.Println()
	}
	//[北京 上海]	0:北京 1:上海
	//[广州 深圳]	0:广州 1:深圳
	//[成都 重庆]	0:成都 1:重庆 */
	//test
	//求数组所有元素的和
	/*a := [...]int{1, 3, 5, 7, 8}
	var sum int = 0
	for _, val := range a {
		sum += val
	}
	fmt.Printf("数组a元素和为：%d\n", sum)
	//找出数组中和为10的两个元素的下标
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == 10 {
				fmt.Printf("{ %d , %d }\n", i, j)
			}
		}
	}*/
}

// 切片(slice)
// 切片要初始化之后才能使用
func test06() {
	/*//直接定义切片
	m := []int{5, 6, 7, 8, 9, 10}
	fmt.Println(m)
	fmt.Printf("%T\n", m)
	a := [5]int{1, 2, 3, 4, 5}
	//基于数组得到切片
	b := a[1:5] //左闭右开 输出[2 3]
	c := a[1:2] //仍输出[2]

	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	//切片再次切片
	d := b[0:len(b)]
	fmt.Println(d)
	fmt.Printf("%T\n", d)
	//make函数构造切片 make(切片类型，切片长度，指向数组的长度（不写默认和切片一样长）)
	e := make([]int, 5, 10)
	fmt.Println(e)
	fmt.Printf("%T\n", e)*/
	/*// nil 切片的默认值（就像int的默认值是0一样）长度 容量均为0
	//不是所有长度和容量都为0的切片都是nil 因此不能用切片是否为nil判断切片长度为0
	var a []int     //声明int类型切片
	var b = []int{} //声明并初始化
	c := make([]int, 0)
	fmt.Println(a, len(a), cap(a))
	fmt.Println(b, len(b), cap(b))
	fmt.Println(c, len(c), cap(c))
	if a == nil {
		fmt.Println("a==nil")
	}
	if b == nil {
		fmt.Println("b==nil")
	}
	if c == nil {
		fmt.Println("b==nil")
	}*/
	//切片的赋值拷贝
	/*a := make([]int, 3) //[0 0 0]
	b := a              //共用一个底层数组
	b[0] = 100
	fmt.Println(a)
	fmt.Println(b)*/
	/*// 切片的遍历 for / for range
	a := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}
	fmt.Println()
	for index, value := range a {
		fmt.Println(index, value)
	}*/
	// 切片的扩容
	//var a []int // 此时并没有申请内存
	//a[0] = 100
	//fmt.Println(a)// 越界
	/*var a []int       // 此时并没有申请内存
	a = append(a, 10) // 必须使用一个变量接收append的返回值
	// 底层数组大小不够扩容时，会变更指向（像C语言的realloc()）
	fmt.Println(a)

	var b []int
	for i := 0; i < 10; i++ {
		b = append(b, i)
		fmt.Printf("%v len:%d cap:%d ptr:%p\n", b, len(b), cap(b), b)
	} //容量 ：1 2 4 4 8 8 8 8 16 16*/
	/*// copy函数 和赋值拷贝（共用一个底层数组）不同，copy只单纯的把切片的数据复制到了另一个切片里
	a := []int{1, 2, 3, 4, 5}
	b := a
	c := make([]int, 5, 5)
	copy(c, b)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println()
	b[0] = 100
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)*/
	// 切片删除元素
	a := []string{"小王", "小冯", "小张", "小何"}
	// 删除索引为2的元素
	a = append(a[:2], a[3:]...)
	fmt.Println(a)
	// 总结一下就是：要从切片a中删除索引为index的元素，操作方法是
	// a = append(a[:index], a[index+1:]...)

	//-----------------------------------------------------
	// 重点：由于切片扩容时可能因容量不够需要更换底层数组，因此最好在make时就把容量写够
	//
}

func test_06() {
	/*//思考下面代码会输出什么？
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)
	// 切片原本为[     ]五个空字符串
	// 当使用append逐步在a后面追加元素，a变为[     0,1,2,3,4]。
	// 由于make创建a切片的时候容量为10，再次追加元素会导致扩容（变为两倍：20）
	// 后续追加完，结果为：[     0,1,2,3,4,5,6,7,8,9]，（长度15，容量20）无需再次扩容*/

	// 请使用内置的sort包对数组var a = [...]int{3, 7, 8, 9, 1}进行排序
	a := [...]int{3, 7, 8, 9, 1}
	// a[:]得到一个切片，把数组a从头切到尾
	sort.Ints(a[:])
	fmt.Println(a)
}

/*func main() {
	test_06()
}*/
