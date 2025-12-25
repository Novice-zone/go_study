package main

import "github.com/gin-gonic/gin"

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Index(c *gin.Context) {

	c.JSON(200, Response{
		Code: 0,
		Msg:  "success",
		Data: map[string]any{},
	})
}

func Index1(c *gin.Context) {
	// gin.H 是一个 map[string]any 类型的别名,会按键名字母顺序输出
	//最好使用定义结构体的方式 保证输出格式统一
	c.JSON(200, gin.H{
		"code": 111,
		"msg":  "success",
		"data": map[string]any{},
	})
}

func main() {
	// gin.SetMode("release") // 发布模式,会屏蔽一些调试信息
	// 1.初始化
	r := gin.Default()

	//2.挂载路由
	r.GET("/index", Index)
	r.GET("/index1", Index1)

	// 3.绑定端口运行
	r.Run("127.0.0.1:8080") // http.ListenAndServe

}
