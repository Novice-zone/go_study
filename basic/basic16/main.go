package main

import (
	"encoding/json"
	"fmt"
)

// 结构体字段可见性和JSON序列化

// go语言包中定义的标识符首字母大写就对外可见
// 结构体字段首字母大写就对外可见

type student struct {
	ID   int
	Name string
}

func newStudent(id int, name string) student {
	return student{
		ID:   id,
		Name: name,
	}
}

// 结构体tag，使用键值对的形式：`key1:"value1" key2:"value2"`
// 表示当使用json包的时候Title显示为title
type class struct {
	Title    string `json:"title"`
	Students []student
}

func main() {
	//创建c1班
	c1 := class{
		Title:    "冲刺班001",
		Students: make([]student, 0, 20),
	}
	//添加学生
	for i := 0; i < 10; i++ {
		tmpStu := newStudent(i, fmt.Sprintf("stu%02d", i))
		c1.Students = append(c1.Students, tmpStu)
	}
	//fmt.Printf("%v\n", c1)
	// JSON序列化：go语言数据->JSON格式字符串
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("json marshal failed ,err:", err)
		return
	}
	fmt.Printf("%T\n", data)
	fmt.Printf("%s\n", data)
	// JSON反序列化：JSON格式字符串->go语言数据
	jsonStr := `{"Title":"冲刺班001","Students":[{"ID":0,"Name":"糯小米"},{"ID":1,"Name":"stu01"},{"ID":2,"Name":"stu02"}]}`
	var c2 class
	err = json.Unmarshal([]byte(jsonStr), &c2)
	if err != nil {
		fmt.Println("json unmarshal failed,err:", err)
		return
	}
	fmt.Printf("%v\n", c2)
}
