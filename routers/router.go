package router

import (
	"gin-demo/config"
	"gin-demo/middleware"
	"gin-demo/routers/testRouter"
	"github.com/gin-gonic/gin"
)


var (
	demogroup *gin.RouterGroup
)

func InitRouter() *gin.Engine {
	gin.SetMode(config.Conf.Server.Mode)
	r := gin.New()
	
	//初始化参数校验
	if err := middleware.TransInit("zh"); err != nil {
		middleware.Logger.Errorf("init trans failed:%v", err)
		return nil
	}

	// 404处理
	/*
	r.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method
		msg := fmt.Sprintf("%s %s not found", method, path)
		resp.NotFound(c,nil,msg)
	})
	 */

	InitRegisterRoute(r)

	return r
}
func InitRegisterRoute(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/gin-demo")
	
	v1 := g.Group("/v1")
	//v2 := g.Group("/v2")
	testRouter.RegisterTestRouter(v1)
	return g
}
