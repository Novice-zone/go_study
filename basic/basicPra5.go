package main

import "fmt"

// 基础语法：数组
func test_array() {
	//数组大小从声明时就确定，之后不可修改
	//基本定义：var 数组变量名 [元素数量]T
	//var a [4]int
	//var b [10]int
	//fmt.Println(a)
	//fmt.Println(b)
	//a=b是错的，因为[4]int和[10]int是不同的数据类型
	//通过下标访问，从0开始
	//初始化
	//1.定义时就初始化
	//var cityArray = [4]string{"北京", "上海", "广州", "深圳"}
	//fmt.Println(cityArray)
	////2.编译器推导数组长度
	//var boolArray = [...]bool{true, true, false, false, false}
	//fmt.Println(boolArray)
	////使用索引值初始化
	////var nums [5]int
	////for i := 0; i < 5; i++ {
	////	nums[i] = i
	////	fmt.Println(nums[i])
	////}
	////var langArray = [...]string{0: "Golang", 4: "C++", 9: "Java"}
	//langArray := [...]string{0: "Golang", 4: "C++", 9: "Java"}
	//fmt.Println(langArray)
	//
	//fmt.Printf("%T\n", langArray)

	//数组遍历
	//var cityArray = [4]string{"北京", "上海", "广州", "深圳"}
	//1.for循环遍历
	//for i := 0; i < len(cityArray); i++ {
	//	fmt.Println(cityArray[i])
	//}
	//2.for range 遍历
	//for index, value := range cityArray {
	//	fmt.Println(index, value)
	//}
	//for index := range cityArray {
	//	fmt.Println(index)
	//}
	//for _, value := range cityArray {
	//	fmt.Println(value)
	//}

	//二维数组
	cityArray := [...][2]string{ //多维数组只有外层可以用...让编译器推导大小
		{"北京", "西安"},
		{"上海", "杭州"},
		{"重庆", "成都"},
		{"广州", "深圳"}, //这里必须要有,
	}
	//fmt.Println(cityArray[2][0])
	//二维数组遍历
	for _, v1 := range cityArray {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	// 数组是值类型
	x := [3]int{1, 2, 3}
	fmt.Println(x)
	modifyArray(x)
	fmt.Println(x) //不改变数组的值
	//数组支持== !=等操作
	//[n]*T是指针数组 x:=[3]*int
	//*[n]T是数组指针 x:=*[3]int
}

func modifyArray(x [3]int) {
	x[0] = 100
}

// 练习
func praArray() {
	//数组元素求和
	/*a := [...]int{1, 3, 5, 7, 8}
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	fmt.Println(sum)*/
	//找出和为指定数值的两个元素的下标
	b := [...]int{1, 3, 5, 7, 8} //找出和为8的两个元素的下标
	/*for i := 0; i < len(b); i++ {
		for j := i + 1; j < len(b); j++ {
			if i == j {
				continue
			} else {
				if b[i]+b[j] == 8 {
					fmt.Printf("(%d,%d)\n", i, j)
				}
			}
		}
	}*/
	for i := 0; i < len(b); i++ {
		for j := i + 1; j < len(b); j++ {
			if b[i]+b[j] == 8 {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}

func main() {
	praArray()
}
