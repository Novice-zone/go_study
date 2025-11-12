package main

import "fmt"

type student struct {
	id    int
	name  string
	class string
}

// student的构造函数
func newStudent(id int, name, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}

type studentMgr struct {
	allStudents []*student
}

// newStudentMgr的构造函数
func newStudentMgr() *studentMgr {
	return &studentMgr{
		allStudents: make([]*student, 0, 100),
	}
}

// 验证学号是否已存在
func (s studentMgr) isIdExists(id int) bool {
	for _, stu := range s.allStudents {
		if stu.id == id {
			return true
		}
	}
	return false
}

// 1.添加学生
func (s *studentMgr) addStudent(newStu *student) {
	s.allStudents = append(s.allStudents, newStu)
}

// 2.编辑学生
func (s *studentMgr) modifyStudent(newStu *student) {
	//先通过id找
	for i, v := range s.allStudents {
		if newStu.id == v.id { //唯一的学号，相同表示有这个学生
			s.allStudents[i] = newStu //根据切片索引把新学生赋值进来
			fmt.Printf("学号%d已更新成功\n", newStu.id)
			return
		}
	}
	//找不到
	fmt.Printf("系统内无学号：%d的学生，输入有误\n", newStu.id)
}

// 3.展示学生
func (s *studentMgr) showStudent() {
	for _, v := range s.allStudents {
		fmt.Printf("学号：%d 姓名：%s 班级：%s\n", v.id, v.name, v.class)
	}
}
