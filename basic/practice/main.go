package main

//## 1. 基础类型转换与计算
//
//**要求：**
//
//将摄氏度转换为华氏度（公式：华氏度 = 摄氏度 × 9/5 + 32）。
//
//- 声明常量`boilingPointC`(100℃)
//- 接收用户输入的摄氏度值
//- 用两种方式打印结果：带单位字符串和不带单位

//const boilingPointC = 100.0
//
//func main() {
//	var celsius float64
//	fmt.Print("请输入摄氏度：")
//	fmt.Scanln(&celsius)
//
//	var fahrenheit float64
//	fahrenheit = (celsius*9)/5 + 32
//	fmt.Printf("对应的华氏度为：%.2f\n", fahrenheit) //不带单位
//	fmt.Printf("%.2f℃=%.2f℉", celsius, fahrenheit)
//}

// ## 2. 切片动态管理
//
// **要求**：
//
// 动态管理学生成绩切片。
//
// - 初始化空切片保存成绩
// - 添加三次成绩（85, 92, 78）
// - 删除第二个成绩
// - 打印最终成绩切片与平均分
//func main() {
//	//空切片
//	scores := []int{}
//
//	//append动态添加成绩
//	scores = append(scores, 85, 92, 78)
//	fmt.Println(scores) //输出原始数据
//
//	//删除第二个成绩
//	scores = append(scores[0:1], scores[2])
//	fmt.Println(scores) //检查是否删除
//
//	//打印最终成绩，并且平均分
//	fmt.Printf("最终成绩为：%v\n", scores)
//	total := 0
//	for _, v := range scores {
//		total += v
//	}
//	average := float64(total) / float64(len(scores))
//	fmt.Printf("平均分是：%.2f（保留两位小数）\n", average)
//}

//## 3. Map数据检索
//
//**要求**：创建国家-首都映射表
//
//- 初始化包含三组数据（中国→北京，日本→东京，美国→华盛顿）
//- 接收用户输入国家名
//- 检查输入是否存在，打印相应首都或错误信息
