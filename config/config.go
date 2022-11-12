package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var Conf = new(Config)
func InitConfig() {
	path, err := os.Getwd()
	if err != nil {
		log.Printf("get pwd path failed:%v\n", err)
		os.Exit(-1)
	}
	filepath := path + CONFIG_DIR
	c := viper.New()
	c.SetConfigFile(CONFIG_FILE)
	c.AddConfigPath(filepath)         //设置读取的文件路径
	c.SetConfigName(CONFIG_FILE_NAME) //设置读取的文件名
	c.SetConfigType(CONFIG_TYPE)      //chaos设置文件的类型

	err = c.ReadInConfig() // 搜索并读取配置文件
	if err != nil {        // 处理错误
		log.Printf("read config file failed:%v\n", err)
		os.Exit(-1)
	}
	err = c.Unmarshal(&Conf) //将配置文件绑定到config上
	if err != nil {
		log.Printf("unmarshal config info failed:%v\n", err)
		os.Exit(-1)
	}

	log.Printf("init config success")
}
