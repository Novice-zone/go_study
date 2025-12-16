package main

import (
	"fmt"
	xmm "go_study/basic/basic18/package_demo/calc"
	//当包名冲突就起个别名/包名很长起个别名
)

// Go不允许导入包不使用
// Go不允许循环引用包！！！

// 代码都在GOPATH目录下的话
// 导入包的路径要从$GOPATH/src后面写起

// import "go_study/basic/basic18/package_demo/calc"

func main() {
	fmt.Println("hi")
	ret := xmm.Add(4, 2)
	fmt.Println(ret)
	fmt.Println(xmm.Name)
}
