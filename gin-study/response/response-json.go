package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化
	r := gin.Default()

	// 挂载路由
	r.GET("/index", func(c *gin.Context) {
		//c.JSON(201, gin.H{
		//	"code": 0,
		//	"msg":  "success",
		//	"data": map[string]any{},
		//})
		//res.OK()
		//res.OKWithData()
		//res.OKWithMsg()
	})

	// 绑定端口
	r.Run(":8080")
}
