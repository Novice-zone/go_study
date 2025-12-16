package calc

import (
	"fmt"
	"go_study/basic/basic18/cuteDog"
)

// 标识符首字母大写才可以在包外被使用
// 通常不对外的标识符首字母小写

// Name 是一个用于测试的全局变量
var Name = "糯小米"

// Add 是一个计算两个int之和的函数
func Add(x, y int) int {
	cuteDog.CuteDog()
	return x + y
}

// 初始化函数init
// init函数在包导入时自动执行
// init函数无参无返
func init() {
	fmt.Println("calc.init()")
	fmt.Println(Name) // 导入包时全局变量先执行？？
}

// 执行顺序： 全部声明 --> init() --> main()
// 如果多层导入，从内到外依次执行
//一个包的初始化过程是按照代码中引入的顺序来进行的
//所有在该包中声明的init函数都将被串行调用并且仅调用执行一次。
//每一个包初始化的时候都是先执行依赖的包中声明的init函数再执行当前包中声明的init函数。
//确保在程序的main函数开始执行时所有的依赖包都已初始化完成。
