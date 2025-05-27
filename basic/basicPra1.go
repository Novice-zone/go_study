package main

import "fmt"

/*
var
func test() (int, string) {
	return 500, "good"
}

// go语言一条语句不需要分号;结尾
// 函数外语句必须以关键字开始
var salary = 30000 //全局变量
// 变量定义
func main() {
	//标准变量定义,不初始化则都有对应值
	var name string
	var age int
	var sex bool
	fmt.Println(name, age, sex) //变量必须要使用
	//批量定义
	//定义+初始化
	var (
		a string  = "hello"
		b int     = 0
		c float32 = 100.00
	)
	fmt.Println(a, b, c)
	//常用变量定义
	//短变量声明，只能在函数内部使用
	num := "123456789"
	score := 100
	isOk := true
	fmt.Println(num, score, isOk, salary)
	//匿名变量,_用于占位，表示忽略值
	x, _ := test()
	_, y := test()
	fmt.Println(x, y)
}
*/

//常量

//const pi = 3.1415926
//const e = 2.7

//const (
//	pi = 3.14
//	e  = 2.7
//)

//const (
//	n1 = 10
//	n2 //这种情况，n2，n3默认与上面的n1相同
//	n3
//)

// 在 const 常量组里，iota 从 0 开始，每一行递增 1。
//const (
//	n0 = iota //0
//	n1 = iota //1
//	n2 //2
//	n3 //3
//	n4 //4
//	n5 //5
//)

//iota的几个常用示例

//const (
//	n0 = iota
//	n1 //1
//	n2 //2
//	_ //_用于忽略一些值
//	n4 //4
//)

//const (
//	n0 = iota
//	n1 //1
//	n2 //2
//	n3 = 100
//	n4 = iota //4
//	n5        //5
//)
//const n6 = iota //0

// iota定义数量级
//const (
//	_  = iota
//	KB = 1 << (10 * iota)
//	MB = 1 << (10 * iota)
//	GB = 1 << (10 * iota)
//)

const (
	a, b = iota + 1, iota + 2
	c, d
	e, f
)

func main() {
	//fmt.Println(pi, e)
	//fmt.Println(n1, n2, n3)
	//fmt.Println(n0, n1, n2, n3, n4, n5, n6)
	//fmt.Println(KB, MB, GB)
	fmt.Println(a, b, c, d, e, f)
}
