package main

import "fmt"

type dog struct{}

func (d dog) say() {
	fmt.Println("woof")
}

type cat struct{}

func (c cat) say() {
	fmt.Println("meow")
}

type person struct {
	name string
}

func (p person) say() {
	fmt.Println("ahhhhh")
}

// 接口不管是什么类型，只管你要实现什么方法
// 定义一个抽象类型，只要实现了say()方法的类型都可以称之为sayer类型
type sayer interface {
	say()
}

func da(arg sayer) {
	arg.say() //不管传进来的是什么，都打，打了就叫，执行say方法
}

func main() {
	c1 := cat{}
	da(c1)

	d1 := dog{}
	da(d1)

	p1 := person{
		name: "fanbei",
	}
	da(p1)
}
