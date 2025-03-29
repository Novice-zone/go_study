package main

import "fmt"

// if判断
func main() {
	//1.基本写法
	//var score = 65
	//if score >= 90 {
	//	fmt.Println("A")
	//} else if score > 75 {
	//	fmt.Println("B")
	//} else {
	//	fmt.Println("C")
	//}
	//2.if判断的特殊写法，score只在if代码块之中生效
	//if score := 65; score >= 90 {
	//	fmt.Println("A")
	//} else if score > 75 {
	//	fmt.Println("B")
	//} else {
	//	fmt.Println("C")
	//}
	//score := 100
	//if score > 95 {
	//	fmt.Println("A+")
	//} else if score > 90 {
	//	fmt.Println("A")
	//}

	/*//1.for循环基本写法
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	//2.省略初始化，但是;不能省略
	i := 1
	for ; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	//3.省略初始化和迭代操作
	j := 5
	for j < 10 {
		fmt.Printf("%d ", j)
		j++
	}
	fmt.Println()
	//4.无限循环
	//for {
	//	fmt.Println("死循环...")
	//}
	//5.break 跳出for循环
	for i := 0; i < 5; i++ {
		fmt.Printf("循环第%d次\n", i+1)
		if i == 3 {
			fmt.Println("跳出循环")
			break
		}
	}
	//6.continue 继续下一次循环
	for i := 0; i <= 5; i++ {
		if i == 3 {
			continue //跳过这次循环
		}
		fmt.Println(i)
	}*/
	//7.for range
	//暂时省略

	//switch case
	//finger := 3
	//if finger == 1 {
	//	fmt.Println("大拇指")
	//}else if fniger==2{
	//	fmt.Println("食指")
	//}//...太长太复杂
	//1.
	//finger := 3
	//switch finger {
	//case 1:
	//	fmt.Println("大拇指")
	//case 2:
	//	fmt.Println("食指")
	//case 3:
	//	fmt.Println("中指")
	//case 4:
	//	fmt.Println("无名指")
	//case 5:
	//	fmt.Println("小拇指")
	//default:
	//	fmt.Println("无效值")
	//}
	//2.case一次判断多个值
	num := 5
	switch num {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8, 0:
		fmt.Println("偶数")
	}
	//3，case中使用表达式
	age := 47
	switch {
	case age >= 18:
		fmt.Println("adult")
	case age < 17:
		fmt.Println("child")
	default:
		fmt.Println("...")

	}
}
