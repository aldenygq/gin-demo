package testRouter

import (
	"gin-demo/app/test"
	"github.com/gin-gonic/gin"
)

func RegisterTestRouter(v1 *gin.RouterGroup) {
	t := v1.Group("/test")
	{
		t.GET("/getinfo", test.GetInfo)
	}
}
