package calc

import (
	"fmt"
	"go_study/basic/basic18/cuteDog"
)

// 标识符首字母大写才可以在包外被使用
// 通常不对外的标识符首字母小写

// Name是一个用于测试的全局变量
var Name = "糯小米"

// Add 是一个计算两个int之和的函数
func Add(x, y int) int {
	cuteDog.CuteDog()
	return x + y
}

// init函数在包导入时自动执行
// init函数无参无返
func init() {
	fmt.Println("calc.init()")
	fmt.Println(Name) // 导入包时全局变量先执行？？
}
