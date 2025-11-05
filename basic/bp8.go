package main

import "fmt"

// 通过嵌套结构体实现“继承”

// 动物
type Animal struct {
	name string
}

// 构造函数newAnimal
func newAnimal(name string) *Animal {
	return &Animal{
		name: name,
	}
}

// Animal专属方法move，设立接收者a，不改变接收者的值，不用指针类型
func (a Animal) move() {
	fmt.Printf("%s是一个动物，它会动\n", a.name)
}

// 狗
type Dog struct {
	size string
	//Animal *Animal
	*Animal //匿名嵌套
}

// 构造函数
func newDog(size, name string) *Dog {
	return &Dog{
		size:   size,
		Animal: newAnimal(name),
	}
}

// 专属方法，汪汪叫
//
//	func (d Dog) woof() {
//		fmt.Printf("%s狗%s可以汪汪叫\n", d.size,d.Animal.name)
//	}
func (d Dog) woof() {
	fmt.Printf("%s狗%s可以汪汪叫\n", d.size, d.name)
}

func main() {
	d1 := newDog("小", "糯米")
	d1.move()
	d1.woof()
	d1.size = "中小"
	d1.name = "米米"
	d1.move()
	d1.woof()
}
