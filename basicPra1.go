package main

import "fmt"

// 变量定义
func main() {
	x := 10        // 使用 := 进行定义并初始化
	fmt.Println(x) // 输出语句 10
	x, y := 10, 20
	x, y = y+3, x+2   // 先计算等号右边值，然后再对x、y变量赋值
	fmt.Println(x, y) // 输出语句  结果为：23 12
}
