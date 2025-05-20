package main

func foo_01() (int, string) {
	return 14, "shuzhi"
}

func test01() {
	//变量
	/*var name string
	var age int
	var isOK bool

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isOK)

	var (
		a int16   = -10
		b uint16  = 10
		c float32 = 15.53
		d bool    = true
	)
	fmt.Println(a, b, c, d)*/

	/*x, _ := foo_01()
	_, y := foo_01()
	fmt.Println("x=", x)
	fmt.Println("y=", y)*/

	//常量

	/*
		//const pi=3.14159
		//const e=2.7182
		const (
			pi = 3.1415
			e  = 2.7182
		)
		fmt.Println(pi, e)
		const (
			n1 = 10
			n2
			n3
		)
		fmt.Println(n1, n2, n3)*/

	//iota
	/*const (
		n0 = iota
		n1
		_
		n3
		n4
	)
	fmt.Println(n0,n1, n3,n4)*/

	/*const (
		n0 = iota
		n1
		n2 = 200
		n3 = iota
		n4
		_
		n6
	)
	fmt.Println(n0, n1, n2, n3, n4, n6)*/

	/*	const (
			a, b = iota + 1, iota + 2
			c, d
			e, _
		)
		fmt.Println(a, b)
		fmt.Println(c, d)
		fmt.Println(e)*/
}

func test02() () {
	//基本数据类型
	
}
func main() {
	test01()
}
