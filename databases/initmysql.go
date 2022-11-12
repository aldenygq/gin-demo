package databases

import (
	"gin-demo/config"
	"gin-demo/middleware"
	"log"
	"os"
	
	"strconv"
	"time"
	
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Sql *gorm.DB
func InitDB() {
	var err error
	// 数据库的类型
	dbType := config.Conf.Mysql.Type
	// Mysql配置信息
	mysqlName := config.Conf.Mysql.Dbname
	mysqlUser := config.Conf.Mysql.User
	mysqlPwd := config.Conf.Mysql.Pwd
	mysqlPort := strconv.Itoa(config.Conf.Mysql.Port)
	mysqlCharset := config.Conf.Mysql.Dbcharset
	mysqlHost := config.Conf.Mysql.Host
	
	var dataSource string
	dataSource = mysqlUser + ":" + mysqlPwd + "@tcp(" + mysqlHost + ":" +
		mysqlPort + ")/" + mysqlName + "?charset=" + mysqlCharset +
		"&parseTime=" + "true" + "&loc=" + "Local"
	log.Printf("dataSource:%v",dataSource)
	Sql, err = gorm.Open(dbType, dataSource)
	if Sql == nil {
		log.Printf("gorm mysql client invalid")
		os.Exit(-1)
	}
	if err != nil {
		log.Printf("connect mysql failed:%v", err)
		os.Exit(-1)
	}
	// 设置连接池，空闲连接
	Sql.DB().SetMaxIdleConns(config.Conf.Mysql.MaxIdleConns)
	// 打开链接
	Sql.DB().SetMaxOpenConns(config.Conf.Mysql.MaxOpenConns)
	//连接超时
	Sql.DB().SetConnMaxLifetime(time.Second * time.Duration(config.Conf.Mysql.MaxConnLifeTime))
	Sql.SetLogger(middleware.Logger)
	// 表明禁用后缀加s
	Sql.SingularTable(true)
	
	log.Printf("init db success\n")
}

