package net

import (
	"net"
)

//校验ip是否合法
func CheckIp(str string) bool {
	result := net.ParseIP(str)
	if result == nil {
		return false
	}
	
	return true
}

//获取本地ip
func GetLocalIp() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	if conn == nil {
		return ""
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	return localAddr[0:idx]
}
