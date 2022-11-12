package test

import (
	"gin-demo/middleware"
	"gin-demo/service/test"
	"gin-demo/tools/resp"
	"github.com/gin-gonic/gin"
)


func GetInfo(c *gin.Context)  {
	data,msg,err := test.GetInfo()
	if err != nil {
		middleware.Logger.Errorf("get info failed:%v",err)
		resp.Error(c,err)
	}
	resp.Ok(c,data,msg)
}
