package main

import "fmt"

// 空接口
// 接口中没有定义任何方法时就是空接口
// 任意类型都实现了空接口 因此空接口可以存储任意值
// type xxx interface{}// 一般不用提前定义

//func main() {
//	var x interface{} // 空接口变量x
//
//	x = "hello"
//	fmt.Println(x)
//
//	x = 100
//	fmt.Println(x)
//}

// 空接口的应用
// 1.空接口类型作为函数的参数
// 2.可以作为map的value
func main() {
	var x interface{}
	x = "hello"
	//fmt.Println(x)

	x = 1
	//fmt.Println(x)

	x = 42
	//fmt.Println(x)

	x = true
	//fmt.Println(x)

	x = "nomi"
	//var m = make(map[string]interface{}, 16)
	//m["name"] = "nomi"
	//m["age"] = 1
	//m["hobby"] = []string{"吃饭", "睡觉"}
	//fmt.Println(m)

	// 断言 是括号里的类型就是true，不是就报错
	//ret := x.(bool)
	//fmt.Println(ret)

	ret, ok := x.(int) //猜的不对时，ok=false,ret=猜的类型的零值
	if !ok {
		fmt.Println("不是字符串类型")
	} else {
		fmt.Println("是字符串类型", ret)
	}

	// 使用switch进行类型断言
	switch value := x.(type) {
	case string:
		fmt.Printf("是string类型,value:%v\n", value)
	case int:
		fmt.Printf("是int类型,value:%v\n", value)
	case bool:
		fmt.Printf("是bool类型,value:%v\n", value)
	case float32:
		fmt.Printf("是float32类型,value:%v\n", value)
	default:
		fmt.Println("猜不到")
	}
}
