package databases

import (
	"fmt"
	"gin-demo/config"
	"gin-demo/middleware"
	"github.com/go-redis/redis"
	"log"
	"os"
)

var RedisClient *redis.Client

func IntiRedis() {
	ip := config.Conf.Redis.Ip
	port := config.Conf.Redis.Port
	url := fmt.Sprintf("%s:%s", ip, port)
	middleware.Logger.Infof("url:%v",url)
	r := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "",
		DB:       0,
	})
	
	RedisClient = r
	if RedisClient == nil {
		log.Printf("connent redisl client failed")
		os.Exit(-1)
	}
	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Printf("check redis client connection failed:%v",err)
		os.Exit(-1)
	}
	log.Printf("init redis client success\n")
}

