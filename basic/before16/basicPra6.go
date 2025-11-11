package main

import "fmt"

// 切片
// 切片本质就是对底层数组的封装
// 包含三个信息：底层数组的指针，切片长度，切片容量
func test_Slice() {
	//定义
	//var name []T 括号里不用写长度
	//var a []string
	//var b = []int{}
	//var c = []bool{false, true}
	//
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)

	//基于数组得到切片
	/*a := [5]int{17, 18, 27, 40, 60}
	//[]里面的是索引值（下标）
	b := a[1:4]
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	//基于切片的切片
	c := b[0:1]
	fmt.Println(c)
	fmt.Printf("%T\n", c)
	//make函数构造切片make(类型,切片长度,切片容量（底层数组长度）)
	d := make([]int, 5, 10)
	fmt.Println(d)
	//通过len()函数获取切片长度
	e := len(d)
	fmt.Println(e)
	//通过cap()函数获取切片容量
	fmt.Println(cap(d))*/

	//切片不能直接"=="比较,不能用nil判断是否为空，用len即可
	/*var a []int //声明
	fmt.Println(a == nil)
	fmt.Println("长度:", len(a)) //0
	fmt.Println("容量:", cap(a)) //0

	b := []int{1, 2, 3}
	fmt.Println(b == nil)
	fmt.Println("长度:", len(b)) //3
	fmt.Println("容量:", cap(b)) //3

	//fmt.Println(a==b) 编译错误

	s := make([]int, 3, 5)
	fmt.Println(s == nil)
	fmt.Println("长度:", len(s)) // 输出: 3
	fmt.Println("容量:", cap(s)) // 输出: 5*/

	//切片赋值拷贝
	/*a := make([]int, 3) //[0] [0] [0]
	b := a              //b和a指向一块内存空间
	b[0] = 100
	fmt.Println(a)
	fmt.Println(b)*/

	//切片遍历
	a := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(a); i++ {
		//fmt.Println(i, a[i])
		fmt.Printf("%d ", a[i])
	}
	fmt.Println()
	for _, value := range a {
		//fmt.Println(index, value)
		fmt.Printf("%d ", value)
	}
}

func main() {
	test_Slice()
}
