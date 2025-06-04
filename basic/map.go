package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// map
// map是一种无序的基于key-value（键值对）的数据结构
// map[KeyType]ValueType
func main() {
	// 1.map的基本使用
	//var scoreMap map[string]int
	//fmt.Println(scoreMap == nil)
	//scoreMap = make(map[string]int, 8)
	//scoreMap["小邱"] = 85
	//scoreMap["小段"] = 80
	//scoreMap["小邹"] = 90
	//fmt.Println(scoreMap)
	//fmt.Println(scoreMap["小段"])
	//fmt.Printf("type:%T\n", scoreMap)
	// 声明时填充元素
	//Stu := map[string]string{
	//	"三班":  "周翻倍",
	//	"十二班": "feifei",
	//}
	//fmt.Println(Stu)

	// 2.判断某个键值是否存在
	// value  ,   ok :=         map[key]
	// 对应键值，是否存在（bool）  map类型变量名["要找的键"]
	//scoreMap := make(map[string]int)
	//scoreMap["巴萨"] = 85
	//scoreMap["皇马"] = 80
	//scoreMap["马竞"] = 79
	//v, ok := scoreMap["马竞"]
	//fmt.Println(v, ok)
	//if ok {
	//	fmt.Println(v)
	//} else {
	//	fmt.Println("查无此队")
	//}
	//fmt.Println(scoreMap)
	// 3.map遍历
	//scoreMap := make(map[string]int)
	//scoreMap["cq"] = 75
	//scoreMap["ln"] = 68
	//scoreMap["sjz"] = 63
	//for k, v := range scoreMap {
	//	fmt.Printf("分数:%d,球队:%s\n", v, k)
	//}
	//// 只输出一项
	//for k := range scoreMap {
	//	fmt.Printf("前三球队分别是：%s\n", k)
	//}
	//for _, v := range scoreMap {
	//	fmt.Printf("前三甲积分分别为:%d\n", v)
	//}
	// 4.delete()函数删除键值对
	// delete(map,key)
	//scoreMap := make(map[string]int)
	//scoreMap["cq"] = 75
	//scoreMap["ln"] = 70
	//scoreMap["sjz"] = 68
	//delete(scoreMap, "sjz")
	//for k, v := range scoreMap {
	//	fmt.Println(k, v)
	//}
	// 按照某个固定顺序遍历map
	var scoreMap = make(map[string]int, 100)
	// 依次添加五十个随机键值对
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100) // 随机生成0~99的随机整数
		scoreMap[key] = value
	}
	/*// 打印检验是否添加成功
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}*/
	// 按照key值从小到大的顺序遍历scoreMap
	// 1.先取出所有key值，存放到slice中
	keys := make([]string, 0, 100) // 长度必须是0！！
	for k := range scoreMap {
		// for range函数遍历map默认第一个变量是键值，第二个是对应值
		// 因此这里就是取出scoreMap的key存放到切片keys中
		keys = append(keys, k)
	}
	// 2.对所有的key做排序
	sort.Strings(keys)
	// 3.按照排序后的key对scoreMap排序
	for _, key := range keys {
		// 这里的第一个变量是索引，第二个是元素值
		fmt.Println(key, scoreMap[key])
	}
}

// make函数
// 专门用于初始化 切片（slice）、映射（map）、通道（channel）
// 用于切片，length 是初始长度，capacity 是底层数组容量（可选，不填则和 length 相等）
// make(type, length, capacity)
// 用于 map 和通道，initialCapacity 对 map 是预分配桶数，对通道是缓冲区大小（可选）
// make(type, initialCapacity)
