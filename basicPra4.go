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
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	i := 1
	for ; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	j := 5
	for j < 10 {
		fmt.Printf("%d ", j)
		j++
	}
}
