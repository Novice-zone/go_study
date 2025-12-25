package res

import "github.com/gin-gonic/gin"

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func OK(c *gin.Context, data any, msg string) {

}

func OKWithData(c *gin.Context, data any) {

}

func OKWithMsg(c *gin.Context, msg string) {

}
