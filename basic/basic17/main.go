package main

import (
	"fmt"
	"os"
)

//学员信息管理

//1.添加学员
//2.编辑学员信息
//3.展示学员信息

func showMenu() {
	fmt.Println("学员信息管理系统")
	fmt.Println("1.添加学员")
	fmt.Println("2.编辑学员信息")
	fmt.Println("3.展示学员信息")
	fmt.Println("4.退出")
}

// 获取用户输入的学员信息
func getInput() *student {
	var (
		id    int
		name  string
		class string
	)
	fmt.Println("展示学员信息")
	fmt.Print("请输入学员学号：")
	fmt.Scanf("%d\n", &id)
	fmt.Print("请输入学员姓名：")
	fmt.Scanf("%s\n", &name)
	fmt.Print("请输入学员班级：")
	fmt.Scanf("%s\n", &class)

	stu := newStudent(id, name, class)
	return stu
}

func main() {

	sm := newStudentMgr()
	for {
		//1.打印系统菜单
		showMenu()
		//2.等待用户选择选项
		var input int
		fmt.Printf("请输入你要操作的序号：")
		fmt.Scanf("%d\n", &input)
		fmt.Println("用户选择的是：", input)
		//3.执行选项
		switch input {
		case 1:
			//添加学员
			stu := getInput()
			if sm.isIdExists(stu.id) {
				fmt.Println("该学号已存在，请重新输入")
				break
			}
			sm.addStudent(stu)
			fmt.Printf("学号 %d（%s）添加成功！\n", stu.id, stu.name)
		case 2:
			//编辑学员
			stu := getInput()
			sm.modifyStudent(stu)
		case 3:
			//展示学员
			sm.showStudent()
		case 4:
			//退出
			fmt.Println("\n退出系统...")
			os.Exit(0)
		}
	}
}
