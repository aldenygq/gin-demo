package databases

import (
	"gin-demo/config"
	
	//"gin-demo/middleware"
	//"log"
)

func Init() {
	if config.Conf.Redis !=  nil {
		IntiRedis()
	}
	
	if config.Conf.Mysql != nil{
		InitDB()
	}
}
