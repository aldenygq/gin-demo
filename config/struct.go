package config

import "time"

type Config struct {
	Log   *LogConfig `mapstructure:"log"`
	Mysql *DataBase  `mapstructure:"mysql"`
	Util   *UtilInfo  `mapstructure:"util"`
	Redis  *RedisInfo  `mapstructure:"redis"`
	Server *ServerInfo `mapstructure:"server"`
}
type UtilInfo struct {

}
type LogConfig struct {
	Loglevel  string        `yaml:"loglevel"`
	Logfile   string        `yaml:"logfile"`
	Logmaxage time.Duration `yaml:"logmaxage"`
}
type DataBase struct {
	Host            string `yaml:"host"`
	User            string `yaml:"user"`
	Dbname          string `yaml:"dbname"`
	Pwd             string `yaml:"pwd"`
	Port            int    `yaml:"port"`
	MaxIdleConns    int    `yaml:"maxIdleConns"`
	MaxOpenConns    int    `yaml:"maxOpenConns"`
	MaxConnLifeTime int    `yaml:"maxConnLifeTime"`
	Type            string `yaml:"type"`
	Dbcharset       string `yaml:"dbcharset"`
}
type ServerInfo struct {
	IsHttps bool `yaml:"ishttps"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Mode string `yaml:"mode"`
	Ssl SslInfo `yaml:"ssl"`
}
type SslInfo struct {
	Key string `json:"key"`
	Pem string `json:"pem"`
}

type RedisInfo struct {
	Ip   string `yaml:"ip"`
	Port string `yaml:"port"`
}
