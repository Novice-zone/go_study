package main

import "fmt"

// 值接收者和指针接收者

// 接口嵌套
type animal interface {
	//move()
	//say()
	mover
	sayer
}

type mover interface {
	move()
}

type sayer interface {
	say()
}

type person struct {
	name string
	age  int8
}

//// 值接收者实现接口：类型的值和类型的指针都能保存到接口中
//func (p person) move() {
//	fmt.Printf("%s\n", p.name)
//}

// 指针接收者实现接口：只有类型指针可以保存到接口中
func (p *person) move() {
	fmt.Printf("%s\n", p.name)
}

// 一个类型可以实现多个接口
func (p *person) say() {
	fmt.Printf("%s在叫\n", p.name)
}
func main() {
	var m mover
	//p1 := person{ //p1是person类型的值
	//	name: "张三",
	//	age:  20,
	//}

	p2 := &person{ //p2是person类型的指针
		name: "李四",
		age:  21,
	}

	// m = p1// 值不能存到m中,因为p1是person类型的值，没有实现mover接口
	m = p2
	m.move()
	fmt.Println(m)

	var s sayer
	s = p2
	s.say()
	fmt.Println(s)

	var a animal
	a = p2
	a.move()
	a.say()
}
