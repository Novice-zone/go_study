package main

import "fmt"

// 函数做参数
//func add(x, y int) int {
//	return x + y
//}
//
//func sub(x, y int) int {
//	return x - y
//}
//func cal(x, y int, op func(int, int) int) int {
//	return op(x, y)
//}
//
//// 函数做返回值
//func do(s string) (func(int, int) int, error) {
//	switch s {
//	case "+":
//		return add, nil
//	case "-":
//		return sub, nil
//	default:
//		err := errors.New("无法识别的操作符")
//		return nil, err
//	}
//}
//func main() {
//	ret := cal(10, 20, add)
//	fmt.Println(ret)
//}

// 匿名函数（没有函数名）
// 匿名函数需要保存到某个变量或者作为立即执行函数
//多用于实现回调函数和闭包
//func main() {
//	//保存到变量
//	add := func(x, y int) {
//		fmt.Println(x + y)
//	}
//	add(10, 20)
//
//	//自执行函数：匿名函数定义完加() 直接执行
//	func(x, y int) {
//		fmt.Println(x + y)
//	}(10, 20)
//}

//func a() func() {
//	name := "小涛"
//	return func() {
//		fmt.Println("hello", name)
//	}
//}
//
//func main() {
//	//闭包=函数+外层变量的引用
//	r := a() //r此时就是一个闭包
//	r() //相当于执行a函数内部的匿名函数
//}

//func a(name string) func() {
//	//name := "小涛"
//	return func() {
//		fmt.Println("hello", name) //先在最内层找name，找不到就去外层，发现外层接收string类型的name，拿来使用
//	}
//}
//func main() {
//	//闭包=函数+外层变量的引用
//	r := a("小涛")
//	r() //相当于执行a函数内部的匿名函数
//}

//func makeSuffixFunc(suffix string) func(string) string {
//	return func(name string) string {
//		if !strings.HasSuffix(name, suffix) {
//			return name + suffix
//		}
//		return name
//	}
//}
//
//func calc(base int) (func(int) int, func(int) int) {
//	add := func(i int) int {
//		base += i
//		return base
//	}
//
//	sub := func(i int) int {
//		base -= i
//		return base
//	}
//	return add, sub
//}
//
//func main() {
//	//	//闭包=函数+外部环境（变量，参数），在这里是匿名函数
//	//	r := makeSuffixFunc(".docx")
//	//	//ret := r("homework")
//	//	//fmt.Println(ret)
//	//	fmt.Println(r("homework"))
//	//}
//	x, y := calc(50)
//	ret1 := x(100)
//	fmt.Println(ret1)
//	ret2 := y(25)
//	fmt.Println(ret2)
//}

//	func funcA() {
//		fmt.Println("func A")
//	}
//
// // recover()必须搭配defer使用。
// // defer一定要在可能引发panic的语句之前定义。
//
//	func funcB() {
//		defer func() {
//			err := recover()
//			//如果程序出出现了panic错误,可以通过recover恢复过来
//			if err != nil {
//				fmt.Println("recover in B")
//			}
//		}()
//		panic("panic in B")
//	}
//
//	func funcC() {
//		fmt.Println("func C")
//	}
//
//	func main() {
//		funcA()
//		funcB()
//		funcC()
//	}
//func main() {
//	//写一个程序，统计一个字符串中每个单词出现的次数。
//	//比如：“how do you do"中how=1 do=2 you=1。
//	//1.定义一个map[string]int
//	var s = "how do you do"
//	var wordCount = make(map[string]int, 10)
//	//2.分割字符串s，拆分为一个个单词
//	words := strings.Split(s, " ")
//	//3.遍历words，统计
//	//for _, word := range words {
//	//	v, ok := wordCount[word]
//	//	if ok {
//	//		//map中有这个单词
//	//		wordCount[word] = v + 1
//	//	} else {
//	//		wordCount[word] = 1
//	//	}
//	//}
//	for _, word := range words {
//		wordCount[word]++
//	}
//	for k, v := range wordCount {
//		fmt.Println(k, v)
//	}
//}

//
////匿名函数：没有函数名的函数
////函数做返回值只能返回匿名函数
//
////使用匿名函数 1.保存到变量 2.立即执行
//
//func main() {
//	//1.保存到变量
//	add := func(x, y int) {
//		fmt.Println(x + y)
//	}
//	add(10, 20)
//
//	//2.立即执行/自执行函数 定义完直接加()
//	func(x, y int) {
//		fmt.Println(x + y)
//	}(50, 50)
//
//}

// 闭包
// 闭包=函数+引用环境
// 闭包 = 嵌套在另一个函数（外层函数）内部的函数（内层函数）
// + 该内层函数引用了外层函数的变量（引用形式）
// + 这个变量没有通过内层函数的参数传入
func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	f := adder()
	fmt.Println(f(1))
	//fmt.Println(f) 不加括号就是一个地址
}
