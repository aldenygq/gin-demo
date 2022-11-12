package tools

import (
	"fmt"
	"time"
)

//耗时统计函数(优雅方式)
func TimeGraceCost() func() {
	start := time.Now()
	return func() {
		tc := time.Since(start)
		fmt.Printf("time cost = %v\n", tc)
	}
}

//耗时统计函数(简洁方式)
func timeCost(start time.Time) {
	tc := time.Since(start)
	fmt.Printf("time cost = %v\n", tc)
}

//获取当前时间(unix时间戳)
func GetCurrntTime() time.Time {
	return time.Now()
}

func GetCurrntTimeStr() string {
		return time.Now().Format("2006-01-02 15:04:05")
}
