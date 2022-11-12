package resp

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

type Response struct {
	// 代码
	Code int `json:"code" example:"200"`
	// 数据集
	Data interface{} `json:"data"`
	// 消息
	Msg string `json:"msg"`
}
func (res *Response) ReturnOK() *Response {
	res.Code = 200
	return res
}

func (res *Response) ReturnError() *Response {
	res.Code = 1001
	return res
}
func (res *Response) ReturnNotFound(code int) *Response {
	res.Code = 404
	return res
}

func Ok(c *gin.Context,data interface{}, msg string) {
	var res Response
	res.Data = data
	if msg != "" {
		res.Msg = msg
	}
	c.JSON(http.StatusOK, res.ReturnOK())
}
func Error(c *gin.Context, err error) {
	var res Response
	res.Msg = err.Error()

	c.JSON(http.StatusOK, res.ReturnError())
}

func NotFound(c *gin.Context,data interface{},msg string) {
	var res Response
	res.Data = data
	if msg != "" {
		res.Msg = msg
	}
	c.JSON(http.StatusNotFound, res.ReturnOK())
}

